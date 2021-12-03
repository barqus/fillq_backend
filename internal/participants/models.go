package participants

import "time"

type Participant struct {
	Id            int       `json:"participant_id"`
	Name          string    `json:"name"`
	Surname       string    `json:"surname"`
	BirthDay      time.Time `json:"birth_day"`
	Description   string    `json:"description"`
	Nickname      string    `json:"nickname"`
	SummonerName  string    `json:"summoner_name"`
	TwitchChannel string    `json:"twitch_channel"`
	Instagram     string    `json:"instagram"`
	Twitter       string    `json:"twitter"`
	Youtube       string    `json:"youtube"`
	GameName      *string    `json:"game_name"`
	IsLive        *bool      `json:"is_live"`
	Title         *string    `json:"title"`
	StartedAt     *string    `json:"started_at"`
	Tier          *string	`json:"tier"`
	Rank          *string	`json:"rank"`
	Points        *int`json:"points"`
	Wins          *int `json:"wins"`
	Losses        *int `json:"losses"`
}
