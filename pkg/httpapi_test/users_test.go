package httpapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func requestWithUser(req *http.Request, user *User) *http.Request {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", user.IDToken))
	return req
}

func TestAuth(t *testing.T) {
	u := newUser()
	v := newStaticVerifier([]*User{u})

	runWithHandler(t, v, func(t *testing.T, h http.Handler) {
		// unauthorized
		{
			req := httptest.NewRequest("POST", "/users", nil)
			rec := httptest.NewRecorder()
			h.ServeHTTP(rec, req)
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
		}

		// authorized as user1
		{
			req := httptest.NewRequest("POST", "/users", nil)
			req = requestWithUser(req, u)
			rec := httptest.NewRecorder()
			h.ServeHTTP(rec, req)
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}
