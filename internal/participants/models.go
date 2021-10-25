package participants

type Participant struct {
	Id int 	`json: "id"`
	Name string `json: "name"`
	Surname string `json: "surname"`
	SummonerName string `json: "summoner_name"`
	StreamName string `json: "stream_name"`
}