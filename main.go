package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	_ "github.com/eleboucher/slack-bot/cog/hello"
	_ "github.com/eleboucher/slack-bot/cog/weather"

	"github.com/eleboucher/slack-bot/slack"
	"github.com/joho/godotenv"
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
