/*
This packege describes the format for posting message.
*/
package slack

//Appstart is struct of first message from slack
type Appstart struct {
	Type     string `json:"type"`
	Token    string `json:"token"`
	ActionTs string `json:"action_ts"`
	Team     struct {
		ID     string `json:"id"`
		Domain string `json:"domain"`
	} `json:"team"`
	User struct {
		ID       string `json:"id"`
		Username string `json:"username"`
		TeamID   string `json:"team_id"`
	} `json:"user"`
	IsEnterpriseInstall bool        `json:"is_enterprise_install"`
	Enterprise          interface{} `json:"enterprise"`
	CallbackID          string      `json:"callback_id"`
	TriggerID           string      `json:"trigger_id"`
}

// SlackRepository controls posts to Slack.
type SlackRepository struct {
	Token        string
	ViewsDirPath string
}

type Config struct {
	Token string `json:"token"`
}
