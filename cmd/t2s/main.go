package main

import (
	"context"
	"database/sql"
	"os"
	"strings"
	"time"

	"github.com/albarin/t2s/pkg/slackoauth"
	"github.com/albarin/t2s/pkg/slackrepo"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
)

func main() {
	cfg := generateConfig()

	db, err := openDB(cfg)
	if err != nil {
		log.Error().Err(err).Msg("error connecting to the db")
		os.Exit(1)
	}
	defer db.Close()

	slackRepo := slackrepo.New(db)
	slackOauthConfig := &oauth2.Config{
		ClientID:     cfg.slack.clientID,
		ClientSecret: cfg.slack.clientSecret,
		Scopes:       []string{"chat:write"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://slack.com/oauth/v2/authorize",
			TokenURL: "https://slack.com/api/oauth.v2.access",
		},
		RedirectURL: cfg.slack.redirectURL,
	}

	app := application{
		config: cfg,
		logger: zerolog.New(os.Stderr).With().Timestamp().Logger(),

		slackOauth: slackoauth.New(slackOauthConfig, slackRepo),
	}

	err = app.serve()
	if err != nil {
		log.Error().Err(err).Int("port", cfg.port).Msg("error starting the server")
		os.Exit(1)
	}
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.dbDSN)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (app application) twitchOAuthURL(slackUserID string) string {
	url := "https://id.twitch.tv/oauth2/authorize?access_type=offline&client_id={clientID}&redirect_uri={redirectURL}&response_type=code&scope=user:read:subscriptions&state={state}"

	replacer := strings.NewReplacer(
		"{clientID}", app.config.twitch.clientID,
		"{redirectURL}", app.config.twitch.redirectURL,
		"{state}", slackUserID,
	)

	return replacer.Replace(url)
}
