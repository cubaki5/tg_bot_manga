package not_existed_command_handler

import (
	"tgbot/internal/infrastructure/telegram/telegram_models"
)

type UseCase interface {
	NotExistedCommand() string
}

type NotExistedCommandHandler struct {
	uc UseCase
}

func NewNotExistedCommandHandler(uc UseCase) *NotExistedCommandHandler {
	return &NotExistedCommandHandler{uc: uc}
}

func (nc NotExistedCommandHandler) Handle(targetInfo telegram_models.TGMessageInfo) (string, error) {
	return nc.uc.NotExistedCommand(), nil
}
