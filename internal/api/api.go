package api

import (
	"github.com/barqus/fillq_backend/internal/participants"
	"github.com/go-chi/chi"
)

func HandlerAPIv1(router chi.Router) {

	router.Route("/participants", func(r chi.Router) {
		r.Get("/", participants.GetAllParticipants)
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