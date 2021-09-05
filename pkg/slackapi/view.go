package slackapi

import (
	"bytes"
	"net/http"
)

type Block struct {
	Type    string `json:"type"`
	BlockId string `json:"block_id"`
	Text    struct {
		Type     string `json:"type"`
		Text     string `json:"text"`
		Verbatim bool   `json:"verbatim"`
	} `json:"text"`
	Accessory struct {
		Type     string `json:"type"`
		ActionId string `json:"action_id"`
		Text     struct {
			Type  string `json:"type"`
			Text  string `json:"text"`
			Emoji bool   `json:"emoji"`
		} `json:"text"`
		Value string `json:"value"`
		Url   string `json:"url"`
	} `json:"accessory"`
	Elements []Element `json:"elements"`
}

type Element struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type PlainTextBlock struct {
	Type     string    `json:"type"`
	Elements []Element `json:"elements"`
}

type View struct {
	ID              string  `json:"id"`
	TeamId          string  `json:"team_id"`
	Type            string  `json:"type"`
	Blocks          []Block `json:"blocks"`
	PrivateMetadata string  `json:"private_metadata"`
	CallbackId      string  `json:"callback_id"`
	State           struct {
		Values struct {
		} `json:"values"`
	} `json:"state"`
	Hash  string `json:"hash"`
	Title struct {
		Type  string `json:"type"`
		Text  string `json:"text"`
		Emoji bool   `json:"emoji"`
	} `json:"title"`
	ClearOnClose       bool        `json:"clear_on_close"`
	NotifyOnClose      bool        `json:"notify_on_close"`
	Close              interface{} `json:"close"`
	Submit             interface{} `json:"submit"`
	PreviousViewId     interface{} `json:"previous_view_id"`
	RootViewId         string      `json:"root_view_id"`
	AppId              string      `json:"app_id"`
	ExternalId         string      `json:"external_id"`
	AppInstalledTeamId string      `json:"app_installed_team_id"`
	BotId              string      `json:"bot_id"`
}

func (view *View) IsEmpty() bool {
	return view.ID == ""
}

type ResponseView struct {
	OK   bool `json:"ok,omitempty"`
	View View `json:"view"`
}

func (api *API) ViewPublish(view []byte, token string) ([]byte, error) {
	return api.makeRequest(token, http.MethodPost, "/views.publish", bytes.NewBuffer(view))
}

func (api *API) ViewUpdate(view []byte, token string) ([]byte, error) {
	return api.makeRequest(token, http.MethodPost, "/views.update", bytes.NewBuffer(view))
}
