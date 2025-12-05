package main

import (
	"log"
	"telegram_bot/internal/config"
	"telegram_bot/internal/telegram"
)

func main() {
	cfg := config.MustLoad()

	bot, err := telegram.NewBot(cfg)
	if err != nil {
		log.Fatal("Failed to create bot:", err)
	}

	if err := bot.Start(); err != nil {
		log.Fatal("Failed to start bot:", err)
	}
}
