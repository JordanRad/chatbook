package design

import (
	. "goa.design/goa/v3/dsl"
)

// ------ POST login/: Handles user login and authentication.
// ------ POST register/: Handles user registration and account creation.
// POST friend/: Adds a user to the friend list.
// DELETE friend/: Removes a user from the friend list.
// ------ GET profile: Retrieves user profile information, including the list of friends.
// ------ PUT profile/: Allows modification of the user's profile, such as updating the name.

var _ = Service("user", func() {
	Description("User service is responsible for handling user data and requests")
	HTTP(func() {
		Path("/api/v1/users") // Prefix to HTTP path of all requests.
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

	Method("getUserProfile", func() {
		Result(UserProfileResponse)
		HTTP(func() {
			GET("/")
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
})

var Friend = Type("Friend", func() {
	Attribute("id", Int, "User ID")
	Attribute("email", String, "Email")
	Attribute("firstName", String, "First name")
	Attribute("lastName", String, "Last name")

	Required("id", "email", "firstName", "lastName")
})

var UserProfileResponse = Type("UserProfileResponse", func() {
	Attribute("id", Int, "User ID")
	Attribute("email", String, "Email")
	Attribute("firstName", String, "First name")
	Attribute("lastName", String, "Last name")
	Attribute("friendsList", ArrayOf(Friend), "Friendslist")

	Required("id", "email", "firstName", "lastName", "friendsList")
})
