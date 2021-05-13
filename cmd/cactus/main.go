package main

import (
	"log"
	"time"

	"github.com/huyng12/cactus/config"

	"github.com/huyng12/cactus/internal/handler"
	"github.com/huyng12/cactus/internal/wallet"
	"gopkg.in/tucnak/telebot.v2"
)

func main() {
	cfg := config.Load()

	// Intitate a wallet for doing transactions
	wlt := wallet.New()

	// Intitate a bot for processing commands
	bot, err := telebot.NewBot(telebot.Settings{
		Token:       cfg.TelegramBotToken,
		Poller:      &telebot.LongPoller{Timeout: 2 * time.Second},
		Synchronous: true,
	})
	if err != nil {
		log.Fatalf("can not awake bot, err = %s", err.Error())
	}

	log.Print("Bot is ready [o_o]")

	bot.Handle("/health_check", handler.HealthCheck(bot))
	bot.Handle("/balance", handler.CheckBalance(bot, wlt))
	bot.Start()
}
