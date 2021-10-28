package participants

import (
	//"database/sql"
	//"gitlab.com/idoko/bucketeer/models"
	database "github.com/barqus/fillq_backend/internal/database"
)

type Storage interface {
	getAllParticipants() ([]*Participant, error)
}

type participantStorage struct {
	db *database.Database
}

func MustNewStorage(db *database.Database) Storage {
	return &participantStorage{
		db: db,
	}
}
func (pS participantStorage) getAllParticipants() ([]*Participant, error) {

	rows, err := pS.db.Conn.Query("SELECT * FROM participants ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	allParticipants := make([]*Participant, 0)
	for rows.Next() {
		item := &Participant{}
		err = rows.Scan(
			&item.Id,
			&item.Name,
			&item.Surname,
			&item.BirthDay,
			&item.description,
			&item.Nickname,
			&item.SummonerName,
			&item.TwitchChannel,
		)
		if err != nil {
			return allParticipants, err
		}
		allParticipants = append(allParticipants, item)
	}
	return allParticipants, nil
}

//func (db Database) AddItem(item *models.Item) error {
//	var id int
//	var createdAt string
//	query := `INSERT INTO items (name, description) VALUES ($1, $2) RETURNING id, created_at`
//	err := db.Conn.QueryRow(query, item.Name, item.Description).Scan(&id, &createdAt)
//	if err != nil {
//		return err
//	}
//	item.ID = id
//	item.CreatedAt = createdAt
//	return nil
//}
//func (db Database) GetItemById(itemId int) (models.Item, error) {
//	item := models.Item{}
//	query := `SELECT * FROM items WHERE id = $1;`
//	row := db.Conn.QueryRow(query, itemId)
//	switch err := row.Scan(&item.ID, &item.Name, &item.Description, &item.CreatedAt); err {
//	case sql.ErrNoRows:
//		return item, ErrNoMatch
//	default:
//		return item, err
//	}
//}
//func (db Database) DeleteItem(itemId int) error {
//	query := `DELETE FROM items WHERE id = $1;`
//	_, err := db.Conn.Exec(query, itemId)
//	switch err {
//	case sql.ErrNoRows:
//		return ErrNoMatch
//	default:
//		return err
//	}
//}
//func (db Database) UpdateItem(itemId int, itemData models.Item) (models.Item, error) {
//	item := models.Item{}
//	query := `UPDATE items SET name=$1, description=$2 WHERE id=$3 RETURNING id, name, description, created_at;`
//	err := db.Conn.QueryRow(query, itemData.Name, itemData.Description, itemId).Scan(&item.ID, &item.Name, &item.Description, &item.CreatedAt)
//	if err != nil {
//		if err == sql.ErrNoRows {
//			return item, ErrNoMatch
//		}
//		return item, err
//	}
//	return item, nil
//}
