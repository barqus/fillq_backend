package scheduler

import (
	"github.com/KnutZuidema/golio"
	"github.com/barqus/fillq_backend/internal/database"
	"github.com/barqus/fillq_backend/internal/leagueoflegends"
	"github.com/jasonlvhit/gocron"
	"os"
)

func updateDatabase(databaseClient *database.Database, lolClient *golio.Client) {
	leagueoflegends.UpdateSummonersInformation(databaseClient,lolClient)
}

func RunTasks() {
	databaseClient, _ := database.Initialize(os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	lolClient := leagueoflegends.SetupRiotClient()
	gocron.Every(60).Seconds().Do(updateDatabase,databaseClient, lolClient)
	<- gocron.Start()
}

