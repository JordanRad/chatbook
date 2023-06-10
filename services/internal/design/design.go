package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var _ = API("chatbook-microservices", func() {
	Title("Chatbook Microservices")

	//NOTE(JordanRad): Replace localhost with the deployed origin
	cors.Origin("/.*localhost.*/", func() {
		cors.Headers("*")
		cors.Methods("GET", "POST", "DELETE", "PUT", "PATCH", "OPTIONS")
		cors.MaxAge(600)
	})
})

var _ = Service("info", func() {
	Description("Application info")

	HTTP(func() {
		Path("/api/user-management/v1/info") // Prefix to HTTP path of all requests.
	})

	Method("getInfo", func() {

		Result(OperationStatusResponse)
		HTTP(func() {
			GET("/")
		})
	})
})

var OperationStatusResponse = Type("OperationStatusResponse", func() {
	Attribute("message", String, "Operation status")

	Required("message")
})
