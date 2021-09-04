package slackapi

import (
	"io"
	"io/ioutil"
	"net/http"
)

type API struct {
	client  *http.Client
	baseURL string
}

func New(client *http.Client, baseURL string) *API {
	return &API{
		client:  client,
		baseURL: baseURL,
	}
}

func (api *API) makeRequest(token, method, path string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, api.baseURL+path, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+token)

	res, err := api.client.Do(req)
	if err != nil {
		return nil, err
	}

	responseView, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return responseView, nil
}
