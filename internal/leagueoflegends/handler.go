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
	callAmount := 0
	participantsStorage := participants.MustNewStorage(databaseClient)
	lolStorage := MustNewStorage(databaseClient)
	allParticipants, err := participantsStorage.GetAllParticipants()
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, item := range allParticipants {
		fmt.Println("UPDATING: ", item.Nickname)
		summoner, err := client.Riot.LoL.Summoner.GetByName(item.SummonerName)
		callAmount++

		fmt.Println("CALL AMOUNT:", callAmount)
		if err != nil {
			fmt.Println("ERR", err)
			continue
		}

		summonersLeagueInformation, err := client.Riot.LoL.League.ListBySummoner(summoner.ID)
		if err != nil {
			fmt.Println("ERR", err)
			continue
		}
		callAmount++
		fmt.Println("CALL AMOUNT:", callAmount)
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
					fmt.Println(err)
					continue
				}

				if *summonerExists == true {
					err = lolStorage.UpdateSummonerLeagueByID(currentParticipantsSummonerLeague, summoner.PUUID)
					if err != nil {
						fmt.Println(err)
						continue
					}
				} else {
					err = lolStorage.AddSummonerLeague(currentParticipantsSummonerLeague)
					if err != nil {
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

		wins, loses, err := GetSummonersMatchHistory(client, summoner.PUUID,&callAmount)
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

func GetSummonersMatchHistory(client *golio.Client, PUUID string, callAmount *int) (int, int, error) {
	queue := 420
	matchListOptions := lol.MatchListOptions{Queue: &queue}
	*callAmount++
	fmt.Println("CALL AMOUNT:", *callAmount)
	arrayOfMatches, err := client.Riot.LoL.Match.List(PUUID, 0, 10, &matchListOptions)
	if err != nil {
		fmt.Println("ERR", err)
		return 0, 0, err
	}
	wins := 0
	losses := 0
	for _, item := range arrayOfMatches {
		*callAmount++
		fmt.Println("CALL AMOUNT:", *callAmount)
		matchesInfo, err := client.Riot.LoL.Match.Get(item)
		if err != nil {
			fmt.Println("ERR", err)
			continue
		}
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
