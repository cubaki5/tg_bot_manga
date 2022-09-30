package list_handler

import (
	"tgbot/internal/infrastructure/telegram/telegram_models"
)

type UseCase interface {
	List() (interface{}, error)
}

type ListHandler struct {
	uc UseCase
}

func NewListHandler(uc UseCase) *ListHandler {
	return &ListHandler{uc: uc}
}

func (gh *ListHandler) Handle(targetInfo telegram_models.TGMessageInfo) (string, error) {
	return "", nil
}
