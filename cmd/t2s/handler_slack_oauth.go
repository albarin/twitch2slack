package main

import (
	"net/http"
)

func (app application) handleSlackAuthorization() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		if code == "" {
			app.badRequestResponse(w, MissingCodeParamErr)
			return
		}

		slackAuth, err := app.slackOauth.Authorize(code)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		http.Redirect(w, r, app.twitchOAuthURL(slackAuth.UserID), http.StatusSeeOther)
	}
}
