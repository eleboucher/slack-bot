package slack

import (
	"bytes"
	"log"

	"github.com/genesixx/slack-bot/bot"
	"github.com/nlopes/slack"
)

var (
	api    *slack.Client
	rtm    *slack.RTM
	buf    bytes.Buffer
	logger = log.New(&buf, "logger: ", log.Lshortfile)
)

func sendResponse(message *bot.Response) {
	channel := message.Channel

	if message.ThreadTimestamp != "" {
		message.Options = append(message.Options, slack.MsgOptionTS(message.Timestamp))
	}
	api.PostMessage(channel, message.Options...)
}

func Run(token string) {
	api = slack.New(token)
	rtm = api.NewRTM()

	b := bot.New(sendResponse)

	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.ConnectedEvent:
			logger.Println("Ready")
		case *slack.MessageEvent:
			if ev.Msg.User != "" {
				b.ReceiveMessage(&bot.Request{
					Message:         ev.Msg.Text,
					Channel:         ev.Msg.Channel,
					User:            ev.Msg.User,
					Timestamp:       ev.Msg.Timestamp,
					ThreadTimestamp: ev.Msg.ThreadTimestamp,
				})
			}
		case *slack.RTMError:
			logger.Fatal(ev.Error())
		case *slack.InvalidAuthEvent:
			logger.Fatal("Invalid credentials")
			return
		}
	}
}
