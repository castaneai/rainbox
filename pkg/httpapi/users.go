package httpapi

import (
	"encoding/json"
	"net/http"

	"cloud.google.com/go/firestore"

	"github.com/castaneai/rainbox/pkg/rainbox"

	"github.com/go-chi/chi"
)

func usersApi(r chi.Router, v rainbox.Verifier, store *firestore.Client) {
	repo := rainbox.NewFirestoreUserRepository(store)
	us := rainbox.NewUserService(repo)

	r.Use(Authenticator(v))
	r.Get("/me", userHandler(us, getMyUser))
	r.Post("/", userHandler(us, createUser))
}

func userHandler(us *rainbox.UserService, f func(rainbox.UserID, *rainbox.UserService, http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, _ := r.Context().Value(userIDCtxKey).(rainbox.UserID)
		if userID == rainbox.InvalidUserID {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		f(userID, us, w, r)
	}
}

func getMyUser(userID rainbox.UserID, us *rainbox.UserService, w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, err := us.Get(ctx, userID)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	enc := json.NewEncoder(w)
	if err := enc.Encode(user);err != nil{
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func createUser(userID rainbox.UserID, us *rainbox.UserService, w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	displayName := "new user"
	if err := us.Register(ctx, userID, displayName);err != nil{
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
