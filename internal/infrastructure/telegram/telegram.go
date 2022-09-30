package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/labstack/gommon/log"
	"strings"
	"tgbot/internal/infrastructure/telegram/telegram_models"
)

type Handler interface {
	Handle(targetInfo telegram_models.TGMessageInfo) (string, error)
}

type TelegramBot struct {
	bot      *tgbotapi.BotAPI
	handlers map[string]Handler
}

func NewTGBot(handlers map[string]Handler) *TelegramBot {

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Error(err)
	}

	return &TelegramBot{
		bot:      bot,
		handlers: handlers,
	}
}

func (t *TelegramBot) Run() {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 3
	updateConfig.Limit = 50

	updatesChan := t.bot.GetUpdatesChan(updateConfig)

	for update := range updatesChan {
		go func(update tgbotapi.Update) {
			command := strings.Split(update.Message.Text, " ")
			t.updateHandle(update, command[0])
		}(update)
	}
}

func (t *TelegramBot) updateHandle(update tgbotapi.Update, command string) {

	if _, ok := t.handlers[command]; !ok {
		command = "not_existed_command"
	}

	tgInfo := telegram_models.NewTGMessageInfo(telegram_models.TitleInfo(strings.Replace(update.Message.Text, command, "", -1)))

	msg, err := t.handlers[command].Handle(tgInfo)
	if err != nil {
		log.Error(err)
	}

	chattable := tgbotapi.NewMessage(update.Message.Chat.ID, msg)

	_, err = t.bot.Send(chattable)
	if err != nil {
		log.Error(err)
	}
}
