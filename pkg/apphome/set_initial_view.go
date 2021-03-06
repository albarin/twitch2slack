package apphome

import (
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

func (appHome *AppHome) SetInitialView(event *slackevents.AppHomeOpenedEvent) error {
	slackUserID := event.User

	slackAuth, err := appHome.slackRepo.GetByUserID(slackUserID)
	if err != nil {
		return err
	}

	twitchAuth, err := appHome.twitchRepo.GetBySlackUserID(slackUserID)
	if err != nil {
		return err
	}

	// TODO: consider to verify the token
	follows, err := appHome.twitchAPI.UserFollows(twitchAuth.UserID, twitchAuth.AccessToken)
	if err != nil {
		return err
	}

	var options []*slack.OptionBlockObject
	for _, follow := range follows {
		option := &slack.OptionBlockObject{
			Text: &slack.TextBlockObject{
				Type: slack.PlainTextType,
				Text: follow.ToName,
			},
			Value: follow.ToID,
		}

		options = append(options, option)
	}

	//altOptions := []*slack.OptionBlockObject{
	//	{
	//		Text: &slack.TextBlockObject{
	//			Type:     slack.PlainTextType,
	//			Text:     "asdas",
	//			Emoji:    false,
	//			Verbatim: false,
	//		},
	//		Value:       "asdf",
	//		Description: nil,
	//		URL:         "",
	//	},
	//}

	view := slack.HomeTabViewRequest{
		Type: slack.VTHomeTab,
		Blocks: slack.Blocks{
			BlockSet: []slack.Block{
				slack.ActionBlock{
					Type: slack.MBTAction,
					Elements: &slack.BlockElements{
						ElementSet: []slack.BlockElement{
							slack.SelectBlockElement{
								Type: slack.OptTypeStatic,
								Placeholder: &slack.TextBlockObject{
									Type: slack.PlainTextType,
									Text: "Select a channel",
								},
								Options:  options,
								ActionID: "select-follow",
							},
						},
					},
				},
			},
		},
	}

	api := slack.New(slackAuth.AccessToken)

	// TODO: what about the hash parameter???
	viewResponse, err := api.PublishView(slackUserID, view, "")
	if err != nil {
		return err
	}

	err = viewResponse.Err()
	if err != nil {
		return err
	}

	return nil
}
