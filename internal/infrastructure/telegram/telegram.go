package telegram

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/labstack/gommon/log"
)

type Handler interface {
	Handle(targetInfo TGMessageInfo) (string, error)
}

type Bot struct {
	bot               *tgbotapi.BotAPI
	handlers          map[string]Handler
	notExistedCommand Handler
}

func NewTGBot(handlers map[string]Handler, notExistedCommand Handler) *Bot {

	botToken := os.Getenv("TG_BOT_TOKEN")

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Error(err)
	}

	return &Bot{
		bot:               bot,
		handlers:          handlers,
		notExistedCommand: notExistedCommand,
	}
}

func (t *Bot) Run() {
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

func (t *Bot) updateHandle(update tgbotapi.Update, command string) {
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
	if err != nil {
		log.Error(err)
	}

	chattable := tgbotapi.NewMessage(update.Message.Chat.ID, msg)

	_, err = t.bot.Send(chattable)
	if err != nil {
		log.Error(err)
	}
}
