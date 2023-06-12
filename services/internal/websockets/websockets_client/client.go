package websocketsclient

import "github.com/gorilla/websocket"

// Client implements the WebSocketClient interface.
type Client struct {
	conn *websocket.Conn
}

// NewClient creates a new WebSocket client.
func NewClient() *Client {
	return &Client{}
}

// Connect establishes a WebSocket connection with the specified URL.
func (c *Client) Connect(url string) error {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return err
	}

	c.conn = conn
	return nil
}
