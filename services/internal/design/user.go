package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("user", func() {
	Description("User service is responsible for handling user data and requests")
	HTTP(func() {
		Path("/api/user-management/v1/users") // Prefix to HTTP path of all requests.
	})

	Method("register", func() {
		Payload(func() {
			Attribute("firstName", String, "User first name")
			Attribute("lastName", String, "User last name")
			Attribute("email", String, "User email")
			Attribute("password", String, "User password")
			Attribute("confirmedPassword", String, "Confirmed password of the user")
			Required("email", "password", "confirmedPassword", "firstName", "lastName")
		})

		Result(RegisterResponse)
		Error("UniqueEmailError")
		Error("UnmatchingPassowrds")

		HTTP(func() {
			POST("/register")
			Response("UniqueEmailError", StatusConflict, func() {
				Description("Email already exists")
			})
			Response("UnmatchingPassowrds", StatusConflict, func() {
				Description("Provided passwords are not the same")
			})
		})
	})

	Method("getProfile", func() {
		Result(UserProfileResponse)
		HTTP(func() {
			GET("/profile")
		})

	})

	Method("updateProfileNames", func() {
		Payload(func() {
			Attribute("firstName", String, "Updated user first name")
			Attribute("lastName", String, "Updated user last name")
			Required("firstName", "lastName")
		})
		Result(OperationStatusResponse)
		HTTP(func() {
			PUT("/profile")
		})
	})

	Method("addFriend", func() {
		Payload(func() {
			Attribute("id", String, "User ID to add")

			Required("id")
		})
		Result(OperationStatusResponse)
		HTTP(func() {
			POST("/friend")
		})
	})

	Method("removeFriend", func() {
		Payload(func() {
			Attribute("id", String, "User ID to delete")

			Required("id")
		})
		Result(OperationStatusResponse)
		HTTP(func() {
			DELETE("/friends/{id}")
		})
	})
})

var Friend = Type("Friend", func() {
	Attribute("id", String, "User ID")
	Attribute("email", String, "Email")
	Attribute("firstName", String, "First name")
	Attribute("lastName", String, "Last name")

	Required("id", "email", "firstName", "lastName")
})

var UserProfileResponse = Type("UserProfileResponse", func() {
	Attribute("id", String, "User ID")
	Attribute("email", String, "Email")
	Attribute("firstName", String, "First name")
	Attribute("lastName", String, "Last name")
	Attribute("friendsList", ArrayOf(Friend), "Friendslist")

	Required("id", "email", "firstName", "lastName", "friendsList")
})
