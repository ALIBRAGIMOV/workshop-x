package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"workshop/internal/handler"
)

func main()  {
	h := handler.NewHandler()
	r := chi.NewRouter()

	r.Get("/hello", h.Hello)

	log.Print("starting server")
	err := http.ListenAndServe(":8080", r)
	log.Print("shutting server down")
	log.Fatal(err)

}
