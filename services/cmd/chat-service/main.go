package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/JordanRad/chatbook/services/cmd/chat-service/chat"
	"github.com/JordanRad/chatbook/services/cmd/chat-service/db/dbchat"
	"github.com/JordanRad/chatbook/services/cmd/chat-service/notifiation"
	"github.com/JordanRad/chatbook/services/cmd/user-management-service/info"

	"github.com/JordanRad/chatbook/services/internal/auth"
	"github.com/JordanRad/chatbook/services/internal/auth/jwt"
	"github.com/JordanRad/chatbook/services/internal/databases/postgresql"
	migrations "github.com/JordanRad/chatbook/services/internal/databases/postgresql/migrations/chat"
	infosrv "github.com/JordanRad/chatbook/services/internal/gen/http/info/server"
	infosvc "github.com/JordanRad/chatbook/services/internal/gen/info"
	"github.com/JordanRad/chatbook/services/internal/middleware"
	websocketsserver "github.com/JordanRad/chatbook/services/internal/websockets/websockets_server"

	notificationsrv "github.com/JordanRad/chatbook/services/internal/gen/grpc/notification/server"
	notificationsvc "github.com/JordanRad/chatbook/services/internal/gen/notification"

	chatsvc "github.com/JordanRad/chatbook/services/internal/gen/chat"
	chatsrv "github.com/JordanRad/chatbook/services/internal/gen/http/chat/server"

	"google.golang.org/grpc"

	notificationsprotobuf "github.com/JordanRad/chatbook/services/internal/gen/grpc/notification/pb"
	goahttp "goa.design/goa/v3/http"
)

func main() {
	// Read the configuration
	config, err := configFromEnv()
	if err != nil {
		log.Fatalf("Config file cannot be read: %v", err)
	}
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

	//Note (JordanRad): Add store
	// Initialize Chat Service
	chatStore := &dbchat.Store{
		DB: db,
	}

	chatService := chat.NewService(logger, chatStore)
	var chatEndpoints *chatsvc.Endpoints = chatsvc.NewEndpoints(chatService)

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

	// Initialize Chat Server
	var chatServer *chatsrv.Server = chatsrv.New(chatEndpoints, mux, dec, enc, nil, nil)
	userStore := auth.NewStore(db)
	jwtService := &jwt.JWTService{}
	chatServer.Use(middleware.AuthenticateRequest(userStore, jwtService))

	chatsrv.Mount(mux, chatServer)

	go func() {
		s := notifiation.NewService()

		grpcEndpoints := notificationsvc.NewEndpoints(s)
		grpcsrv := notificationsrv.New(grpcEndpoints, nil)
		grpcServer := grpc.NewServer()
		notificationsprotobuf.RegisterNotificationServer(grpcServer, grpcsrv)
		lis, err := net.Listen("tcp", "localhost:5002")
		if err != nil {
			panic(err)
		}

		log.Printf("Notifications gRPC server has just started on  %s ...\n", lis.Addr().String())
		if err := grpcServer.Serve(lis); err != nil {
			panic(err)
		}
	}()

	go func() {
		websocketsServer := websocketsserver.NewServer()

		log.Printf("Websockets server started on  %d ...\n", 6001)
		err := websocketsServer.Start()
		if err != nil {
			panic(err)
		}
	}()
	// Start the HTTP server
	address := fmt.Sprintf("%s:%d", config.HTTP.Host, config.HTTP.Port)
	log.Printf("Chat service has just started on %s ...\n", address)
	http.ListenAndServe(address, mux)
}
