package leagueoflegends

import (
	"database/sql"
	"github.com/barqus/fillq_backend/config"
	"github.com/barqus/fillq_backend/internal/database"
)

type Storage interface {
	AddSummonerLeague(summoner SummonerLeague) error
	UpdateSummonerLeagueByID(summoner SummonerLeague, summonerID string) error
	summonerAlreadyExists(summonerID string) (*bool, error)
}

type leagueoflegendsStorage struct {
	db *database.Database
}

func MustNewStorage(db *database.Database) Storage {
	return &leagueoflegendsStorage{
		db: db,
	}
}

// TODO: KAS BUNA JEI SUZAIDZIA ZAIDIMA?
func (lol leagueoflegendsStorage) AddSummonerLeague(summoner SummonerLeague) error {
	query := `INSERT INTO summoners (puuid, summoner_name, tier, rank, points, wins, losses, participant_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := lol.db.Conn.Exec(query, summoner.PUUID, summoner.SummonerName, &summoner.Tier, &summoner.Rank, &summoner.Points, summoner.Wins, summoner.Losses, summoner.ParticipantID)
	if err != nil {
		return err
	}

	return nil
}

func (lol leagueoflegendsStorage) summonerAlreadyExists(PUUID string) (*bool, error) {
	//select exists(select 1 from summoners where puuid='vIKJHI4OihKYP2LVylIORWtwtbeCh0k8WYGfbKzt2cUkadzGiES5TpQwKQDYduY-kxwQ_XVTwWcstQ')
	query := `SELECT EXISTS(SELECT 1 FROM summoners WHERE puuid=$1)`

	row := lol.db.Conn.QueryRow(query, PUUID)
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

// TODO: KAS BUNA JEI NULL WINS
func (lol leagueoflegendsStorage) UpdateSummonerLeagueByID(summoner SummonerLeague, summonerID string) error {
	query := `UPDATE summoners SET tier=$1, rank=$2, points=$3, wins=$4, losses=$5 WHERE puuid=$6`

	_, err := lol.db.Conn.Exec(query, summoner.Tier, summoner.Rank, summoner.Points, summoner.Wins, summoner.Losses, summonerID)
	if err != nil {
		return err
	}

	return nil
}
