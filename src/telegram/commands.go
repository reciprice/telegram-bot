package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

func DefineBasicCommands() {
	bot.Handle("/start", func(m *tb.Message) {
		if !m.Private() {
			return
		}

		bot.Send(m.Sender, "Hello!", &tb.ReplyMarkup{})
	})
}
