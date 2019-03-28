package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserProfile(t *testing.T) {
	r := newRouter()
	ts := httptest.NewServer(r)
	defer ts.Close()

	req := httptest.NewRequest("GET", "/users/profile", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}
