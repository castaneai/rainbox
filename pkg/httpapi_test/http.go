package httpapi

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"

	"github.com/castaneai/rainbox/pkg/httpapi"
	"github.com/castaneai/rainbox/pkg/rainbox"

	"go.uber.org/dig"
)

func runWithHandler(t *testing.T, v rainbox.Verifier, f func(*testing.T, http.Handler)) {
	c := dig.New()
	ctx := context.Background()
	if err := setupDIContainerForTesting(ctx, c, v); err != nil {
		t.Fatal(err)
	}
	if err := c.Invoke(func(v rainbox.Verifier, svc rainbox.Services) {
		h := httpapi.NewHandler(v, &svc)
		f(t, h)
	}); err != nil {
		t.Fatal(err)
	}
}

func tryRequest(h http.Handler, method, path string, authUser *User, reqBody map[string]interface{}) *httptest.ResponseRecorder {
	var body io.Reader
	if reqBody != nil {
		form := url.Values{}
		for k, v := range reqBody {
			form.Add(k, fmt.Sprintf("%s", v))
		}
		body = strings.NewReader(form.Encode())
	}
	req := httptest.NewRequest(method, path, body)
	if strings.ToUpper(method) == "POST" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	if authUser != nil {
		req = requestWithUser(req, authUser)
	}
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	return rec
}

func requestWithUser(req *http.Request, user *User) *http.Request {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", user.IDToken))
	return req
}

func setupDIContainerForTesting(ctx context.Context, c *dig.Container, v rainbox.Verifier) error {
	if err := c.Provide(func() (*firebase.App, error) {
		return firebase.NewApp(ctx, nil)
	}); err != nil {
		return err
	}
	if err := c.Provide(func(app *firebase.App) (*firestore.Client, error) {
		if e := os.Getenv("FIRESTORE_EMULATOR_HOST"); e == "" {
			return nil, fmt.Errorf("env not set: FIRESTORE_EMULATOR_HOST")
		}
		return app.Firestore(ctx)
	}); err != nil {
		return err
	}
	if err := c.Provide(func() rainbox.Verifier {
		return v
	}); err != nil {
		return err
	}

	if err := rainbox.SetupDIContainer(c); err != nil {
		return err
	}
	return nil
}
