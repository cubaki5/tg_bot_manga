package delete_handler

import (
	"tgbot/internal/infrastructure/telegram/telegram_models"
)

type UseCase interface {
	Delete() error
}

type DeleteHandler struct {
	uc UseCase
}

func NewDeleteHandler(uc UseCase) *DeleteHandler {
	return &DeleteHandler{uc: uc}
}

func (dh *DeleteHandler) Handle(targetInfo telegram_models.TGMessageInfo) (string, error) {
	return "", nil
}
