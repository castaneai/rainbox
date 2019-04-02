package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/castaneai/rainbox/pkg/rainbox"
	"github.com/go-chi/chi"
	"gopkg.in/go-playground/validator.v9"
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

type createUserRequest struct {
	DisplayName string `validate:"required,min=2,max=40"`
}

func createUser(userID rainbox.UserID, sv *rainbox.Services, w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	reqBody := &createUserRequest{DisplayName: r.FormValue("displayName")}
	val := validator.New()
	if err := val.Struct(reqBody); err != nil {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	ctx := r.Context()
	if err := sv.User.Register(ctx, userID, reqBody.DisplayName); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
