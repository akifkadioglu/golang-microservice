package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"*"},
		AllowedHeaders: []string{"*"},
		MaxAge:         300,
	}))
	mux.Use(middleware.Heartbeat("/ping"))
	mux.Get("/", app.Broker)
	return mux
}
