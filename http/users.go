package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	firebase "firebase.google.com/go"

	"github.com/castaneai/rainbox/pkg"

	"github.com/castaneai/rainbox/http/middleware"

	"github.com/go-chi/chi"
)

func routeUsers(r chi.Router) {
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	auth, err := app.Auth(ctx)
	if err != nil {
		log.Fatal(err)
	}
	store, err := app.Firestore(ctx)
	if err != nil {
		log.Fatal(err)
	}
	repo := pkg.NewFirestoreUserRepository(store)
	us := pkg.NewUserService(repo)

	// public api
	r.Post("/sign-in", signIn)

	// private api
	r.Group(func(r chi.Router) {
		r.Use(middleware.Authenticator(auth, us))
		r.Get("/profile", userHandler(profile))
	})
}

func userHandler(f func(*pkg.User, http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, _ := r.Context().Value(middleware.UserCtxKey).(*pkg.User)
		f(user, w, r)
	}
}

func signIn(w http.ResponseWriter, r *http.Request) {
	// TODO:
}

func profile(user *pkg.User, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("username: %s", user.DisplayName)))
}
