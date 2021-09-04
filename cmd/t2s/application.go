package main

import (
	"fmt"
	"net/http"

	"github.com/albarin/t2s/pkg/slackoauth"
	"github.com/albarin/t2s/pkg/slackrepo"
	"github.com/albarin/t2s/pkg/twitchoauth"
	"github.com/rs/zerolog"
)

type application struct {
	config      config
	logger      zerolog.Logger
	slackOauth  *slackoauth.Service
	twitchOauth *twitchoauth.Service
	slackRepo   *slackrepo.Repo
}

func (app application) serve() error {
	server := &http.Server{
		Addr:    fmt.Sprintf("localhost:%d", app.config.port),
		Handler: app.routes(),
	}

	app.logger.Info().Int("port", app.config.port).Msg("starting server")

	return server.ListenAndServe()
}
