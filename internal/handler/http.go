package handler

import (
	"github.com/barqus/fillq_backend/internal/api"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func HandleHTTP() {
	r := chi.NewRouter()

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	l := logrus.New()

	l.SetFormatter(&logrus.JSONFormatter{})
	r.Route("/api/", func(r chi.Router) {
		r.Route("/v1", api.HandlerAPIv1)
	})
	logrus.Info("START")
	l.Info(os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	l.Info(os.Environ())
	l.Info("TEST")
	l.Info("SERVER STARTED...")
	panic(http.ListenAndServe(":8080",r))

}
