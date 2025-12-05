package telegram

import (
	"log"
	"telegram_bot/internal/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot     *tgbotapi.BotAPI
	config  config.Config
	handler *Handler
}

func NewBot(cfg config.Config) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(cfg.Telegram.Token)
	if err != nil {
		return nil, err
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &Bot{
		bot:     bot,
		config:  cfg,
		handler: NewHandler(bot),
	}, nil
}

func (b *Bot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			b.handler.HandleMessage(update.Message)
		} else if update.CallbackQuery != nil {
			b.handler.HandleCallback(update.CallbackQuery)
		}
	}

	return nil
}
