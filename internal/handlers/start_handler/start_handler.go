package start_handler

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type UseCase interface {
	Start() string
}

type StartHandler struct {
	uc UseCase
}

func NewStartHandler(uc UseCase) *StartHandler {
	return &StartHandler{uc: uc}
}

func (sh StartHandler) Handle(update tgbotapi.Update) (tgbotapi.Chattable, error) {
	startMsg := sh.uc.Start()
	chattable := tgbotapi.NewMessage(update.Message.Chat.ID, startMsg)
	return chattable, nil
}
