package start_handler

import (
	"tgbot/internal/infrastructure/telegram"
)

type UseCase interface {
	Start() string
}

type StartHandler struct {
	uc UseCase
}

func NewStartHandler(uc UseCase) *StartHandler {
	return &StartHandler{uc: uc}
}

func (sh StartHandler) Handle(targetInfo telegram.TGMessageInfo) (string, error) {
	return sh.uc.Start(), nil
}
