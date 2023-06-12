package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("notification", func() {
	Description("Notification service is responsible for handling user data and requests")
	HTTP(func() {
		Path("/api/chat/v1/notifications") // Prefix to HTTP path of all requests.
	})

	Method("notifyUserNamesUpdate", func() {
		Payload(func() {
			Field(1, "id", String)
			Field(2, "firstName", String)
			Field(3, "oldFirstName", String)
			Field(4, "lastName", String)
			Field(5, "oldLastName", String)
			Field(6, "ts", String)
			Required("id", "oldLastName", "oldFirstName", "firstName", "lastName", "ts")
		})

		GRPC(func() {

		})

	})
})

// var UpdatedUserNotification = Type("UpdatedUserNotification", func() {
// 	Field(1, "id", String)
// 	Field(2, "firstName", String)
// 	Field(3, "oldFirstName", String)
// 	Field(4, "lastName", String)
// 	Field(5, "oldLastName", String)
// 	Required("id", "oldLastName", "oldFirstName", "firstName", "lastName")
// })
