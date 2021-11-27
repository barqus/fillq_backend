package pickems

import (
	"database/sql"
	"github.com/barqus/fillq_backend/config"
	"github.com/barqus/fillq_backend/internal/database"
)

type Storage interface {
	getUsersPickems(id int) ([]*PickEm, error)
	createUsersPickems(usersPickems []*PickEm) error
	deleteUsersAllPickems(userID int) error
}

type pickemStorage struct {
	db *database.Database
}

func MustNewStorage(db *database.Database) Storage {
	return &pickemStorage{
		db: db,
	}
}

func (p pickemStorage) getUsersPickems(id int) ([]*PickEm, error) {
	rows, err := p.db.Conn.Query("SELECT user_id, participant_id, p.nickname, position FROM pickems JOIN participants p on p.id = pickems.participant_id WHERE user_id=$1;", id)

	if err != nil {
		if err == nil {
			return nil, config.PICKEMS_NOT_FOUND
		}
		return nil, err
	}

	defer rows.Close()
	allUsersPickems := make([]*PickEm, 0)
	for rows.Next() {
		item := &PickEm{}
		err = rows.Scan(
			&item.UserID,
			&item.ParticipantID,
			&item.Nickname,
			&item.Position,
		)
		if err != nil {
			return nil, err
		}
		allUsersPickems = append(allUsersPickems, item)
	}

	return allUsersPickems, nil
}

func (p pickemStorage) createUsersPickems(usersPickems []*PickEm) error {
	for _, item := range usersPickems {
		query := `INSERT INTO pickems (user_id,participant_id,position) VALUES ($1, $2, $3)`
		_, err := p.db.Conn.Exec(query, item.UserID, item.ParticipantID, item.Position)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p pickemStorage) deleteUsersAllPickems(userID int) error {
	query := `DELETE FROM pickems WHERE user_id = $1;`
	_, err := p.db.Conn.Exec(query, userID)
	switch err {
	case sql.ErrNoRows:
		return sql.ErrNoRows
	default:
		return err
	}
	return nil
}