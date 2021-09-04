package twitchoauth

import (
	"context"

	"github.com/albarin/t2s/pkg/twitchapi"
	"github.com/albarin/t2s/pkg/twitchrepo"
	"golang.org/x/oauth2"
)

type Service struct {
	repo      *twitchrepo.Repo
	oauth     *oauth2.Config
	twitchAPI *twitchapi.API
}

func New(oauthConfig *oauth2.Config, repo *twitchrepo.Repo, twitchAPI *twitchapi.API) *Service {
	return &Service{
		repo:      repo,
		twitchAPI: twitchAPI,
		oauth:     oauthConfig,
	}
}

func (s *Service) Authorize(code string, slackUserID string) (*twitchrepo.Auth, error) {
	token, err := s.oauth.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	auth, err := s.buildAuth(token, slackUserID)
	if err != nil {
		return nil, err
	}

	err = s.repo.Create(auth)
	if err != nil {
		return nil, err
	}

	return auth, nil
}

func (s *Service) buildAuth(token *oauth2.Token, slackUserID string) (*twitchrepo.Auth, error) {
	user, err := s.twitchAPI.GetUser(token.AccessToken)
	if err != nil {
		return nil, err
	}

	return &twitchrepo.Auth{
		UserID:       user.ID,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		SlackUserID:  slackUserID,
	}, nil
}
