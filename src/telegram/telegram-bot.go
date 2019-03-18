package main

import (
	"log"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	config Configuration
	bot    *tb.Bot
)

func Init() {
	config = Load("config/config.json")
	var err error
	bot, err = tb.NewBot(tb.Settings{
		Token:  config.Token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	Init()
	DefineBasicCommands()
	DefineCommands()
	bot.Start()
}
