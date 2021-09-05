package twitchapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
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

func (api *API) UserFollows(userID, token string) ([]Follows, error) {
	url := fmt.Sprintf("%s/users/follows?from_id=%s", api.baseURL, userID)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Client-ID", api.clientID)

	res, err := api.client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var payload FollowsPayload
	err = json.Unmarshal(body, &payload)
	if err != nil {
		return nil, err
	}

	return payload.Data, nil
}
