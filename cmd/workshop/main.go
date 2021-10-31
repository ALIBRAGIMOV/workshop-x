package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"net/http"
	"os"
	"workshop/internal/config"
	"workshop/internal/handler"
)

func main() {
	cfg := config.Server{}
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	h := handler.NewHandler()
	r := chi.NewRouter()

	r.Get("/hello", h.Hello)

	path := cfg.Host + ":" + cfg.Port

	log.Printf("starting server at %s", path)
	err = http.ListenAndServe(":" + os.Getenv("PORT"), r)
	log.Printf("shutting server down with %v", err)
}
