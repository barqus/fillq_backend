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
}
