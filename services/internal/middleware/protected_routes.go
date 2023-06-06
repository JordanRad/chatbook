package middleware

type ProtectedRoute struct {
	HTTPMethod   string
	URL          string
	AllowedRoles []string
}

// TODO(JordanRad): use regex instead
var protectedRoutes = []ProtectedRoute{
	{
		HTTPMethod: "GET",
		URL:        "/users",
	},
	{
		HTTPMethod:   "PUT",
		URL:          "/users/settings",
		AllowedRoles: []string{"sfit_user", "sfit_coach"},
	},
	{
		HTTPMethod:   "PUT",
		URL:          "/users/password",
		AllowedRoles: []string{"sfit_user", "sfit_coach"},
	},
	{
		HTTPMethod:   "PUT",
		URL:          "/users/profile",
		AllowedRoles: []string{"sfit_user", "sfit_coach"},
	},
	{
		HTTPMethod:   "DELETE",
		URL:          "/users",
		AllowedRoles: []string{"sfit_user", "sfit_coach"},
	},
	{
		HTTPMethod:   "POST",
		URL:          "/body-metrics/records",
		AllowedRoles: []string{"sfit_user", "sfit_coach"},
	},
	{
		HTTPMethod:   "PUT",
		URL:          "/body-metrics/records",
		AllowedRoles: []string{"sfit_user", "sfit_coach"},
	},
	{
		HTTPMethod:   "GET",
		URL:          "/body-metrics/records/statistics",
		AllowedRoles: []string{"sfit_user", "sfit_coach"},
	},
	{
		HTTPMethod:   "POST",
		URL:          "/plans",
		AllowedRoles: []string{"sfit_user", "sfit_coach"},
	},
	{
		HTTPMethod:   "GET",
		URL:          "/plans",
		AllowedRoles: []string{"sfit_user", "sfit_coach"},
	},
	{
		HTTPMethod:   "PUT",
		URL:          "/plans",
		AllowedRoles: []string{"sfit_user", "sfit_coach"},
	},
	{
		HTTPMethod:   "DELETE",
		URL:          "/plans",
		AllowedRoles: []string{"sfit_user", "sfit_coach"},
	},
	{
		HTTPMethod:   "GET",
		URL:          "/coaches/trainees",
		AllowedRoles: []string{"sfit_coach"},
	},
	{
		HTTPMethod:   "POST",
		URL:          "/users/profile-picture",
		AllowedRoles: []string{"sfit_user", "sfit_coach"},
	},
	{
		HTTPMethod:   "GET",
		URL:          "/workouts",
		AllowedRoles: []string{"sfit_user", "sfit_coach"},
	},
	{
		HTTPMethod:   "POST",
		URL:          "/workouts",
		AllowedRoles: []string{"sfit_user", "sfit_coach"},
	},
	{
		HTTPMethod: "PUT",
		URL:        "/workouts",
		AllowedRoles: []string {"sfit_user", "sfit_coach"},
	},
	{
		HTTPMethod: "DELETE",
		URL:        "/workouts",
		AllowedRoles: []string {"sfit_user", "sfit_coach"},
	},
}
