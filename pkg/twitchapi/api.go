package twitchapi

import (
	"net/http"
)

type API struct {
	baseURL  string
	clientID string
	client   *http.Client
}

func New(client *http.Client, baseURL, clientID string) *API {
	return &API{
		baseURL:  baseURL,
		clientID: clientID,
		client:   client,
	}
}
