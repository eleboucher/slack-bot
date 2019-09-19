package main

import (
	"github.com/nlopes/slack"
)

var (
	api *slack.Client
	rtm *slack.RTM
)

func sendResponse(message *Response) {
	channel := message.Channel

	if message.ThreadTimestamp != "" {
		message.options = append(message.options, slack.MsgOptionTS(message.Timestamp))
	}
	api.PostMessage(channel, message.options...)
}

func Run(token string) {
	api = slack.New(token)
	rtm = api.NewRTM()

	bot := New(sendResponse)

	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.ConnectedEvent:
			logger.Println("Ready")
		case *slack.MessageEvent:
			if ev.Msg.User != "" {
				bot.receiveMessage(&Request{
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
