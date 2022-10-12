package telegram

import (
	"errors"
	"fmt"
	"os"
	"strconv"

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

func NewTGBot() (*Bot, error) {

	botToken := os.Getenv("TG_BOT_TOKEN")
	chatID := os.Getenv("CHAT_ID")

	if botToken == "" || chatID == "" {
		return nil, errors.New("fail to get environment value")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return nil, err
	}

	return &Bot{
		bot: bot,
	}, nil
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

	msg, err = t.RequestRouter(tgInfo, command)
	if err != nil {
		log.Error(err)
	}

	chattable := tgbotapi.NewMessage(update.Message.Chat.ID, msg)

	_, err = t.bot.Send(chattable)
	if err != nil {
		log.Error(err)
	}
}

func (t *Bot) RequestRouter(tgInfo TGMessageInfo, command string) (string, error) {

	auth, err := authenticationMiddleware(tgInfo)
	if err != nil {
		return authenticationFail, err
	}

	if !auth {
		return notDefinedUser, nil
	}

	if _, ok := t.handlers[command]; !ok {
		msg, _ := t.notExistedCommand.Handle(tgInfo) //notExistedCommand does not return an error under any circumstances
		return msg, nil
	}

	msg, err := t.handlers[command].Handle(tgInfo)

	return msg, err
}

func authenticationMiddleware(tgInfo TGMessageInfo) (bool, error) {
	ID, err := strconv.Atoi(os.Getenv("CHAT_ID"))
	if err != nil {
		return false, err
	}

	if tgInfo.ID != ChatID(ID) {
		return false, nil
	}

	return true, nil
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
