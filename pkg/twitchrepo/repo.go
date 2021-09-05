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

func (r Repo) GetBySlackUserID(slackUserID string) (*Auth, error) {
	var auth Auth

	query := `
		SELECT ta.user_id, ta.access_token, ta.refresh_token, ta.slack_user_id
		FROM twitch_auths ta
		JOIN slack_auths sa on sa.user_id = ta.slack_user_id
		WHERE sa.user_id = $1
	`

	args := []interface{}{&auth.UserID, &auth.AccessToken, &auth.RefreshToken, &auth.SlackUserID}
	err := r.db.QueryRow(query, slackUserID).Scan(args...)
	if err != nil {
		return nil, err
	}

	return &auth, nil
}
