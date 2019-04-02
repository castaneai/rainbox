package httpapi

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	user1 := newUser()
	v := newStaticVerifier([]*User{user1})

	runWithHandler(t, v, func(t *testing.T, h http.Handler) {
		// unauthorized
		{
			res := tryRequest(h, "POST", "/users", nil, nil)
			assert.Equal(t, http.StatusUnauthorized, res.Code)
		}

		// create user record with auth
		{
			res := tryRequest(h, "POST", "/users", user1, nil)
			assert.Equal(t, http.StatusOK, res.Code)
		}
	})
}
