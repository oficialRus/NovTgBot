package telegram

import (
	"log"
	"telegram_bot/internal/usecases/start"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler struct {
	bot          *tgbotapi.BotAPI
	startUseCase *start.UseCase
}

func NewHandler(bot *tgbotapi.BotAPI) *Handler {
	return &Handler{
		bot:          bot,
		startUseCase: start.NewUseCase(),
	}
}

func (h *Handler) HandleMessage(message *tgbotapi.Message) {
	if message.IsCommand() {
		switch message.Command() {
		case "start":
			h.startUseCase.HandleStart(h.bot, message.Chat.ID)
		default:
			log.Printf("Unknown command: %s", message.Command())
		}
	}
}

func (h *Handler) HandleCallback(callback *tgbotapi.CallbackQuery) {
	// Обработка нажатий на inline кнопки
	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "Вы нажали на кнопку: "+callback.Data)
	h.bot.Send(msg)

	// Отвечаем на callback, чтобы убрать индикатор загрузки
	callbackConfig := tgbotapi.NewCallback(callback.ID, "")
	h.bot.Request(callbackConfig)
}
