package twitch

import (
	"database/sql"
	"fmt"
	"github.com/barqus/fillq_backend/config"
	"github.com/barqus/fillq_backend/internal/database"
)

type Storage interface {
	AddTwitchAccount(twitch TwitchObject) error
	UpdateTwitchAccountByID(twitch TwitchObject, twitchID string) error
	twitchAccountAlreadyExists(twitchID string) (*bool, error)
}


type twitchStorage struct {
	db *database.Database
}

func MustNewStorage(db *database.Database) Storage {
	return &twitchStorage{
		db: db,
	}
}

func (ts twitchStorage) AddTwitchAccount(twitch TwitchObject) error {
	fmt.Println(twitch)
	query := `INSERT INTO twitch (twitch_id, display_name, game_name, is_live, title, started_at, participant_id) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := ts.db.Conn.Exec(query, twitch.TwitchID, twitch.DisplayName, twitch.GameName, twitch.IsLive, twitch.Title, twitch.StartedAt, twitch.ParticipantID)
	if err != nil {
		return err
	}

	return nil
}

func (ts twitchStorage) UpdateTwitchAccountByID(twitch TwitchObject, twitchID string) error {
	query := `UPDATE twitch SET is_live=$1, title=$2, started_at=$3, game_name=$4 WHERE twitch_id=$5`

	_, err := ts.db.Conn.Exec(query, twitch.IsLive, twitch.Title, twitch.StartedAt, twitch.GameName, twitchID)
	if err != nil {
		return err
	}

	return nil
}

func (ts twitchStorage) twitchAccountAlreadyExists(twitchID string) (*bool, error) {
	//select exists(select 1 from summoners where puuid='vIKJHI4OihKYP2LVylIORWtwtbeCh0k8WYGfbKzt2cUkadzGiES5TpQwKQDYduY-kxwQ_XVTwWcstQ')
	query := `SELECT EXISTS(SELECT 1 FROM twitch WHERE twitch_id=$1)`

	row := ts.db.Conn.QueryRow(query, twitchID)
	exists := false
	switch err := row.Scan(&exists); err {
	case sql.ErrNoRows:
		return nil, config.USER_NOT_FOUND
	case nil:
		return &exists, nil
	default:
		return nil, err
	}
}

