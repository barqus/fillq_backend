package handler

import (
	"github.com/barqus/fillq_backend/internal/api"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"net/http"
)

func HandleHTTP() {
	r := chi.NewRouter()
	l := logrus.New()

	l.SetFormatter(&logrus.JSONFormatter{})

	r.Route("/api/", func(r chi.Router) {
		r.Route("/v1", api.HandlerAPIv1)
	})
	l.Info("SERVER STARTED...")
	panic(http.ListenAndServe(":8080",r))

}