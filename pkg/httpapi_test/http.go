package httpapi

import (
	"context"
	"fmt"
	"net/http"
	"os"
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
