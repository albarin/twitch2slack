package main

import (
	"net/http"
)

func (app application) handleTwitchAuthorization() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		if code == "" {
			app.badRequestResponse(w, MissingCodeParamErr)
			return
		}

		slackUserID := r.URL.Query().Get("state")
		if slackUserID == "" {
			app.badRequestResponse(w, MissingStateParamErr)
			return
		}

		_, err := app.twitchOauth.Authorize(code, slackUserID)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		slackAuth, err := app.slackRepo.GetByUserID(slackUserID)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		http.Redirect(w, r, app.slackHomeTabURL(slackAuth.TeamID), http.StatusSeeOther)
	}
}
