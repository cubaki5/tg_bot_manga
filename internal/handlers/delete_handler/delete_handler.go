package delete_handler

import (
	"tgbot/internal/infrastructure/telegram"
	"tgbot/internal/models/models_types"
)

type UseCase interface {
	Delete(titleName models_types.TitleName) (string, error)
}

type DeleteHandler struct {
	uc UseCase
}

func NewDeleteHandler(uc UseCase) *DeleteHandler {
	return &DeleteHandler{uc: uc}
}

func (dh *DeleteHandler) Handle(targetInfo telegram.TGMessageInfo) (string, error) {
	name := models_types.TitleName(targetInfo.Title)
	return dh.uc.Delete(name)
}
