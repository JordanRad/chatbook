package websocketsclient

// WebSocketClient defines the methods required for a websocket client.
type WebSocketClient interface {
	SendMessage(message string) error
	ReceiveMessage() ([]byte, error)
}
