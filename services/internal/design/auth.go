package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("auth", func() {
	Description("Authentication service is responsible for handling user data and requests")
	HTTP(func() {
		Path("/api/user-management/v1/auth") // Prefix to HTTP path of all requests.
	})

	Method("refreshToken", func() {
		Payload(func() {
			Attribute("email", String, "User email")
			Attribute("refreshToken", String, "Refresh token")
			Required("email", "refreshToken")
		})

		Result(LoginResponse)
		HTTP(func() {
			POST("/refresh-token")
		})

	})

	Method("login", func() {
		Payload(func() {
			Attribute("email", String, "User email")
			Attribute("password", String, "User password")
			Required("email", "password")
		})
		Result(LoginResponse)
		Error("WrongCredentials")
		HTTP(func() {
			POST("/login")
			Response("WrongCredentials", StatusNotFound, func() {
				Description("Email and/or password don't match")
			})
		})
	})

})

var RegisterResponse = Type("RegisterResponse", func() {
	Attribute("message", String, "Operation status")
	Required("message")
})

var LoginResponse = Type("LoginResponse", func() {
	Attribute("email", String, "User's email")
	Attribute("token", String, "JWT Token")
	Attribute("refresh_token", String, "Refresh token")
	Attribute("role", String, "User's role")
	Attribute("id", String, "User account's ID")

	Required("email", "token", "refresh_token", "role")
})
