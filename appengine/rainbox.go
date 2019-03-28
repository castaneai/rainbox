package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	firebase "firebase.google.com/go"
	"github.com/castaneai/rainbox/pkg/rainbox"

	"github.com/castaneai/rainbox/pkg/httpapi"
)

func main() {
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	store, err := app.Firestore(ctx)
	if err != nil {
		log.Fatal(err)
	}
	auth, err := app.Auth(ctx)
	if err != nil {
		log.Fatal(err)
	}
	verifier := rainbox.NewFirebaseAuthVerifier(auth)

	handler := httpapi.NewHandler(verifier, store)

	addr := ":8080"
	if p := os.Getenv("PORT"); p != "" {
		addr = fmt.Sprintf(":%s", p)
	}
	log.Printf("Listening on %s...", addr)
	log.Fatal(http.ListenAndServe(addr, handler))
}
