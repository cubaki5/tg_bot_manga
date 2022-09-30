package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/labstack/gommon/log"
	"strings"
)

type Handler interface {
	Handle(update tgbotapi.Update) (tgbotapi.Chattable, error)
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

	updatesChan := t.bot.GetUpdatesChan(updateConfig)

	for update := range updatesChan {
		go func(update tgbotapi.Update) {
			command := strings.Split(update.Message.Text, " ")
			switch {
			case command[0] == "/start":
				t.updateHandle(update, command[0])
			case command[0] == "/add":
				t.updateHandle(update, command[0])
			case command[0] == "/delete":
				t.updateHandle(update, command[0])
			case command[0] == "/list":
				t.updateHandle(update, command[0])
			default:
				t.updateHandle(update, "")
			}
		}(update)
	}
}

func (t *TelegramBot) updateHandle(update tgbotapi.Update, command string) {

	if command == "" {
		chattable := tgbotapi.NewMessage(update.Message.Chat.ID, "You should write command from list")

		_, err := t.bot.Send(chattable)
		if err != nil {
			log.Error(err)
		}

	} else {
		chattable, err := t.handlers[command].Handle(update)
		if err != nil {
			log.Error(err)
		}

		_, err = t.bot.Send(chattable)
		if err != nil {
			log.Error(err)
		}
	}
}
