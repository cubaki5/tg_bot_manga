package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/labstack/gommon/log"
	"os"
)

type Handler interface {
	Handle(targetInfo TGMessageInfo) (string, error)
}

type TelegramBot struct {
	bot               *tgbotapi.BotAPI
	handlers          map[string]Handler
	notExistedCommand Handler
}

func NewTGBot(handlers map[string]Handler, notExistedCommand Handler) *TelegramBot {

	botToken := os.Getenv("TG_BOT_TOKEN")

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Error(err)
	}

	return &TelegramBot{
		bot:               bot,
		handlers:          handlers,
		notExistedCommand: notExistedCommand,
	}
}

func (t *TelegramBot) Run() {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 3
	updateConfig.Limit = 50

	updatesChan := t.bot.GetUpdatesChan(updateConfig)

	for update := range updatesChan {
		go func(update tgbotapi.Update) {
			t.updateHandle(update, update.Message.Command())
		}(update)
	}
}

func (t *TelegramBot) updateHandle(update tgbotapi.Update, command string) {
	var (
		msg string
		err error
	)

	tgInfo := NewTGMessageInfo(TitleInfo(update.Message.CommandArguments()))

	if _, ok := t.handlers[command]; !ok {
		msg, err = t.notExistedCommand.Handle(tgInfo)
	} else {
		msg, err = t.handlers[command].Handle(tgInfo)
	}

	chattable := tgbotapi.NewMessage(update.Message.Chat.ID, msg)

	_, err = t.bot.Send(chattable)
	if err != nil {
		log.Error(err)
	}
}
