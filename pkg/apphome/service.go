package apphome

import (
	"github.com/albarin/t2s/pkg/slackapi"
	"github.com/albarin/t2s/pkg/slackrepo"
	"github.com/albarin/t2s/pkg/twitchapi"
	"github.com/albarin/t2s/pkg/twitchrepo"
)

type AppHome struct {
	slackAPI   *slackapi.API
	twitchAPI  *twitchapi.API
	slackRepo  *slackrepo.Repo
	twitchRepo *twitchrepo.Repo
}

func New(slackAPI *slackapi.API, twitchAPI *twitchapi.API, slackRepo *slackrepo.Repo, twitchRepo *twitchrepo.Repo) *AppHome {
	return &AppHome{
		slackAPI:   slackAPI,
		twitchAPI:  twitchAPI,
		slackRepo:  slackRepo,
		twitchRepo: twitchRepo,
	}
}
