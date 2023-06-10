// Code generated by goa v3.11.3, DO NOT EDIT.
//
// HTTP request path constructors for the user service.
//
// Command:
// $ goa gen github.com/JordanRad/chatbook/services/internal/design -o
// ./internal

package server

import (
	"fmt"
)

// RegisterUserPath returns the URL path to the user service register HTTP endpoint.
func RegisterUserPath() string {
	return "/api/user-management/v1/users/register"
}

// GetProfileUserPath returns the URL path to the user service getProfile HTTP endpoint.
func GetProfileUserPath() string {
	return "/api/user-management/v1/users/profile"
}

// UpdateProfileNamesUserPath returns the URL path to the user service updateProfileNames HTTP endpoint.
func UpdateProfileNamesUserPath() string {
	return "/api/user-management/v1/users/profile"
}

// AddFriendUserPath returns the URL path to the user service addFriend HTTP endpoint.
func AddFriendUserPath(id string) string {
	return fmt.Sprintf("/api/user-management/v1/users/friend/%v", id)
}

// RemoveFriendUserPath returns the URL path to the user service removeFriend HTTP endpoint.
func RemoveFriendUserPath(id string) string {
	return fmt.Sprintf("/api/user-management/v1/users/friends/%v", id)
}
