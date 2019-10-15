package bot

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

type CMD struct {
	Command         string
	Option          string
	Channel         string
	User            string
	Timestamp       string
	ThreadTimestamp string
}

type cmdFunc func(cmd *CMD) (*Response, error)

type Cog struct {
	cmd      string
	helper   string
	function cmdFunc
}

type PeriodicCog struct {
	Function    func() (*Response, error)
	CronSetting string
}

var (
	commands        = make(map[string]*Cog)
	periodicCommand = make(map[string]*PeriodicCog)
)

//RegisterCommand Register command
func RegisterCommand(cmd, helper string, function cmdFunc) {
	log.Infof("Adding Command %s\n", cmd)
	commands[cmd] = &Cog{
		cmd:      cmd,
		helper:   helper,
		function: function,
	}
}

func RegisterPeriodicCommand(cmd string, config *PeriodicCog) {
	log.Infof("Adding Periodic Command %s\n", cmd)

	periodicCommand[cmd] = config
}

func (b *Bot) handleCMD(cmd *CMD) {
	log.Infof("received new Command %s\n", cmd.Command)
	if cmd.Command == "help" {
		b.sendHelper(cmd)
		return
	}

	c := commands[cmd.Command]

	if c == nil {
		log.Errorf("Command %s not found\n", cmd.Command)
	}

	resp, err := c.function(cmd)

	if err != nil {
		log.Errorf("Command %s error: %s\n", cmd.Command, err)
	}

	b.sendResponse(resp)
}

func (b *Bot) sendHelper(cmd *CMD) {
	helper := "```\n"
	helper += `
Bot Usage

Command:
	`

	for _, cog := range commands {
		cmd := fmt.Sprintf("%s:", cog.cmd)
		helper += fmt.Sprintf("\t%-12s %s\n", cmd, cog.helper)
	}
	helper += "\n```"

	resp := &Response{
		Message:         helper,
		Channel:         cmd.Channel,
		Timestamp:       cmd.Timestamp,
		ThreadTimestamp: cmd.ThreadTimestamp,
	}
	b.sendResponse(resp)
}
