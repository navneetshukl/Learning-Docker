package main

import (
	"fmt"
	"go_forms/routes"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	mux := chi.NewRouter()

	mux.Get("/", routes.Home)
	mux.Post("/getdata", routes.Data)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
	fmt.Println("Server started at port 8080")
}
