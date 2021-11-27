package leagueoflegends

//fmt.Println(leagueItem.LeaguePoints, leagueItem.SummonerName, leagueItem.Wins, leagueItem.Losses, leagueItem.Rank, leagueItem.Tier)

type SummonerLeague struct {
	PUUID         string
	SummonerName  string
	Tier          *string
	Rank          *string
	Points        *int
	Wins          int
	Losses        int
	ParticipantID int
}
