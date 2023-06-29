package websocketsserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Store interface {
	SaveConversationMessage(ID, senderID, content string) error
}
type ChatConnection struct {
	ID   string
	conn *websocket.Conn
}

// Server implements the WebSocketServer interface.
type Server struct {
	upgrader    *websocket.Upgrader
	connections []*ChatConnection
	store       Store
}

// NewServer creates a new WebSocket server.
func NewServer(s Store) *Server {
	upgrader := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// Allow connections from any origin
			return true
		},
	}

	return &Server{
		upgrader:    upgrader,
		connections: make([]*ChatConnection, 0),
		store:       s,
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
		err := conn.conn.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

// BroadcastMessage sends a message to all connected clients.
func (s *Server) BroadcastMessage(message []byte, senderID string) {
	type OutgoingMessage struct {
		Content  string `json:"content"`
		SenderID string `json:"senderID"`
	}

	m := OutgoingMessage{
		Content:  string(message),
		SenderID: senderID,
	}

	for _, conn := range s.connections {
		err := conn.conn.WriteJSON(m)
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

	conversationID := r.URL.Query().Get("conversationID")
	senderID := r.URL.Query().Get("senderID")

	chatConn := &ChatConnection{
		conn: conn,
		ID:   conversationID,
	}

	s.connections = append(s.connections, chatConn)

	for {
		_, msg, err := chatConn.conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading WebSocket message: %v", err)
			break
		}

		s.BroadcastMessage(msg, senderID)

		err = s.store.SaveConversationMessage(conversationID, senderID, string(msg))
		if err != nil {
			fmt.Printf("error saving message to db: %v", err.Error())
		}
	}

	err = conn.Close()
	if err != nil {
		log.Printf("Error closing connection: %v", err)
	}
}
