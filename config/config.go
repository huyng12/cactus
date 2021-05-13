package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	// TelegramBotToken is access token for listening/messaging in Telegram
	// Generate it by message `/newbot` to @BotFather
	TelegramBotToken string `required:"true" split_words:"true"`
}

func Load() *Config {
	var cfg Config

	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal("Load config failed", err)
	}

	return &cfg
}
