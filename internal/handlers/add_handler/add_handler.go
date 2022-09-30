package add_handler

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type UseCase interface {
	Add() error
}

type AddHandler struct {
	uc UseCase
}

func NewAddHandler(uc UseCase) *AddHandler {
	return &AddHandler{uc: uc}
}

func (a AddHandler) Handle(update tgbotapi.Update) (tgbotapi.Chattable, error) {
	//TODO implement me
	panic("implement me")
}
