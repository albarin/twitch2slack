package main

import (
	"context"
	"database/sql"
	"os"
	"strings"
	"time"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg := generateConfig()

	db, err := openDB(cfg)
	if err != nil {
		log.Error().Err(err).Msg("error connecting to the db")
		os.Exit(1)
	}
	defer db.Close()

	app := application{
		config: cfg,
		logger: zerolog.New(os.Stderr).With().Timestamp().Logger(),
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

func twitchOAuthURL(clientID string, redirectURL string) string {
	url := "https://id.twitch.tv/oauth2/authorize?access_type=offline&client_id={clientID}&redirect_uri={redirectURL}&response_type=code&scope=user:read:subscriptions&state={state}"

	replacer := strings.NewReplacer(
		"{clientID}", clientID,
		"{redirectURL}", redirectURL,
	)

	return replacer.Replace(url)
}
