// cmd/api/main.go
package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"featurn/internal/api"
)

func main() {
	r := chi.NewRouter()
	api.RegisterRoutes(r)

	log.Println("🚀 Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
