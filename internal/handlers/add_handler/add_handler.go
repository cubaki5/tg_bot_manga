package add_handler

import (
	"tgbot/internal/infrastructure/telegram"
	"tgbot/internal/models/models_types"
)

type AddUseCase interface {
	Add(URL models_types.URL) (string, error)
}

type AddHandler struct {
	AddUC AddUseCase
}

func NewAddHandler(addUC AddUseCase) *AddHandler {
	return &AddHandler{
		AddUC: addUC,
	}
}

func (a AddHandler) Handle(targetInfo telegram.TGMessageInfo) (string, error) {
	URL := models_types.URL(targetInfo.Title)
	return a.AddUC.Add(URL)
}
