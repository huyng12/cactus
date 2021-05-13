package handler

import (
	"fmt"

	"github.com/huyng12/cactus/internal/wallet"
	"gopkg.in/tucnak/telebot.v2"
)

func CheckBalance(bot *telebot.Bot, w wallet.Walleter) func(msg *telebot.Message) {
	return func(msg *telebot.Message) {
		currentBalance := w.GetBalance()
		_, _ = bot.Send(msg.Chat, fmt.Sprintf("💵 Balance: %d", currentBalance))
	}
}
