package subscriptionsrepo

import (
	"database/sql"
)

type Repo struct {
	db *sql.DB
}

func New(db *sql.DB) *Repo {
	return &Repo{db: db}
}

type Subscription struct {
	ID           string `db:"subscription_id"`
	SlackUserID  string `db:"slack_auth_id"`
	TwitchUserID string `db:"twitch_auth_id"`
}

func (r Repo) GetSubscriptionByID(subscriptionID string) ([]Subscription, error) {
	query := `
        SELECT subscription_id, twitch_user_id, slack_user_id 
        FROM subscriptions 
        WHERE subscription_id = $1
	`

	rows, err := r.db.Query(query, subscriptionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	subscriptions := make([]Subscription, 0)
	for rows.Next() {
		var sub Subscription
		err := rows.Scan(&sub.ID, &sub.TwitchUserID, &sub.SlackUserID)
		if err != nil {
			return nil, err
		}

		subscriptions = append(subscriptions, sub)
	}

	return subscriptions, nil
}
