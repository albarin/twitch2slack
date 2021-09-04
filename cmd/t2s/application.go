package main

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog"
)

type application struct {
	config config
	logger zerolog.Logger
}

func (app application) serve() error {
	server := &http.Server{
		Addr:    fmt.Sprintf("localhost:%d", app.config.port),
		Handler: app.routes(),
	}

	app.logger.Info().Int("port", app.config.port).Msg("starting server")

	return server.ListenAndServe()
}
