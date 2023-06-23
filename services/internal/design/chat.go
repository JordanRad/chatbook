package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("chat", func() {
	Description("User service is responsible for handling user data and requests")
	HTTP(func() {
		Path("/api/chat/v1/") // Prefix to HTTP path of all requests.
	})

	Method("getConversationHistory", func() {
		Result(ChatHistoryResponse)
		Payload(func() {
			Attribute("ID", String, "Conversation ID")
			Attribute("limit", Int, "Messages count", func() {
				Default(200)
			})
			Attribute("beforeTimestamp", Int64, "Before timestamp", func() {
				Default(1257894000)
			})

			Required("ID")

		})
		HTTP(func() {
			GET("/conversations/{ID}/history")
			Param("limit")
			Param("beforeTimestamp")
		})
	})

	Method("searchInConversation", func() {
		Result(ChatHistoryResponse)
		Payload(func() {
			Attribute("ID", String, "Conversation ID")
			Attribute("limit", Int, "Messages count", func() {
				Default(200)
			})
			Attribute("searchInput", String, "Input", func() {
				Default("200")
				MinLength(5)
			})

			Required("ID")

		})
		HTTP(func() {
			GET("/conversations/{ID}")
			Param("limit")
			Param("searchInput")
		})
	})

	Method("getConversationsList", func() {
		Result(ConversationsListResponse)
		Payload(func() {
			Attribute("limit", Int, "Messages count", func() {
				Default(100)
			})
		})
		HTTP(func() {
			GET("/conversations")
			Param("limit")
		})

	})

	Method("addConversation", func() {
		Result(OperationStatusResponse)
		Payload(func() {
			Attribute("participants", ArrayOf(Friend), "Participants")

			Required("participants")
		})
		HTTP(func() {
			POST("/conversation")

		})

	})

})

var ChatHistoryResponse = Type("ChatHistoryResponse", func() {
	Attribute("id", String, "Chatroom ID")
	Attribute("count", Int, "Messages Count")
	Attribute("messages", ArrayOf(ConversationMessage), "Chat history")

	Required("id", "count", "messages")
})

var ConversationMessage = Type("ConversationMessage", func() {
	Attribute("senderID", String, "Sender ID")
	Attribute("timestamp", Float64, "Timestamp of the message")
	Attribute("content", String, "Message Content")

	Required("timestamp", "senderID", "content")
})

var ConversationsListResponse = Type("ConversationsListResponse", func() {
	Attribute("total", Int, "Messages Count")
	Attribute("resources", ArrayOf(Conversation), "Chat history")

	Required("total", "resources")
})

var Conversation = Type("Conversation", func() {
	Attribute("ID", String, "Conversation ID")
	Attribute("lastMessageSenderID", String, "Sender ID")
	Attribute("lastMessageContent", String, "Last message")
	Attribute("lastMessageDeliveredAt", String, "TS for delivered time")

	Required("ID", "lastMessageSenderID", "lastMessageContent", "lastMessageDeliveredAt")
})
