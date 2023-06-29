package websocketsserver

// WebSocketServer defines the methods required for a websocket server.
type WebSocketServer interface {
	Start() error
	Close() error
	BroadcastMessage(message []byte)
}
