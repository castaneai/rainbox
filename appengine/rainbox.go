package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/castaneai/rainbox/pkg/httpapi"
)

func main() {
	handler := httpapi.NewHandler()

	addr := ":8080"
	if p := os.Getenv("PORT"); p != "" {
		addr = fmt.Sprintf(":%s", p)
	}
	log.Printf("Listening on %s...", addr)
	log.Fatal(http.ListenAndServe(addr, handler))
}
