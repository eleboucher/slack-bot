package hello

import "github.com/eleboucher/slack-bot/bot"

func hello(cmd *bot.CMD) (*bot.Response, error) {
	return &bot.Response{
		Message:         "Hello <@" + cmd.User + ">!",
		Channel:         cmd.Channel,
		Timestamp:       cmd.Timestamp,
		ThreadTimestamp: cmd.ThreadTimestamp,
	}, nil
}

func init() {
	bot.RegisterCommand("hello", "just say hello", hello)
}
