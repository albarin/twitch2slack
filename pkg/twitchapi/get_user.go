package twitchapi

import (
	"fmt"

	"github.com/nicklaw5/helix"
)

type UserPayload struct {
	Data []User `json:"data"`
}

type User struct {
	ID string `json:"id"`
}

func (api *API) GetUser(userToken string) (*helix.User, error) {
	client, err := helix.NewClient(&helix.Options{
		ClientID:        api.clientID,
		UserAccessToken: userToken,
	})

	users, err := client.GetUsers(&helix.UsersParams{})
	if err != nil {
		return nil, err
	}

	if users.Error != "" {
		return nil, fmt.Errorf("getUsersFollows error: %s", users.Error)
	}

	return &users.Data.Users[0], nil
}
