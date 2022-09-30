package add_handler

import (
	"tgbot/internal/infrastructure/telegram/telegram_models"
)

type UseCase interface {
	Add() error
}

type AddHandler struct {
	uc UseCase
}

func NewAddHandler(uc UseCase) *AddHandler {
	return &AddHandler{uc: uc}
}

func (a AddHandler) Handle(targetInfo telegram_models.TGMessageInfo) (string, error) {
	//TODO implement me
	panic("implement me")
}
