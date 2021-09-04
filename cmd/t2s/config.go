package main

import (
	"flag"
	"os"
)

type config struct {
	env    string
	port   int
	dbDSN  string
	twitch struct {
		clientID     string
		clientSecret string
		redirectURL  string
		oauthURL     string
		eventSecret  string
	}
	slack struct {
		appID        string
		clientID     string
		clientSecret string
		redirectURL  string
	}
}

func generateConfig() config {
	var cfg config

	flag.StringVar(&cfg.env, "env", "dev", "Environment to run the app")
	flag.IntVar(&cfg.port, "port", 3030, "API server port")
	flag.StringVar(&cfg.dbDSN, "dsn", os.Getenv("DATABASE_URL"), "Database dsn")
	flag.Parse()

	cfg.twitch.clientID = os.Getenv("TWITCH_CLIENT_ID")
	cfg.twitch.clientSecret = os.Getenv("TWITCH_CLIENT_SECRET")
	cfg.twitch.redirectURL = os.Getenv("TWITCH_REDIRECT_URL")
	cfg.twitch.eventSecret = os.Getenv("TWITCH_EVENT_SECRET")

	cfg.slack.appID = os.Getenv("SLACK_APP_ID")
	cfg.slack.clientID = os.Getenv("SLACK_CLIENT_ID")
	cfg.slack.clientSecret = os.Getenv("SLACK_CLIENT_SECRET")
	cfg.slack.redirectURL = os.Getenv("SLACK_REDIRECT_URL")

	return cfg
}
