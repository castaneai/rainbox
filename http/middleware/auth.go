package middleware

import (
	"context"
	"net/http"
	"strings"

	rainbox "github.com/castaneai/rainbox/pkg"

	"firebase.google.com/go/auth"
)

var UserCtxKey = &contextKey{"User"}

// contextKey is a value for use with context.WithValue. It's used as
// a pointer so it fits in an interface{} without allocation. This technique
// for defining context keys was copied from Go 1.7's new use of context in net/http.
type contextKey struct {
	name string
}

func (k *contextKey) String() string {
	return k.name
}

func Authenticator(ac *auth.Client, us *rainbox.UserService) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			idToken := idTokenFromHeader(r)
			token, err := ac.VerifyIDToken(ctx, idToken)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			if token == nil {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			user, err := us.SignIn(ctx, token.UID)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			ctx = newContext(ctx, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func newContext(ctx context.Context, user *rainbox.User) context.Context {
	ctx = context.WithValue(ctx, UserCtxKey, user)
	return ctx
}

func idTokenFromHeader(r *http.Request) string {
	bearer := r.Header.Get("Authorization")
	if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
		return bearer[7:]
	}
	return ""
}
