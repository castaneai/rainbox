package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/castaneai/rainbox/pkg/rainbox"

	"github.com/go-chi/chi"
)

func usersApi(r chi.Router, v rainbox.Verifier, services *rainbox.Services) {
	r.Use(Authenticator(v))
	r.Get("/me", apiHandler(services, getMyUser))
	r.Post("/", apiHandler(services, createUser))
}

func getMyUser(userID rainbox.UserID, sv *rainbox.Services, w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, err := sv.User.Get(ctx, userID)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	enc := json.NewEncoder(w)
	if err := enc.Encode(user); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func createUser(userID rainbox.UserID, sv *rainbox.Services, w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	displayName := "new user"
	if err := sv.User.Register(ctx, userID, displayName); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
