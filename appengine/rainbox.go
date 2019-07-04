package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"firebase.google.com/go/auth"

	"cloud.google.com/go/firestore"

	"go.uber.org/dig"

	firebase "firebase.google.com/go"
	"github.com/castaneai/rainbox/pkg/rainbox"

	"github.com/castaneai/rainbox/pkg/httpapi"
)

func main() {
	ctx := context.Background()
	c := dig.New()
	if err := setupDIContainer(ctx, c); err != nil {
		log.Fatal(err)
	}

	if err := c.Invoke(func(verifier rainbox.Verifier, services rainbox.Services) error {
		handler := httpapi.NewHandler(verifier, &services)
		addr := ":8080"
		if p := os.Getenv("PORT"); p != "" {
			addr = fmt.Sprintf(":%s", p)
		}
		log.Printf("Listening on %s...", addr)
		return http.ListenAndServe(addr, handler)
	}); err != nil {
		log.Fatal(err)
	}
}

func setupDIContainer(ctx context.Context, c *dig.Container) error {
	if err := c.Provide(func() (*firebase.App, error) {
		return firebase.NewApp(ctx, nil)
	}); err != nil {
		return err
	}
	if err := c.Provide(func(app *firebase.App) (*firestore.Client, error) {
		return app.Firestore(ctx)
	}); err != nil {
		return err
	}
	if err := c.Provide(func(app *firebase.App) (*auth.Client, error) {
		return app.Auth(ctx)
	}); err != nil {
		return err
	}
	if err := c.Provide(func(auth *auth.Client) rainbox.Verifier {
		return rainbox.NewFirebaseAuthVerifier(auth)
	}); err != nil {
		return err
	}

	if err := rainbox.SetupDIContainer(c); err != nil {
		return err
	}
	return nil
}
