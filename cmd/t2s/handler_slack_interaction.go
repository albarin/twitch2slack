package main

import (
	"encoding/json"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/slack-go/slack"
)

func (app application) handleSlackInteractions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: consider to verify the secret
		payload := r.FormValue("payload")

		var interaction slack.InteractionCallback
		err := json.Unmarshal([]byte(payload), &interaction)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		//spew.Dump("SLACK INTERACTION", interaction)

		if interaction.Type == slack.InteractionTypeBlockActions {
			action := interaction.ActionCallback.BlockActions[0].ActionID

			if action == "select-follow" {
				spew.Dump(action)
			}
		}
	}
}
