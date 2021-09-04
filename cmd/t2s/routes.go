package main

import (
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var (
	MissingCodeParamErr  = errors.New("missing required code query parameter")
	MissingStateParamErr = errors.New("missing required state query parameter")
)

func (app application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/oauth/slack", app.handleSlackAuthorization())
	router.HandlerFunc(http.MethodGet, "/oauth/twitch", app.handleTwitchAuthorization())

	router.HandlerFunc(http.MethodPost, "/twitch/events", app.handleTwitchEvent())

	return router
}
