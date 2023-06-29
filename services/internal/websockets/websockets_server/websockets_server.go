package websocketsserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Server implements the WebSocketServer interface.
type Server struct {
	upgrader    *websocket.Upgrader
	connections []*websocket.Conn
}

// NewServer creates a new WebSocket server.
func NewServer() *Server {
	upgrader := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// Allow connections from any origin
			return true
		},
	}

	return &Server{
		upgrader:    upgrader,
		connections: make([]*websocket.Conn, 0),
	}
}

// Start starts the WebSocket server and listens for incoming connections.
func (s *Server) Start() error {
	http.HandleFunc("/", s.handleWebSocket)
	return http.ListenAndServe(":6001", nil)
}

// Close closes all active WebSocket connections.
func (s *Server) Close() error {
	for _, conn := range s.connections {
		err := conn.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

// BroadcastMessage sends a message to all connected clients.
func (s *Server) BroadcastMessage(message []byte) {
	for _, conn := range s.connections {
		err := conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Printf("Error sending message: %v", err)
		}
	}
}

// handleWebSocket handles incoming WebSocket connections.
func (s *Server) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection to WebSocket: %v", err)
		return
	}

	s.connections = append(s.connections, conn)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading WebSocket message: %v", err)
			break
		}

		fmt.Printf("Received message: %s\n", string(msg))

		s.BroadcastMessage(msg)
	}

	err = conn.Close()
	if err != nil {
		log.Printf("Error closing connection: %v", err)
	}
}
