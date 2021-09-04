package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/nicklaw5/helix"
)

type eventSubPayload struct {
	Subscription helix.EventSubSubscription `json:"subscription"`
	Challenge    string                     `json:"challenge"`
	Event        json.RawMessage            `json:"event"`
}

func (app application) handleTwitchEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := app.readBody(r)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		if app.isProduction() && !helix.VerifyEventSubNotification(app.config.twitch.eventSecret, r.Header, string(body)) {
			app.serverErrorResponse(w, r, errors.New("no valid signature on subscription"))
			return
		}

		var event eventSubPayload
		err = app.parseJSON(body, &event)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		if event.Challenge != "" {
			w.Write([]byte(event.Challenge))
			return
		}

		switch event.Subscription.Type {
		case helix.EventSubTypeStreamOnline:
			err := app.processStreamOnline(event)
			if err != nil {
				app.serverErrorResponse(w, r, err)
				return
			}
		case helix.EventSubTypeStreamOffline:
			err := app.processStreamOffline(event)
			if err != nil {
				app.serverErrorResponse(w, r, err)
				return
			}
		default:
			app.logger.Info().
				Str("type", event.Subscription.Type).
				Msg("unknown event received")
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (app application) processStreamOnline(event eventSubPayload) error {
	var streamOnlineEvent helix.EventSubStreamOnlineEvent
	err := app.parseJSON(event.Event, &streamOnlineEvent)
	if err != nil {
		return err
	}

	subs, err := app.subsRepo.GetSubscriptionByID(event.Subscription.ID)
	if err != nil {
		return err
	}

	for _, sub := range subs {
		err := app.notifications.SendStreamOnlineNotification(sub.SlackUserID, streamOnlineEvent.BroadcasterUserName, streamOnlineEvent.BroadcasterUserLogin)
		if err != nil {
			// TODO: log something about the error
			spew.Dump(err)
			continue
		}
	}

	return nil
}

func (app application) processStreamOffline(event eventSubPayload) error {
	var streamOfflineEvent helix.EventSubStreamOfflineEvent
	err := app.parseJSON(event.Event, &streamOfflineEvent)
	if err != nil {
		return err
	}

	spew.Dump("stream.offline", streamOfflineEvent)

	return nil
}
