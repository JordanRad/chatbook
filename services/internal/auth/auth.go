package auth

import (
	"context"
	"errors"
)

var ContextKeyUser string = "user"

func UserInContext(ctx context.Context) (*User, error) {
	u, ok := ctx.Value(ContextKeyUser).(*User)
	if !ok {
		return nil, errors.New("cannot convert context key to User struct, please check if this route is included in the middleware's protected routes")
	}

	return u, nil
}
