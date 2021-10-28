package api

import (
	"github.com/barqus/fillq_backend/internal/database"
	"github.com/barqus/fillq_backend/internal/participants"
	"github.com/go-chi/chi"
)

func HandlerAPIv1(router chi.Router) {
	//httpClient := common_http.NewClient(http.DefaultClient)
	//dbUser, dbPassword, dbName :=
	//os.Getenv("POSTGRES_USER"),
	//os.Getenv("POSTGRES_PASSWORD"),
	//os.Getenv("POSTGRES_DB")

	databaseClient, _ := database.Initialize("barqus", "root", "fillq-db")
	participantClient := participants.MustNewHttpClient(participants.MustNewService(participants.MustNewStorage(databaseClient)))
	//database := database.Get
	router.Route("/participants", func(r chi.Router) {
		r.Get("/", participantClient.GetAllParticipants)
	})

	//router := mux.NewRouter()
	//
	//router.HandleFunc("/participants", participants.GetAllParticipants).Methods(http.MethodGet)
	//router.HandleFunc("/participants", participants.AddParticipant).Methods(http.MethodPost)
	//
	//
	//router.HandleFunc("/participants/{summoner_name}", leagueoflegends.GetParticipantsLeagueInformation).Methods(http.MethodGet)
	//
	//log.Println("API is running!")
	//http.ListenAndServe(":4000", router)
}
