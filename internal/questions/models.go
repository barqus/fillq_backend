package questions

type QuestionAndAnswer struct {
	ID       int    `json:"id"`
	ParticipantID   int    `json:"user_id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
