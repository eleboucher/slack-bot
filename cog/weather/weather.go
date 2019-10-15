package weather

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/genesixx/slack-bot/bot"
)

func weather(cmd *bot.CMD) (*bot.Response, error) {
	var lat = "48.90"
	var lon = "2.32"

	var options string
	if cmd.Option != "" {
		options = strings.ReplaceAll(cmd.Option, " ", "+")
	} else {
		options = lat + "," + lon
	}

	res, err := http.Get("http://en.wttr.in/" + options + "?T0")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}
	meteo := doc.Find("pre").Text()

	return &bot.Response{
		Message:         "```" + meteo + "```",
		Channel:         cmd.Channel,
		Timestamp:       cmd.Timestamp,
		ThreadTimestamp: cmd.ThreadTimestamp,
	}, nil
}

func init() {
	bot.RegisterCommand("weather", "give the weather to the given location", weather)
}
