package httpapi

import (
	"github.com/castaneai/rainbox/pkg/rainbox"
	"github.com/go-chi/chi"
)

func postsAPI(r chi.Router, v rainbox.Verifier, services *rainbox.Services) {
	r.Use(Authenticator(v))
}
