package httpapi

import (
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/castaneai/rainbox/pkg/rainbox"

	"github.com/go-chi/chi"
)

func NewHandler(v rainbox.Verifier, store *firestore.Client) http.Handler {
	r := chi.NewRouter()
	r.Route("/users", apiWithDeps(v, store, usersApi))
	return r
}

func apiWithDeps(v rainbox.Verifier, store *firestore.Client, f func(chi.Router, rainbox.Verifier, *firestore.Client)) func(chi.Router) {
	return func(r chi.Router) {
		f(r, v, store)
	}
}
