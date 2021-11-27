package leagueoflegends

//RGAPI-d7047eaa-db5e-47a7-867f-8b860b6b441f

import (
	//"encoding/json"
	"fmt"
	"github.com/KnutZuidema/golio"
	"github.com/KnutZuidema/golio/api"
	"github.com/KnutZuidema/golio/riot/lol"
	"github.com/barqus/fillq_backend/internal/database"
	"github.com/barqus/fillq_backend/internal/participants"
	"os"
)

// TODO: atsimint kad riot games api limitina requestus turbut reikia daryt scheduleri kuris callins del informacijos
// TODO: update rows instead of add
func UpdateSummonersInformation(databaseClient *database.Database, client *golio.Client) {
	participantsStorage := participants.MustNewStorage(databaseClient)
	lolStorage := MustNewStorage(databaseClient)
	fmt.Println("TRIGGERED UPDATE")
	allParticipants, err := participantsStorage.GetAllParticipants()
	if err != nil {
		fmt.Println("HERE")
		fmt.Println(err.Error())

	}
	for _, item := range allParticipants {
		summoner, err := client.Riot.LoL.Summoner.GetByName(item.SummonerName)
		if err != nil {
			// TODO: ADD ERORR HERE
			//common_http.WriteErrorResponse(w, err)
			//common_http.WriteJSONResponse(w, http.StatusOK, allSummoners)
			fmt.Println("HERE2")
			fmt.Println(err.Error())
			fmt.Println(item.SummonerName)
			continue
		}

		summonersLeagueInformation, err := client.Riot.LoL.League.ListBySummoner(summoner.ID)
		//summonersLeagueInformation
		summonerLeagueFound := false
		for _, leagueItem := range summonersLeagueInformation {
			if leagueItem.QueueType == "RANKED_SOLO_5x5" {
				currentParticipantsSummonerLeague := SummonerLeague{
					PUUID:         summoner.PUUID,
					SummonerName:  leagueItem.SummonerName,
					Tier:          &leagueItem.Tier,
					Rank:          &leagueItem.Rank,
					Points:        &leagueItem.LeaguePoints,
					Wins:          leagueItem.Wins,
					Losses:        leagueItem.Losses,
					ParticipantID: item.Id,
				}

				summonerLeagueFound = true
				summonerExists, err := lolStorage.summonerAlreadyExists(currentParticipantsSummonerLeague.PUUID)
				if err != nil {
					fmt.Println("HERE3")
					fmt.Println(err)
					continue
				}

				if *summonerExists == true {
					err = lolStorage.UpdateSummonerLeagueByID(currentParticipantsSummonerLeague, summoner.PUUID)
					if err != nil {
						fmt.Println("HERE4")
						fmt.Println(err)
						continue
					}
				} else {
					err = lolStorage.AddSummonerLeague(currentParticipantsSummonerLeague)
					if err != nil {
						fmt.Println("HERE5")
						fmt.Println(err)
						continue
					}
				}
				continue
			}
		}

		if summonerLeagueFound == true {
			continue
		}

		wins, loses, err := GetSummonersMatchHistory(client, summoner.PUUID)
		currentParticipantsSummonerLeague := SummonerLeague{
			PUUID:         summoner.PUUID,
			SummonerName:  summoner.Name,
			Tier:          nil,
			Rank:          nil,
			Points:        nil,
			Wins:          wins,
			Losses:        loses,
			ParticipantID: item.Id,
		}
		summonerExists, err := lolStorage.summonerAlreadyExists(currentParticipantsSummonerLeague.PUUID)

		if err != nil {
			fmt.Println("HERE6")
			fmt.Println(err.Error())
			continue
		}

		if *summonerExists == true {
			err = lolStorage.UpdateSummonerLeagueByID(currentParticipantsSummonerLeague, summoner.PUUID)
			if err != nil {
				fmt.Println("HERE7")
				fmt.Println(err)
				continue
			}
		} else {
			err = lolStorage.AddSummonerLeague(currentParticipantsSummonerLeague)
			if err != nil {
				fmt.Println("HERE8")
				fmt.Println(err)
				continue
			}
		}
	}
}

func SetupRiotClient() *golio.Client {
	apiKey := os.Getenv("RIOT_GAMES_API_KEY")
	fmt.Println(apiKey)
	client := golio.NewClient(apiKey, golio.WithRegion(api.RegionEuropeWest))

	return client
}

func GetSummonersMatchHistory(client *golio.Client, PUUID string) (int, int, error) {
	queue := 420
	matchListOptions := lol.MatchListOptions{Queue: &queue}
	arrayOfMatches, _ := client.Riot.LoL.Match.List(PUUID, 0, 10, &matchListOptions)
	// TODO: ERROR CATCHING
	wins := 0
	losses := 0
	for _, item := range arrayOfMatches {
		matchesInfo, _ := client.Riot.LoL.Match.Get(item)
		for _, matchParticipant := range matchesInfo.Info.Participants {
			if matchParticipant.PUUID == PUUID {
				if matchParticipant.Win == true {
					wins++
				} else {
					losses++
				}
			}
		}
	}

	return wins, losses, nil
}
func getLeagueInformationBySummonerName() {
	// TODO: REDO THIS SO IT IS CALLED AS SCHEDULED TASK
	//vars := mux.Vars(r)
	//summonerName := vars["summoner_name"]

	apiKey := os.Getenv("RIOT_GAMES_API_KEY")
	client := golio.NewClient(apiKey, golio.WithRegion(api.RegionEuropeWest))

	//summoners := [12]string{"Fill to Heaven", "Lenkijos Princas", "Old Man Raizzo", "57864393", "jaunutis1111", "elosanta fill", "DNA Domas", "Piridacas V2", "Gurklys FILLQ", "PoveikioKaralius", "urbis šliurbis"}
	summoners := [1]string{"saulius1"}

	allSummoners := make([]*lol.Summoner, 0)
	for _, item := range summoners {
		summoner, err := client.Riot.LoL.Summoner.GetByName(item)
		if err != nil {
			//common_http.WriteErrorResponse(w, err)
			//common_http.WriteJSONResponse(w, http.StatusOK, allSummoners)

			return
		}
		//SummonersLeagueInformation, err := client.Riot.LoL.League.ListBySummoner(summoner.ID)
		//common_http.WriteJSONResponse(w, http.StatusOK, SummonersLeagueInformation)

		allSummoners = append(allSummoners, summoner)
		fmt.Println(summoner)
	}

	// /lol/match/v5/matches/by-puuid/{puuid}/ids?queue=420&start=0&count=11
	queue := 420
	matchListOptions := lol.MatchListOptions{Queue: &queue}
	arrayOfMatches, _ := client.Riot.LoL.Match.List(allSummoners[0].PUUID, 0, 10, &matchListOptions)
	tmp := make([]*lol.Match, 0)
	for _, item := range arrayOfMatches {
		matchInfo, _ := client.Riot.LoL.Match.Get(item)
		tmp = append(tmp, matchInfo)
	}
	fmt.Println(tmp)
	//common_http.WriteJSONResponse(w, http.StatusOK, tmp)

	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//	return
	//}
	//
	//SummonersLeagueInformation, err := client.Riot.LoL.League.ListBySummoner(summoner.ID)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//	return
	//}
	//
	//w.Header().Add("Content-Type", "application/json")
	//w.WriteHeader(http.StatusOK)
	//json.NewEncoder(w).Encode(SummonersLeagueInformation)
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

	output := <-c

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
