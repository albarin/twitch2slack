package twitchrepo

import (
	"database/sql"
)

type Auth struct {
	UserID       string `db:"user_id"`
	AccessToken  string `db:"access_token"`
	RefreshToken string `db:"refresh_token"`
	SlackUserID  string `db:"slack_user_id"`
}

type Repo struct {
	db *sql.DB
}

func New(db *sql.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r Repo) Create(auth *Auth) error {
	query := `
        INSERT INTO twitch_auths (user_id, access_token, refresh_token, slack_user_id) 
        VALUES ($1, $2, $3, $4)
	`

	args := []interface{}{auth.UserID, auth.AccessToken, auth.RefreshToken, auth.SlackUserID}
	_, err := r.db.Exec(query, args...)

	return err
}
