package questions

import (
	"database/sql"
	"github.com/barqus/fillq_backend/config"
	database "github.com/barqus/fillq_backend/internal/database"
)

type Storage interface {
	addNewQnAToParticipant(qna *QuestionAndAnswer) error
	getAllQnAByParticipantID(participantID int) ([]*QuestionAndAnswer, error)
	deleteQuestionByID(id int) error
}

type qnaStorage struct {
	db *database.Database
}

func MustNewStorage(db *database.Database) Storage {
	return &qnaStorage{
		db: db,
	}
}

func (qnaS qnaStorage) addNewQnAToParticipant(qna *QuestionAndAnswer) error {
	query := `INSERT INTO questions (id, participant_id, question, answer) VALUES ($1, $2, $3, $4)`
	_, err := qnaS.db.Conn.Exec(query, &qna.ParticipantID, &qna.Question, &qna.Answer)
	if err != nil {
		return err
	}
	return nil
}

func (qnaS qnaStorage) getAllQnAByParticipantID(participantID int) ([]*QuestionAndAnswer, error) {
	rows, err := qnaS.db.Conn.Query("SELECT id, participant_id, question, answer FROM questions WHERE participant_id=$1;", participantID)

	if err != nil {
		if err == nil {
			return nil, config.PICKEMS_NOT_FOUND
		}
		return nil, err
	}

	defer rows.Close()
	allParticipantsQNA := make([]*QuestionAndAnswer, 0)
	for rows.Next() {
		item := &QuestionAndAnswer{}
		err = rows.Scan(
			&item.ID,
			&item.ParticipantID,
			&item.Question,
			&item.Answer,
		)
		if err != nil {
			return nil, err
		}
		allParticipantsQNA = append(allParticipantsQNA, item)
	}

	return allParticipantsQNA, nil
}

func (qnaS qnaStorage) deleteQuestionByID(id int) error {
	query := `DELETE FROM questions WHERE id = $1;`
	_, err := qnaS.db.Conn.Exec(query, id)
	switch err {
	case sql.ErrNoRows:
		return sql.ErrNoRows
	default:
		return err
	}
	return nil
}
