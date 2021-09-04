package slackapi

import (
	"bytes"
	"fmt"
	"net/http"
)

func (api *API) PostMessage(token string, channel string, blocks []byte) ([]byte, error) {
	body := []byte(fmt.Sprintf(`{"channel":%q, "blocks":%s}`, channel, string(blocks)))

	res, err := api.makeRequest(token, http.MethodPost, "/chat.postMessage", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	return res, nil
}
