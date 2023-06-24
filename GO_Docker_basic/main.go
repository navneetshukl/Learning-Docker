package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	mux := chi.NewRouter()

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is home page"))
	})
	mux.Get("/about", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is About page"))
	})

	http.ListenAndServe(":8080", mux)
}
