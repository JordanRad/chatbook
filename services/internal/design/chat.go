package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("chat", func() {
	Description("User service is responsible for handling user data and requests")
	HTTP(func() {
		Path("/api/chat/v1/") // Prefix to HTTP path of all requests.
	})

	Method("getChatHistory", func() {
		Result(ChatHistoryResponse)
		Payload(func() {
			Attribute("ID", String, "Chatroom ID")
		})
		HTTP(func() {
			GET("/chatrooms/{ID}/history")
		})

	})

})

var ChatHistoryResponse = Type("ChatHistoryResponse", func() {
	Attribute("id", String, "Chatroom ID")
	Attribute("count", Int, "Messages Count")
	Attribute("messages", ArrayOf(ChatMessage), "Chat history")

	Required("id", "count", "messages")
})

var ChatMessage = Type("ChatMessage", func() {
	Attribute("previousMessageID", String, "Previous message ID")
	Attribute("nextMessageID", String, "Next message ID")
	Attribute("timestamp", Float64, "Timestamp of the message")

	Required("timestamp", "previousMessageID", "nextMessageID")
})
