package main

import (
	"os"

	_ "github.com/genesixx/slack-bot/cog"
	"github.com/genesixx/slack-bot/slack"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
}

func main() {
	slack.Run(os.Getenv("SLACK_TOKEN"))
}
