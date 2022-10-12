package telegram

import (
	"fmt"
	"os"

	"tgbot/internal/models"

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

func NewTGBot() *Bot {

	botToken := os.Getenv("TG_BOT_TOKEN")

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Error(err)
	}

	return &Bot{
		bot: bot,
	}
}

func (t *Bot) AppendBotWithHandlers(handlers map[string]Handler, notExistedCommand Handler) {
	t.handlers = handlers
	t.notExistedCommand = notExistedCommand
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

	tgInfo := NewTGMessageInfo(TitleInfo(update.Message.CommandArguments()), ChatID(update.Message.Chat.ID))

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

func (t *Bot) PostMsg(title models.Title, user models.User) error {
	msg := fmt.Sprintf(msgTemplate, title.Name, title.URL)
	chattable := tgbotapi.NewMessage(user.ID.Int64(), msg)

	_, err := t.bot.Send(chattable)
	if err != nil {
		return err
	}
	return nil
}
