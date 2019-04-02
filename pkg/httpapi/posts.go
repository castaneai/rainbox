package httpapi

import (
	"net/http"

	"github.com/unrolled/render"

	"github.com/castaneai/rainbox/pkg/rainbox"
	"github.com/go-chi/chi"
)

func postsAPI(r chi.Router, v rainbox.Verifier, services *rainbox.Services) {
	r.Use(Authenticator(v))
	r.Get("/count", apiHandler(services, countAllPosts))
}

func countAllPosts(userID rainbox.UserID, sv *rainbox.Services, w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	cnt, err := sv.Post.CountAll(ctx)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	ren := render.New()
	if err := ren.JSON(w, http.StatusOK, map[string]interface{}{
		"count": cnt,
	}); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
