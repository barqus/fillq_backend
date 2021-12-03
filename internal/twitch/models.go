package twitch

type TwitchOauth struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

type TwitchReturnedObject struct {
	Data []struct {
		BroadcasterLanguage string        `json:"broadcaster_language"`
		BroadcasterLogin    string        `json:"broadcaster_login"`
		DisplayName         string        `json:"display_name"`
		GameId              string        `json:"game_id"`
		GameName            string        `json:"game_name"`
		Id                  string        `json:"id"`
		IsLive              bool          `json:"is_live"`
		TagIds              []interface{} `json:"tag_ids"`
		ThumbnailUrl        string        `json:"thumbnail_url"`
		Title               string        `json:"title"`
		StartedAt           string        `json:"started_at"`
	} `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
}

type TwitchObject struct {
	TwitchID      string `json:"twitch_id"`
	ParticipantID int    `json:"participant_id"`
	DisplayName   string `json:"display_name"`
	GameName      string `json:"game_name"`
	IsLive        bool   `json:"is_live"`
	Title         string `json:"title"`
	StartedAt     string `json:"started_at"`
}
