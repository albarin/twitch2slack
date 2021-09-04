package twitchapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type UserPayload struct {
	Data []User `json:"data"`
}

type User struct {
	ID string `json:"id"`
}

func (api *API) GetUser(userToken string) (*User, error) {
	url := fmt.Sprintf("%s/users", api.baseURL)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+userToken)
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

	var payload UserPayload
	err = json.Unmarshal(body, &payload)
	if err != nil {
		return nil, err
	}

	return &payload.Data[0], nil
}
