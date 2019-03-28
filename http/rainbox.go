package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func main() {
	r := newRouter()

	addr := ":8080"
	if p := os.Getenv("PORT"); p != "" {
		addr = fmt.Sprintf(":%s", p)
	}
	log.Printf("Listening on %s...", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}

func newRouter() http.Handler {
	r := chi.NewRouter()
	r.Route("/users", routeUsers)
	return r
}
