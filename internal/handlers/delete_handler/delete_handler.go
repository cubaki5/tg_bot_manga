package delete_handler

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type UseCase interface {
	Delete() error
}

type DeleteHandler struct {
	uc UseCase
}

func NewDeleteHandler(uc UseCase) *DeleteHandler {
	return &DeleteHandler{uc: uc}
}

func (dh *DeleteHandler) Handle(update tgbotapi.Update) (tgbotapi.Chattable, error) {
	return nil, nil
}
