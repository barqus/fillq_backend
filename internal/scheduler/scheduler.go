package scheduler

import (
	"github.com/KnutZuidema/golio"
	"github.com/barqus/fillq_backend/internal/common_http"
	"github.com/barqus/fillq_backend/internal/database"
	"github.com/barqus/fillq_backend/internal/leagueoflegends"
	"github.com/barqus/fillq_backend/internal/twitch"
	"net/http"

	//"github.com/barqus/fillq_backend/internal/twitch"
	"github.com/jasonlvhit/gocron"
	"os"
)

func updateDatabaseLeague(databaseClient *database.Database, lolClient *golio.Client) {
	leagueoflegends.UpdateSummonersInformation(databaseClient,lolClient)
}

func updateDatabaseTwitch(databaseClient *database.Database, lolClient *golio.Client) {
	httpClient := common_http.NewClient(http.DefaultClient)
	twitch.UpdateUserTwitch(httpClient, databaseClient)
}


func RunTasks() {
	databaseClient, _ := database.Initialize(os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	lolClient := leagueoflegends.SetupRiotClient()
	gocron.Every(60).Seconds().Do(updateDatabaseLeague,databaseClient, lolClient)
	gocron.Every(10).Seconds().Do(updateDatabaseTwitch,databaseClient, lolClient)
	<- gocron.Start()
}

