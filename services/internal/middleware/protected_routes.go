package middleware

type ProtectedRoute struct {
	HTTPMethod string
	URL        string
}

// TODO(JordanRad): use regex instead
var protectedRoutes = []ProtectedRoute{
	{
		HTTPMethod: "GET",
		URL:        "/users",
	},
	{
		HTTPMethod: "PUT",
		URL:        "/users/profile",
	},
}
