package main

import (
	"os"

	"github.com/genesixx/slack-bot/slack"
)

func main() {
	slack.Run(os.Getenv("SLACK_TOKEN"))
}
