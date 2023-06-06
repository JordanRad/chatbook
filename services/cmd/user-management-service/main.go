package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/JordanRad/chatbook/services/cmd/user-management-service/info"
	"github.com/JordanRad/chatbook/services/cmd/user-management-service/user"
	"github.com/JordanRad/chatbook/services/cmd/user-management-service/userauth"

	"github.com/JordanRad/chatbook/services/internal/auth"

	"github.com/JordanRad/chatbook/services/internal/auth/encryption"
	"github.com/JordanRad/chatbook/services/internal/auth/jwt"
	"github.com/JordanRad/chatbook/services/internal/databases/postgresql"
	"github.com/JordanRad/chatbook/services/internal/databases/postgresql/user-management/migrations"

	authgen "github.com/JordanRad/chatbook/services/internal/gen/auth"
	authsvc "github.com/JordanRad/chatbook/services/internal/gen/auth"
	authsrv "github.com/JordanRad/chatbook/services/internal/gen/http/auth/server"

	infosrv "github.com/JordanRad/chatbook/services/internal/gen/http/info/server"
	infosvc "github.com/JordanRad/chatbook/services/internal/gen/info"

	usersrv "github.com/JordanRad/chatbook/services/internal/gen/http/user/server"
	usersvc "github.com/JordanRad/chatbook/services/internal/gen/user"

	"github.com/JordanRad/chatbook/services/internal/middleware"

	goahttp "goa.design/goa/v3/http"
)

func main() {
	// Read the configuration
	config, err := configFromEnv()
	if err != nil {
		log.Fatalf("Config file cannot be read: %v", err)
	}
	fmt.Println(config)

	// Connect to database

	db := postgresql.ConnectToDatabase(config.Postgres.User, config.Postgres.Password, config.Postgres.Host, config.Postgres.Port, config.Postgres.DBName)
	migrationTool := migrations.Tool{
		DB: db,
	}

	withMockData := false
	if config.Postgres.Mode == "DEV" {
		withMockData = true
	}

	err = migrationTool.ApplyMigrations(withMockData)
	if err != nil {
		log.Fatalf("Error applying table creating migrations: %v", err)
	}

	// Initialize loger
	logger := log.New(os.Stdout, "", log.LstdFlags)

	// Initialize Info Service
	infoService := info.NewService()
	var infoEndpoints *infosvc.Endpoints = infosvc.NewEndpoints(infoService)

	// Initialize User Profile Service
	userStore := &auth.Store{
		DB: db,
	}

	jwtService := &jwt.JWTService{}
	encryptionTool := &encryption.Encrypter{}

	userService := user.NewService(userStore, encryptionTool, logger)
	var userEndpoints *usersvc.Endpoints = usersvc.NewEndpoints(userService)

	authService := userauth.NewService(userStore, encryptionTool, jwtService, logger)
	var authEndpoints *authsvc.Endpoints = authgen.NewEndpoints(authService)

	// Provide the transport specific request decoder and response encoder.
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	var mux goahttp.Muxer
	{
		mux = goahttp.NewMuxer()
	}

	// Initialize Info Server
	var infoServer *infosrv.Server = infosrv.New(infoEndpoints, mux, dec, enc, nil, nil)
	infosrv.Mount(mux, infoServer)

	// Initialize User Profile Server
	var userServer *usersrv.Server = usersrv.New(userEndpoints, mux, dec, enc, nil, nil)
	userServer.Use(middleware.AuthenticateRequest(userStore, jwtService))
	usersrv.Mount(mux, userServer)

	// Initialize User Server
	var authServer *authsrv.Server = authsrv.New(authEndpoints, mux, dec, enc, nil, nil)
	authServer.Use(middleware.AuthenticateRequest(userStore, jwtService))
	authsrv.Mount(mux, authServer)

	// Start the HTTP server
	address := fmt.Sprintf("%s:%d", config.HTTP.Host, config.HTTP.Port)
	log.Printf("User Management service has just started on %s ...\n", address)
	http.ListenAndServe(address, mux)
}
