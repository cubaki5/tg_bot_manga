package delete_handler

import (
	"tgbot/internal/infrastructure/telegram"
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

func (dh *DeleteHandler) Handle(targetInfo telegram.TGMessageInfo) (string, error) {
	return "", nil
}
