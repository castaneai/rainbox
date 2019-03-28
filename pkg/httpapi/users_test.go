package httpapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
)

func requestWithUser(req *http.Request, user *clientUser) *http.Request {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", user.IDToken))
	return req
}

func (suite *TestSuite) TestAuth() {
	// unauthorized
	{
		req := httptest.NewRequest("POST", "/users", nil)
		rec := httptest.NewRecorder()
		suite.Handler.ServeHTTP(rec, req)
		assert.Equal(suite.T(), http.StatusUnauthorized, rec.Code)
	}

	// authorized as user1
	{
		{
			req := httptest.NewRequest("POST", "/users", nil)
			req = requestWithUser(req, testClientUser1)
			rec := httptest.NewRecorder()
			suite.Handler.ServeHTTP(rec, req)
			assert.Equal(suite.T(), http.StatusOK, rec.Code)
		}

		{
			req := httptest.NewRequest("GET", "/users/me", nil)
			req = requestWithUser(req, testClientUser1)
			rec := httptest.NewRecorder()
			suite.Handler.ServeHTTP(rec, req)
			assert.Equal(suite.T(), http.StatusOK, rec.Code)
		}
	}
}
