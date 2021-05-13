package handler

import (
	"gopkg.in/tucnak/telebot.v2"
)

func HealthCheck(bot *telebot.Bot) func(msg *telebot.Message) {
	return func(msg *telebot.Message) {
		_, _ = bot.Send(msg.Chat, "I'm here 🤨")
	}
}
