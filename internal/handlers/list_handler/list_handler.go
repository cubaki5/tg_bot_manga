package list_handler

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type UseCase interface {
	List() (interface{}, error)
}

type ListHandler struct {
	uc UseCase
}

func NewListHandler(uc UseCase) *ListHandler {
	return &ListHandler{uc: uc}
}

func (gh *ListHandler) Handle(update tgbotapi.Update) (tgbotapi.Chattable, error) {
	return nil, nil
}
