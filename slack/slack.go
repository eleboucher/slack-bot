package slack

import (
	log "github.com/sirupsen/logrus"

	"github.com/eleboucher/slack-bot/bot"
	"github.com/slack-go/slack"
)

var (
	API *slack.Client
	rtm *slack.RTM
)

func sendResponse(message *bot.Response) {
	channel := message.Channel
	message.Options = append(message.Options, slack.MsgOptionText(message.Message, false))
	if message.ThreadTimestamp != "" {
		message.Options = append(message.Options, slack.MsgOptionTS(message.Timestamp))
	}
	_, _, err := API.PostMessage(channel, message.Options...)
	if err != nil {
		log.Println(err)
	}
}

func Run(token string) {
	API = slack.New(token)
	rtm = API.NewRTM()
	b := bot.New(sendResponse)

	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.ConnectedEvent:
			log.Info("Ready")
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
			log.Error(ev.Error())
		case *slack.InvalidAuthEvent:
			log.Error("Invalid credentials")
			return
		}
	}
}
