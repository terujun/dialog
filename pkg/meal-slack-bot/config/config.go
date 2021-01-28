package config

import (
	"github.com/terujun/dialog/pkg/meal-slack-bot/slack"
)

type (
	Config struct {
		Slack slack.Config `json:slack`
	}
)
