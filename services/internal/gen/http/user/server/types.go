// Code generated by goa v3.11.3, DO NOT EDIT.
//
// user HTTP server types
//
// Command:
// $ goa gen github.com/JordanRad/chatbook/services/internal/design -o
// ./internal

package server

import (
	user "github.com/JordanRad/chatbook/services/internal/gen/user"
	goa "goa.design/goa/v3/pkg"
)

// RegisterRequestBody is the type of the "user" service "register" endpoint
// HTTP request body.
type RegisterRequestBody struct {
	// User first name
	FirstName *string `form:"firstName,omitempty" json:"firstName,omitempty" xml:"firstName,omitempty"`
	// User last name
	LastName *string `form:"lastName,omitempty" json:"lastName,omitempty" xml:"lastName,omitempty"`
	// User email
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// User password
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
	// Confirmed password of the user
	ConfirmedPassword *string `form:"confirmedPassword,omitempty" json:"confirmedPassword,omitempty" xml:"confirmedPassword,omitempty"`
}

// UpdateProfileNamesRequestBody is the type of the "user" service
// "updateProfileNames" endpoint HTTP request body.
type UpdateProfileNamesRequestBody struct {
	// Updated user first name
	FirstName *string `form:"firstName,omitempty" json:"firstName,omitempty" xml:"firstName,omitempty"`
	// Updated user last name
	LastName *string `form:"lastName,omitempty" json:"lastName,omitempty" xml:"lastName,omitempty"`
}

// RegisterResponseBody is the type of the "user" service "register" endpoint
// HTTP response body.
type RegisterResponseBody struct {
	// Operation status
	Message string `form:"message" json:"message" xml:"message"`
}

// GetUserProfileResponseBody is the type of the "user" service
// "getUserProfile" endpoint HTTP response body.
type GetUserProfileResponseBody struct {
	// User ID
	ID int `form:"id" json:"id" xml:"id"`
	// Email
	Email string `form:"email" json:"email" xml:"email"`
	// First name
	FirstName string `form:"firstName" json:"firstName" xml:"firstName"`
	// Last name
	LastName string `form:"lastName" json:"lastName" xml:"lastName"`
	// Friendslist
	FriendsList []*FriendResponseBody `form:"friendsList" json:"friendsList" xml:"friendsList"`
}

// UpdateProfileNamesResponseBody is the type of the "user" service
// "updateProfileNames" endpoint HTTP response body.
type UpdateProfileNamesResponseBody struct {
	// Operation status
	Message string `form:"message" json:"message" xml:"message"`
}

// RegisterUniqueEmailErrorResponseBody is the type of the "user" service
// "register" endpoint HTTP response body for the "UniqueEmailError" error.
type RegisterUniqueEmailErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// RegisterUnmatchingPassowrdsResponseBody is the type of the "user" service
// "register" endpoint HTTP response body for the "UnmatchingPassowrds" error.
type RegisterUnmatchingPassowrdsResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// FriendResponseBody is used to define fields on response body types.
type FriendResponseBody struct {
	// User ID
	ID int `form:"id" json:"id" xml:"id"`
	// Email
	Email string `form:"email" json:"email" xml:"email"`
	// First name
	FirstName string `form:"firstName" json:"firstName" xml:"firstName"`
	// Last name
	LastName string `form:"lastName" json:"lastName" xml:"lastName"`
}

// NewRegisterResponseBody builds the HTTP response body from the result of the
// "register" endpoint of the "user" service.
func NewRegisterResponseBody(res *user.RegisterResponse) *RegisterResponseBody {
	body := &RegisterResponseBody{
		Message: res.Message,
	}
	return body
}

// NewGetUserProfileResponseBody builds the HTTP response body from the result
// of the "getUserProfile" endpoint of the "user" service.
func NewGetUserProfileResponseBody(res *user.UserProfileResponse) *GetUserProfileResponseBody {
	body := &GetUserProfileResponseBody{
		ID:        res.ID,
		Email:     res.Email,
		FirstName: res.FirstName,
		LastName:  res.LastName,
	}
	if res.FriendsList != nil {
		body.FriendsList = make([]*FriendResponseBody, len(res.FriendsList))
		for i, val := range res.FriendsList {
			body.FriendsList[i] = marshalUserFriendToFriendResponseBody(val)
		}
	}
	return body
}

// NewUpdateProfileNamesResponseBody builds the HTTP response body from the
// result of the "updateProfileNames" endpoint of the "user" service.
func NewUpdateProfileNamesResponseBody(res *user.OperationStatusResponse) *UpdateProfileNamesResponseBody {
	body := &UpdateProfileNamesResponseBody{
		Message: res.Message,
	}
	return body
}

// NewRegisterUniqueEmailErrorResponseBody builds the HTTP response body from
// the result of the "register" endpoint of the "user" service.
func NewRegisterUniqueEmailErrorResponseBody(res *goa.ServiceError) *RegisterUniqueEmailErrorResponseBody {
	body := &RegisterUniqueEmailErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewRegisterUnmatchingPassowrdsResponseBody builds the HTTP response body
// from the result of the "register" endpoint of the "user" service.
func NewRegisterUnmatchingPassowrdsResponseBody(res *goa.ServiceError) *RegisterUnmatchingPassowrdsResponseBody {
	body := &RegisterUnmatchingPassowrdsResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewRegisterPayload builds a user service register endpoint payload.
func NewRegisterPayload(body *RegisterRequestBody) *user.RegisterPayload {
	v := &user.RegisterPayload{
		FirstName:         *body.FirstName,
		LastName:          *body.LastName,
		Email:             *body.Email,
		Password:          *body.Password,
		ConfirmedPassword: *body.ConfirmedPassword,
	}

	return v
}

// NewUpdateProfileNamesPayload builds a user service updateProfileNames
// endpoint payload.
func NewUpdateProfileNamesPayload(body *UpdateProfileNamesRequestBody) *user.UpdateProfileNamesPayload {
	v := &user.UpdateProfileNamesPayload{
		FirstName: *body.FirstName,
		LastName:  *body.LastName,
	}

	return v
}

// ValidateRegisterRequestBody runs the validations defined on
// RegisterRequestBody
func ValidateRegisterRequestBody(body *RegisterRequestBody) (err error) {
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.Password == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("password", "body"))
	}
	if body.ConfirmedPassword == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("confirmedPassword", "body"))
	}
	if body.FirstName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("firstName", "body"))
	}
	if body.LastName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("lastName", "body"))
	}
	return
}

// ValidateUpdateProfileNamesRequestBody runs the validations defined on
// UpdateProfileNamesRequestBody
func ValidateUpdateProfileNamesRequestBody(body *UpdateProfileNamesRequestBody) (err error) {
	if body.FirstName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("firstName", "body"))
	}
	if body.LastName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("lastName", "body"))
	}
	return
}