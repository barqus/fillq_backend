package api

import (
	"fmt"
	"github.com/barqus/fillq_backend/internal/common_http"
	"github.com/barqus/fillq_backend/internal/database"
	"github.com/barqus/fillq_backend/internal/middleware"
	"github.com/barqus/fillq_backend/internal/participants"
	"github.com/barqus/fillq_backend/internal/pickems"
	"github.com/barqus/fillq_backend/internal/questions"
	"github.com/barqus/fillq_backend/internal/users"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func HandlerAPIv1(router chi.Router) {
	// TODO: READ ENV VARIABLES THROUGH .ENV FILE

	l := logrus.New()

	l.SetFormatter(&logrus.JSONFormatter{})
	databaseClient, err := database.Initialize(os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))

	if err != nil {
		l.Info("TEST")
		l.Info(err)
	}
	httpClient := common_http.NewClient(http.DefaultClient)

	usersClient := users.MustNewHttpClient(users.MustNewService(httpClient, users.MustNewStorage(databaseClient)))
	router.Route("/user", func(r chi.Router) {
		r.Get("/login/{id}", usersClient.LoginUser)
		r.Get("/{id}", usersClient.GetUserByID)
	})

	participantsClient := participants.MustNewHttpClient(participants.MustNewService(participants.MustNewStorage(databaseClient)))
	router.Route("/participants", func(r chi.Router) {
		r.Get("/", participantsClient.GetAllParticipants)
		r.With(adminMiddleware).Post("/", participantsClient.AddParticipant)
		r.With(adminMiddleware).Delete("/{id}", participantsClient.DeleteParticipant)
	})

	pickemsClient := pickems.MustNewHttpClient(pickems.MustNewService(pickems.MustNewStorage(databaseClient)))
	router.Route("/pickems", func(r chi.Router) {
		r.Get("/{user_id}", pickemsClient.GetUsersPickems)
		r.Post("/{user_id}", pickemsClient.CreateUsersPickems)
		r.Delete("/{user_id}", pickemsClient.DeleteUsersAllPickems)
	})

	qnaClient := questions.MustNewHttpClient(questions.MustNewService(questions.MustNewStorage(databaseClient)))
	router.Route("/questions", func(r chi.Router) {
		r.Get("/{participant_id}", qnaClient.GetParticipantsQNA)
		r.With(adminMiddleware).Post("/", qnaClient.AddQNA)
		r.With(adminMiddleware).Delete("/{id}", qnaClient.DeleteQNAByID)
	})

}

func adminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("HERE")
		c, err := r.Cookie("jwt_token")
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// For any other type of error, return a bad request status
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		jwtTokenInformation, err := middleware.VerifyToken(c.Value)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		jwtTokenMetadata := jwtTokenInformation.Claims.(jwt.MapClaims)

		userIsAdmin := jwtTokenMetadata["admin"].(bool)
		if userIsAdmin != true {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
