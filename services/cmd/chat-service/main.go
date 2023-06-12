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

	infosrv "github.com/JordanRad/chatbook/services/internal/gen/http/info/server"
	infosvc "github.com/JordanRad/chatbook/services/internal/gen/info"

	notificationsrv "github.com/JordanRad/chatbook/services/internal/gen/grpc/notification/server"
	notificationsvc "github.com/JordanRad/chatbook/services/internal/gen/notification"

	chatsvc "github.com/JordanRad/chatbook/services/internal/gen/chat"
	chatsrv "github.com/JordanRad/chatbook/services/internal/gen/http/chat/server"

	"github.com/JordanRad/chatbook/services/internal/middleware"
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
	// db := postgresql.ConnectToDatabase(config.Postgres.User, config.Postgres.Password, config.Postgres.Host, config.Postgres.Port, config.Postgres.DBName)
	// migrationTool := migrations.Tool{
	// 	DB: db,
	// }

	// withMockData := false
	// if config.Postgres.Mode == "DEV" {
	// 	withMockData = true
	// }

	// err = migrationTool.ApplyMigrations(withMockData)
	// if err != nil {
	// 	log.Fatalf("Error applying table creating migrations: %v", err)
	// }

	// Initialize loger
	logger := log.New(os.Stdout, "", log.LstdFlags)

	// Initialize Info Service
	infoService := info.NewService()
	var infoEndpoints *infosvc.Endpoints = infosvc.NewEndpoints(infoService)

	//Note (JordanRad): Add store
	// Initialize Chat Service
	chatStore := &dbchat.Store{
		DB: nil,
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

	go func() {
		s := notifiation.NewService()
		grpcEndpoints := notificationsvc.NewEndpoints(s)
		grpcsrv := notificationsrv.New(grpcEndpoints, nil)

		grpcServer := grpc.NewServer()
		fmt.Printf("Notifications gRPC server has just started on %d ...\n", 5002)
		notificationsprotobuf.RegisterNotificationServer(grpcServer, grpcsrv)
		lis, err := net.Listen("tcp", "localhost:5002")

		if err != nil {
			panic(err)
		}
		if err := grpcServer.Serve(lis); err != nil {
			panic(err)
		}
	}()

	// Initialize Info Server
	var infoServer *infosrv.Server = infosrv.New(infoEndpoints, mux, dec, enc, nil, nil)
	infosrv.Mount(mux, infoServer)

	//Note (JordanRad): Add store
	userStore := auth.NewStore(nil)
	j := &jwt.JWTService{}
	// Initialize Chat Server
	var chatServer *chatsrv.Server = chatsrv.New(chatEndpoints, mux, dec, enc, nil, nil)
	chatServer.Use(middleware.AuthenticateRequest(userStore, j))
	chatsrv.Mount(mux, chatServer)

	// notificationsprotobuf "github.com/JordanRad/chatbook/services/internal/gen/grpc/user/pb"
	// notificationsgrpcsrv "github.com/JordanRad/chatbook/services/internal/gen/grpc/user/server"
	// go func() {
	// 	grpcEndpoints := usersvc.NewEndpoints(userService)
	// 	grpcsrv := notificationsgrpcsrv.New(grpcEndpoints, nil)

	// 	grpcServer := grpc.NewServer()
	// 	fmt.Printf("Notifications gRPC server has just started on %d ...\n", 5002)
	// 	notificationsprotobuf.RegisterUserServer(grpcServer, grpcsrv)
	// 	lis, err := net.Listen("tcp", "localhost:5002")

	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	if err := grpcServer.Serve(lis); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// Start the HTTP server
	address := fmt.Sprintf("%s:%d", config.HTTP.Host, config.HTTP.Port)
	log.Printf("Chat service has just started on %s ...\n", address)
	http.ListenAndServe(address, mux)
}
