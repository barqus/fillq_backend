package pickems

type PickEm struct {
	Nickname      string `json:"nickname"`
	UserID        int    `json:"user_id"`
	ParticipantID int    `json:"participant_id"`
	Position      int    `json:"position"`

}
