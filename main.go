package main

import (
	"os"

	_ "github.com/genesixx/slack-bot/cog"
	"github.com/genesixx/slack-bot/slack"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	slack.Run(os.Getenv("SLACK_TOKEN"))
}
