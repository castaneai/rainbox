package httpapi

import (
	"net/http"

	"github.com/castaneai/rainbox/pkg/rainbox"

	"github.com/go-chi/chi"
)

func NewHandler(v rainbox.Verifier, services *rainbox.Services) http.Handler {
	r := chi.NewRouter()
	r.Route("/users", routeGroup(v, services, usersApi))
	return r
}

func routeGroup(v rainbox.Verifier, services *rainbox.Services, f func(chi.Router, rainbox.Verifier, *rainbox.Services)) func(chi.Router) {
	return func(r chi.Router) {
		f(r, v, services)
	}
}

func apiHandler(services *rainbox.Services, f func(rainbox.UserID, *rainbox.Services, http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, _ := r.Context().Value(userIDCtxKey).(rainbox.UserID)
		if userID == rainbox.InvalidUserID {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		f(userID, services, w, r)
	}
}
