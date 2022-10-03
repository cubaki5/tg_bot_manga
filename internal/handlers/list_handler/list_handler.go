package list_handler

import (
	"tgbot/internal/infrastructure/telegram"
	"tgbot/internal/models"
	"tgbot/internal/models/models_types"
)

type UseCase interface {
	List() map[models_types.TitleID]models.Title
}

type ListHandler struct {
	uc UseCase
}

func NewListHandler(uc UseCase) *ListHandler {
	return &ListHandler{uc: uc}
}

func (gh *ListHandler) Handle(targetInfo telegram.TGMessageInfo) (string, error) {
	return "", nil
}
