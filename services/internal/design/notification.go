package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("notification", func() {
	Description("Notification service is responsible for handling user data and requests")

	Method("notifyUserNamesUpdate", func() {
		Payload(func() {
			Field(1, "id", String, func() {
				Meta("rpc:tag", "1")
			})
			Field(2, "firstName", String, func() {
				Meta("rpc:tag", "2")
			})
			Field(3, "oldFirstName", String,
				func() {
					Meta("rpc:tag", "3")
				})
			Field(4, "lastName", String, func() {
				Meta("rpc:tag", "4")
			})
			Field(5, "oldLastName", String, func() {
				Meta("rpc:tag", "5")
			})
			Field(6, "ts", String, func() {
				Meta("rpc:tag", "6")
			})

			Required("id", "oldLastName", "oldFirstName", "firstName", "lastName", "ts")
		})

		Result(BlankResponse)

		GRPC(func() {

		})

	})
})

var BlankResponse = Type("BlankResponse", func() {
	Field(1, "id", String)

	Required("id")
})
