package bot

import (
	"bytes"
	"log"

	"github.com/nlopes/slack"
)

var (
	buf           bytes.Buffer
	logger        = log.New(&buf, "logger: ", log.Lshortfile)
	cmdPrefixes   = []string{"!"}
	messageBuffer = 50
)

type Bot struct {
	handler        ResponseHandler
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
		messagesToSend: make(chan Response, messageBuffer),
		done:           make(chan struct{}),
	}

	go b.processMessage()

	return &b
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
