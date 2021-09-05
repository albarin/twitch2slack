package main

import (
	"net/http"

	"github.com/slack-go/slack/slackevents"
)

func (app application) handleSlackEvents() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := app.readBody(r)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		// TODO: consider to verify the token
		event, err := slackevents.ParseEvent(body, slackevents.OptionNoVerifyToken())
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		switch event.InnerEvent.Type {
		case slackevents.AppHomeOpened:
			appHomeOpenedEvent := event.InnerEvent.Data.(*slackevents.AppHomeOpenedEvent)

			err := app.processAppHomeOpenedEvent(appHomeOpenedEvent)
			if err != nil {
				app.serverErrorResponse(w, r, err)
				return
			}
		}
	}
}

func (app application) processAppHomeOpenedEvent(event *slackevents.AppHomeOpenedEvent) error {
	if event.View.ID == "" {
		err := app.appHome.SetInitialView(event)
		if err != nil {
			return err
		}
	}

	return nil
}
