package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/albarin/t2s/pkg/notifications"
	"github.com/albarin/t2s/pkg/slackoauth"
	"github.com/albarin/t2s/pkg/slackrepo"
	subscriptionsrepo "github.com/albarin/t2s/pkg/subscriptionrepo"
	"github.com/albarin/t2s/pkg/twitchoauth"
	"github.com/rs/zerolog"
)

type application struct {
	config        config
	logger        zerolog.Logger
	slackOauth    *slackoauth.Service
	twitchOauth   *twitchoauth.Service
	slackRepo     *slackrepo.Repo
	subsRepo      *subscriptionsrepo.Repo
	notifications *notifications.Notifications
}

func (app application) serve() error {
	server := &http.Server{
		Addr:    fmt.Sprintf("localhost:%d", app.config.port),
		Handler: app.routes(),
	}

	app.logger.Info().Int("port", app.config.port).Msg("starting server")

	return server.ListenAndServe()
}

func (app application) readBody(r *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (app *application) parseJSON(data []byte, dst interface{}) error {
	dec := json.NewDecoder(bytes.NewReader(data))
	//dec.DisallowUnknownFields()

	err := dec.Decode(dst)
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only contain a single JSON value")
	}

	return nil
}

func (app application) isProduction() bool {
	return app.config.env == "production"
}
