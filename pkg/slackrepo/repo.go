package slackrepo

import "database/sql"

type Auth struct {
	UserID      string `db:"user_id"`
	TeamID      string `db:"team_id"`
	ViewID      string `db:"view_id"`
	AccessToken string `db:"access_token"`
}

type Repo struct {
	db *sql.DB
}

func New(db *sql.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) Create(auth *Auth) error {
	query := `
        INSERT INTO slack_auths (user_id, team_id, view_id, access_token) 
        VALUES ($1, $2, $3, $4)
	`

	args := []interface{}{auth.UserID, auth.TeamID, auth.ViewID, auth.AccessToken}
	_, err := r.db.Exec(query, args...)

	return err
}

func (r *Repo) GetByUserID(userID string) (*Auth, error) {
	var auth Auth

	query := `
		SELECT user_id, team_id, view_id, access_token
		FROM slack_auths
		WHERE user_id = $1
	`

	args := []interface{}{&auth.UserID, &auth.TeamID, &auth.ViewID, &auth.AccessToken}
	err := r.db.QueryRow(query, userID).Scan(args...)
	if err != nil {
		return nil, err
	}

	return &auth, nil
}
