package start

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type UseCase struct {
}

func NewUseCase() *UseCase {
	return &UseCase{}
}

func (uc *UseCase) HandleStart(bot *tgbotapi.BotAPI, chatID int64) {
	// Используем путь из корня проекта: go run запускается из /home/.../NovTgBot
	photo := tgbotapi.NewPhoto(chatID, tgbotapi.FilePath("cmd/photo/fhoto.jpeg"))
	photo.Caption = "Добро пожаловать! Выберите действие:"

	// Создаем inline клавиатуру с 3 кнопками
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("Написать Руслану", "https://t.me/RyslanNovikov"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("Канал", "https://t.me/novikovpromarket"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("Отзывы по работе", "https://t.me/novikovpromarket"),
		),
	)

	photo.ReplyMarkup = keyboard
	bot.Send(photo)
}
