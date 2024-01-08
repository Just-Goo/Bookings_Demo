package main

import (
	"net/http"

	"github.com/Just-Goo/Bookings_Demo/pkg/config"
	"github.com/Just-Goo/Bookings_Demo/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {

	// Using 'Chi' router
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer) // Adding a middleware
	mux.Use(NoSurf) // Using a custom middleware
	mux.Use(SessionLoad) // Using a custom middleware

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
