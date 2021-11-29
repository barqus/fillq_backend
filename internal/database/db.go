package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	HOST = "fillq-db.capzgilbiuua.eu-central-1.rds.amazonaws.com"
	PORT = 5432
)

// ErrNoMatch is returned when we request a row that doesn't exist
var ErrNoMatch = fmt.Errorf("no matching record")

type Database struct {
	Conn *sql.DB
}

func Initialize(username, password, database string) (*Database, error) {
	// TODO: CHECK WHAT VARIABLES ARE RUN WITH MIGRATION
	// export POSTGRESQL_URL=""
	// migrate -database ${POSTGRESQL_URL} -path ./migrations upv
	db := Database{}
	//dsn := fmt.Sprintf( "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	//	HOST, PORT, username, password, database)

	dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		username,
		password,
		HOST,
		PORT,
		"postgres")

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return &db, err
	}
	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return &db, err
	}

	return &db, nil
}
