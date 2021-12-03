package participants

import (
	//"database/sql"
	//"gitlab.com/idoko/bucketeer/models"
	"database/sql"
	database "github.com/barqus/fillq_backend/internal/database"
)

type Storage interface {
	GetAllParticipants() ([]*Participant, error)
	AddParticipant(participant *Participant) error
	deleteParticipant(praticipantId int) error
}

type participantStorage struct {
	db *database.Database
}

func MustNewStorage(db *database.Database) Storage {
	return &participantStorage{
		db: db,
	}
}
func (pS participantStorage) GetAllParticipants() ([]*Participant, error) {

	rows, err := pS.db.Conn.Query(`
		select
			p.*,
			t.game_name,
			t.is_live,
			t.title,
			t.started_at,
			s.tier,
			s."rank",
			s.points,
			s.wins,
			s.losses
		FROM
			participants p 
		inner JOIN twitch t ON LOWER(t.display_name) = LOWER(p.twitch_channel)
		inner join summoners s ON LOWER(s.summoner_name) = LOWER(p.summoner_name) 
		order by id desc 
		`)
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
			&item.Description,
			&item.Nickname,
			&item.SummonerName,
			&item.TwitchChannel,
			&item.Instagram,
			&item.Twitter,
			&item.Youtube,
			&item.GameName,
			&item.IsLive,
			&item.Title,
			&item.StartedAt,
			&item.Tier,
			&item.Rank,
			&item.Points,
			&item.Wins,
			&item.Losses,
		)
		if err != nil {
			return allParticipants, err
		}
		allParticipants = append(allParticipants, item)
	}
	return allParticipants, nil
}

func (pS participantStorage) AddParticipant(participant *Participant) error {
	query := `INSERT INTO participants (name, surname, birth_day, description, nickname, summoner_name, twitch_channel,instagram,twitter,youtube) VALUES ($1, $2, $3, $4, $5, $6, $7,$8,$9,$10)`
	_, err := pS.db.Conn.Exec(query, participant.Name, participant.Surname, participant.BirthDay,
		participant.Description, participant.Nickname, participant.SummonerName, participant.TwitchChannel, participant.Instagram, participant.Twitter, participant.Youtube)
	if err != nil {
		return err
	}

	return nil
}

func (pS participantStorage) deleteParticipant(participantId int) error {
	query := `DELETE FROM participants WHERE id = $1;`
	_, err := pS.db.Conn.Exec(query, participantId)
	switch err {
	case sql.ErrNoRows:
		return sql.ErrNoRows
	default:
		return err
	}
	return nil
}

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
