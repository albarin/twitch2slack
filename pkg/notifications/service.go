package notifications

import (
	"fmt"

	"github.com/albarin/t2s/pkg/slackapi"
	"github.com/albarin/t2s/pkg/slackrepo"
)

type Notifications struct {
	slackAPI  *slackapi.API
	slackRepo *slackrepo.Repo
}

func New(slackAPI *slackapi.API, slackRepo *slackrepo.Repo) *Notifications {
	return &Notifications{
		slackAPI:  slackAPI,
		slackRepo: slackRepo,
	}
}

func (n *Notifications) SendStreamOnlineNotification(slackUserID, broadcasterUserName, broadcasterUserLogin string) error {
	blocks := []byte(fmt.Sprintf(`[
		{
			"type": "section",
			"text": {
				"type": "mrkdwn",
				"text": "*%s* has started a stream!! <https://twitch.tv/%s|Go check it out>!"
			}
		}
	]`, broadcasterUserName, broadcasterUserLogin))

	slackAuth, err := n.slackRepo.GetByUserID(slackUserID)
	if err != nil {
		return err
	}

	_, err = n.slackAPI.PostMessage(slackAuth.AccessToken, slackUserID, blocks)
	if err != nil {
		return err
	}

	// TODO: handle post message response

	return nil
}
