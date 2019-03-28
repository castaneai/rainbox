package httpapi

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func requestWithUser(req *http.Request, user *clientUser) *http.Request {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", user.IDToken))
	return req
}

func TestAuth(t *testing.T) {
	ctx := context.Background()
	user1 := &clientUser{ID: newRandomUserID(), IDToken: "token1"}
	user2 := &clientUser{ID: newRandomUserID(), IDToken: "token2"}
	v := NewTestClientVerifier([]*clientUser{user1, user2})
	store, err := newTestFirestore(ctx)
	if err != nil {
		t.Fatal(err)
	}

	h := NewHandler(v, store)
	ts := httptest.NewServer(h)
	defer ts.Close()

	// unauthorized
	{
		req := httptest.NewRequest("POST", "/users", nil)
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	}

	// authorized as user1
	{
		{
			req := httptest.NewRequest("POST", "/users", nil)
			req = requestWithUser(req, user1)
			rec := httptest.NewRecorder()
			h.ServeHTTP(rec, req)
			assert.Equal(t, http.StatusOK, rec.Code)
		}

		{
			req := httptest.NewRequest("GET", "/users/me", nil)
			req = requestWithUser(req, user1)
			rec := httptest.NewRecorder()
			h.ServeHTTP(rec, req)
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	}

}