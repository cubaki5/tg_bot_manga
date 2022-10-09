package list_handler

import (
	"tgbot/internal/infrastructure/telegram"
)

type UseCase interface {
	List() string
}

type ListHandler struct {
	uc UseCase
}

func NewListHandler(uc UseCase) *ListHandler {
	return &ListHandler{uc: uc}
}

func (gh *ListHandler) Handle(targetInfo telegram.TGMessageInfo) (string, error) {
	return gh.uc.List(), nil
}
