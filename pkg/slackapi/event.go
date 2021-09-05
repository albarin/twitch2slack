package slackapi

type Payload struct {
	Token    string `json:"token"`
	TeamId   string `json:"team_id"`
	ApiAppId string `json:"api_app_id"`
	Event    struct {
		Type    string `json:"type"`
		UserID  string `json:"user"`
		Channel string `json:"channel"`
		Tab     string `json:"tab"`
		View    View   `json:"view"`
		EventTs string `json:"event_ts"`
	} `json:"event"`
	Type           string `json:"type"`
	EventId        string `json:"event_id"`
	EventTime      int    `json:"event_time"`
	Authorizations []struct {
		EnterpriseId        interface{} `json:"enterprise_id"`
		TeamId              string      `json:"team_id"`
		UserId              string      `json:"user_id"`
		IsBot               bool        `json:"is_bot"`
		IsEnterpriseInstall bool        `json:"is_enterprise_install"`
	} `json:"authorizations"`
	IsExtSharedChannel bool `json:"is_ext_shared_channel"`
}
