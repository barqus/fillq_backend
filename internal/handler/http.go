package handler

import (
	"github.com/barqus/fillq_backend/internal/api"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Cors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=ascii")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers","Content-Type,access-control-allow-origin, access-control-allow-headers")
	w.Write([]byte("Hello, World!"))
}
func HandleHTTP() {
	r := chi.NewRouter()

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing

	l := logrus.New()

	l.SetFormatter(&logrus.JSONFormatter{})
	r.Route("/api/", func(r chi.Router) {
		r.Route("/v1", api.HandlerAPIv1)
	})
	l.Info("SERVER STARTED...")
	panic(http.ListenAndServe(":8080",r))

}
