package users

import (
	//"database/sql"
	//"gitlab.com/idoko/bucketeer/models"
	"database/sql"
	"github.com/barqus/fillq_backend/config"
	database "github.com/barqus/fillq_backend/internal/database"
)

type Storage interface {
	addNewUser(user *User) error
	getUserByID(id string) (*User, error)
	getUserByCode(code string) (*User, error)
}

type userStorage struct {
	db *database.Database
}

func MustNewStorage(db *database.Database) Storage {
	return &userStorage{
		db: db,
	}
}

func (u userStorage) getUserByID(id string) (*User, error) {
	user := &User{}
	query := `SELECT * FROM users WHERE id = $1;`

	row := u.db.Conn.QueryRow(query,id)
	switch err := row.Scan(&user.ID, &user.DisplayName, &user.ProfileImageURL, &user.Email, &user.TwitchCode, &user.Role, &user.AccessToken, &user.RefreshToken,&user.JWTToken); err {
		case sql.ErrNoRows:
			return nil, config.USER_NOT_FOUND
		case nil:
			return user, nil
		default:
			return nil, err
	}
}

func (u userStorage) getUserByCode(code string) (*User, error) {
	user := &User{}
	query := `SELECT * FROM users WHERE twitch_code = $1;`

	row := u.db.Conn.QueryRow(query,code)
	switch err := row.Scan(&user.ID, &user.DisplayName, &user.ProfileImageURL, &user.Email, &user.TwitchCode, &user.Role, &user.AccessToken, &user.RefreshToken,&user.JWTToken); err {
	case sql.ErrNoRows:
		return nil, config.USER_NOT_FOUND
	case nil:
		return user, nil
	default:
		return nil, err
	}
}

func (u userStorage) addNewUser(user *User) error {
	query := `INSERT INTO users (id, display_name, profile_image_url, email, twitch_code, role, access_token, refresh_token, jwt_token) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := u.db.Conn.Exec(query, &user.ID, &user.DisplayName, &user.ProfileImageURL, &user.Email, &user.TwitchCode, &user.Role, &user.AccessToken, &user.RefreshToken,&user.JWTToken)
	if err != nil {
		return err
	}
	return nil
}
