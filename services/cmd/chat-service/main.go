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

	infosrv "github.com/JordanRad/chatbook/services/internal/gen/http/info/server"
	infosvc "github.com/JordanRad/chatbook/services/internal/gen/info"

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

	// Initialize Info Server
	var infoServer *infosrv.Server = infosrv.New(infoEndpoints, mux, dec, enc, nil, nil)
	infosrv.Mount(mux, infoServer)

	// Initialize Chat Server
	var chatServer *chatsrv.Server = chatsrv.New(chatEndpoints, mux, dec, enc, nil, nil)
	// chatServer.Use(middleware.AuthenticateRequest(userStore, j))
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
	// Start the HTTP server
	address := fmt.Sprintf("%s:%d", config.HTTP.Host, config.HTTP.Port)
	log.Printf("Chat service has just started on %s ...\n", address)
	http.ListenAndServe(address, mux)
}
