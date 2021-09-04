package slackoauth

import (
	"context"

	"github.com/albarin/t2s/pkg/slackrepo"
	"golang.org/x/oauth2"
)

type Service struct {
	oauth *oauth2.Config
	repo  *slackrepo.Repo
}

func New(oauthConfig *oauth2.Config, repo *slackrepo.Repo) *Service {
	return &Service{
		repo:  repo,
		oauth: oauthConfig,
	}
}

func (s *Service) Authorize(code string) (*slackrepo.Auth, error) {
	token, err := s.oauth.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	auth := buildAuth(token)

	err = s.repo.Create(auth)
	if err != nil {
		return nil, err
	}

	return auth, nil
}

func buildAuth(token *oauth2.Token) *slackrepo.Auth {
	team := token.Extra("team").(map[string]interface{})
	authedUser := token.Extra("authed_user").(map[string]interface{})

	return &slackrepo.Auth{
		UserID:      authedUser["id"].(string),
		TeamID:      team["id"].(string),
		AccessToken: token.AccessToken,
	}
}
