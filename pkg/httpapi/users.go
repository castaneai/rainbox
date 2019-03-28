package httpapi

import (
	"fmt"
	"net/http"

	"cloud.google.com/go/firestore"

	"github.com/castaneai/rainbox/pkg/rainbox"

	"github.com/castaneai/rainbox/pkg/httpapi/middleware"

	"github.com/go-chi/chi"
)

func usersApi(r chi.Router, v rainbox.Verifier, store *firestore.Client) {
	repo := rainbox.NewFirestoreUserRepository(store)
	us := rainbox.NewUserService(repo)

	// public api
	r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
	})

	// private api
	r.Group(func(r chi.Router) {
		r.Use(middleware.Authenticator(v, us))
		r.Get("/profile", userHandler(profile))
	})
}

func userHandler(f func(*rainbox.User, http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, _ := r.Context().Value(middleware.UserCtxKey).(*rainbox.User)
		if user == nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		f(user, w, r)
	}
}

func profile(user *rainbox.User, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("username: %s", user.DisplayName)))
}
