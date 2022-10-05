package add_handler

import (
	"tgbot/internal/infrastructure/telegram"
	"tgbot/internal/models"
)

type UseCase interface {
	Add(title models.Title) error
}

type AddHandler struct {
	uc UseCase
}

func NewAddHandler(uc UseCase) *AddHandler {
	return &AddHandler{uc: uc}
}

func (a AddHandler) Handle(targetInfo telegram.TGMessageInfo) (string, error) {
	//TODO implement me
	panic("implement me")
}
