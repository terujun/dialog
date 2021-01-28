package config

type (
	Config struct {
		Slack slack.Config `json:slack`
	}
)
