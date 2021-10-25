package leagueoflegends
//RGAPI-d7047eaa-db5e-47a7-867f-8b860b6b441f

import (
	"encoding/json"
	"fmt"
	"github.com/KnutZuidema/golio"
	"github.com/KnutZuidema/golio/api"
	"github.com/gorilla/mux"
	"net/http"
)

func GetParticipantsLeagueInformation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	summonerName := vars["summoner_name"]

	apiKey := "RGAPI-d7047eaa-db5e-47a7-867f-8b860b6b441f"
	client := golio.NewClient(apiKey, golio.WithRegion(api.RegionEuropeWest))
	summoner, err := client.Riot.LoL.Summoner.GetByName(summonerName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	SummonersLeagueInformation, err := client.Riot.LoL.League.ListBySummoner(summoner.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(SummonersLeagueInformation)
}

func GetParticipantsMatchInfo(summonerId string) {
	apiKey := "RGAPI-d7047eaa-db5e-47a7-867f-8b860b6b441f"
	client := golio.NewClient(apiKey, golio.WithRegion(api.RegionEuropeWest))
	//summoner, _ := client.Riot.Summoner.GetByID(summonerId)
	summoner, _ := client.Riot.Summoner.GetByName(summonerId)
	//fmt.Printf("Accoun id: %s \n", summoner.RevisionDate)
	//matchList, err := client.Riot.LoL.Spectator.GetCurrent(summoner.ID)

	// grazina lygos informacija wins loses etc.
	//test, _ := client.Riot.LoL.League.ListBySummoner(summoner.ID)

	c := client.Riot.LoL.Match.ListStream(summoner.AccountID)

	output := <- c

	fmt.Print(output)
	//err := nill
	//if err != nil {
	//	fmt.Printf("ERROR")
	//	//fmt.Printf(err.Error())
	//}

	//fmt.Print(matchList)
	fmt.Printf("%s is a level %d summoner\n", summoner.Name, summoner.SummonerLevel)
	//champion, _ := client.DataDragon.GetChampion("Ashe")
	//mastery, err := client.Riot.ChampionMastery.Get(summoner.ID, champion.Key)
	//if err != nil {
	//	fmt.Printf("%s has not played any games on %s\n", summoner.Name, champion.Name)
	//} else {
	//	fmt.Printf("%s has mastery level %d with %d points on %s\n", summoner.Name, mastery.ChampionLevel,
	//		mastery.ChampionPoints, champion.Name)
	//}
}