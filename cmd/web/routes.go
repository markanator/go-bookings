package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/markanator/go-bookings/pkg/config"
	"github.com/markanator/go-bookings/pkg/handlers"
	"net/http"
)

// define app routes
func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	// define middleware
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	// define application routes
	mux.Get("/", handlers.Repo.HomeHandler)
	mux.Get("/about", handlers.Repo.AboutHandler)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
