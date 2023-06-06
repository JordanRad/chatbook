package middleware

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/JordanRad/chatbook/services/internal/auth"
	"github.com/JordanRad/chatbook/services/internal/auth/jwt"
)

func isRouteProtected(method string, URL *url.URL) bool {
	for _, route := range protectedRoutes {
		if route.HTTPMethod == method && strings.Contains(URL.Path, route.URL) {
			return true
		}
	}
	return false
}

func sliceContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func AuthenticateRequest(us auth.UserStore, jwt jwt.JWTClient) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Check if the route is protected
			isProtectedRoute := isRouteProtected(r.Method, r.URL)

			// Check if the JWT is valid
			if isProtectedRoute {
				authorizationHeader := r.Header.Get("Authorization")
				if len(authorizationHeader) < 1 {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("401 - Unauthorized. \nNo Authorization header is present"))
					return
				}
				// Strip the 'Bearer ' prefix
				tokenString := authorizationHeader[7:]

				_, err := jwt.ValidateJWT(tokenString)
				if err != nil {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("401 - Unauthorized. \nYour token has either expired or is invalid"))
					return
				}

				c, err := jwt.ExtractJWTCLaims(tokenString)
				if err != nil {
					fmt.Println(err.Error())
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("500 - Middleware error. \n"))
					return
				}

				u := &auth.User{
					ID:    c.UserID,
					Email: c.Email,
				}

				ctx := context.Background()
				dbu, err := us.GetUserByEmail(ctx, u.Email)
				if err != nil {
					fmt.Println(err.Error())
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("500 - Middleware error, cannot retrieve user from registry \n"))
					return
				}

				u.Password = dbu.Password
				u.FirstName = dbu.FirstName
				u.LastName = dbu.LastName

				// Inject user in the context
				ctx = context.WithValue(r.Context(), auth.ContextKeyUser, u)
				requestWithContext := r.WithContext(ctx)

				h.ServeHTTP(w, requestWithContext)
				return
			} else {
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte("403 - Forbidden: \nUser does not have access to this route."))
			}

		})
	}
}
