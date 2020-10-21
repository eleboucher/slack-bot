package bot

import (
	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
)

var (
	cmdPrefixes   = []string{"!"}
	messageBuffer = 50
)

type Bot struct {
	handler        ResponseHandler
	cron           *cron.Cron
	messagesToSend chan Response
	done           chan struct{}
}

type Request struct {
	Message         string
	Channel         string
	User            string
	Timestamp       string
	ThreadTimestamp string
}

type Response struct {
	Message         string
	Channel         string
	User            string
	Timestamp       string
	ThreadTimestamp string
	Options         []slack.MsgOption
}

type ResponseHandler func(message *Response)

func New(handler ResponseHandler) *Bot {
	b := Bot{
		handler:        handler,
		cron:           cron.New(),
		messagesToSend: make(chan Response, messageBuffer),
		done:           make(chan struct{}),
	}

	go b.processMessage()

	b.processPeriodic()
	return &b
}

func (b *Bot) processPeriodic() {
	for cmd, cog := range periodicCommand {
		b.cron.AddFunc(cog.CronSetting, func() {
			resp, err := cog.Function()
			log.Infof("Sending %#v command with resp: %#v", cmd, resp)

			if err != nil {
				log.Errorf("Command %s error: %s\n", cmd, err)
			}

			b.sendResponse(resp)
		})
	}
	if len(b.cron.Entries()) > 0 {
		log.Info("Starting cron")
		b.cron.Start()
	}
}

func (b *Bot) sendResponse(resp *Response) {
	b.handler(resp)
}

func (b *Bot) ReceiveMessage(req *Request) {
	cmd := Parse(req)

	if cmd == nil {
		return
	}

	b.handleCMD(cmd)
}

func (b *Bot) processMessage() {
	for {
		select {
		case msg := <-b.messagesToSend:
			b.sendResponse(&msg)
		case <-b.done:
			return
		}
	}
}

func (b *Bot) close() {
	close(b.done)
}
