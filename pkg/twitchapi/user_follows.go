package twitchapi

import (
	"fmt"
	"time"

	"github.com/nicklaw5/helix"
)

type FollowsPayload struct {
	Data []Follows `json:"data"`
}

type Follows struct {
	FromId     string    `json:"from_id"`
	FromLogin  string    `json:"from_login"`
	FromName   string    `json:"from_name"`
	ToID       string    `json:"to_id"`
	ToLogin    string    `json:"to_login"`
	ToName     string    `json:"to_name"`
	FollowedAt time.Time `json:"followed_at"`
}

func (api *API) UserFollows(userID, userToken string) ([]helix.UserFollow, error) {
	client, err := helix.NewClient(&helix.Options{
		ClientID:        api.clientID,
		UserAccessToken: userToken,
	})

	userFollows, err := client.GetUsersFollows(&helix.UsersFollowsParams{FromID: userID})
	if err != nil {
		return nil, err
	}

	if userFollows.Error != "" {
		return nil, fmt.Errorf("getUsersFollows error: %s", userFollows.Error)
	}

	// TODO: consider pagination
	return userFollows.Data.Follows, nil
}
