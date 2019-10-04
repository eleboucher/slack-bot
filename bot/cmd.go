package bot

import (
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
	cmd         string
	helper      string
	description string
	function    cmdFunc
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
func RegisterCommand(cmd, helper, description string, function cmdFunc) {
	log.Infof("Adding Command %s", cmd)
	commands[cmd] = &Cog{
		cmd:         cmd,
		helper:      helper,
		description: description,
		function:    function,
	}
}

func RegisterPeriodicCommand(cmd string, config *PeriodicCog) {
	log.Infof("Adding Periodic Command %s", cmd)

	periodicCommand[cmd] = config
}

func (b *Bot) handleCMD(cmd *CMD) {
	c := commands[cmd.Command]

	log.Infof("received new Command %#v", cmd)

	if c == nil {
		log.Errorf("Command %s not found\n", cmd.Command)
	}
	log.Info(c)

	resp, err := c.function(cmd)

	if err != nil {
		log.Errorf("Command %s error: %s\n", cmd.Command, err)
	}

	b.sendResponse(resp)
}
