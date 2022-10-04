package list_use_case

import (
	"tgbot/internal/models"
	"tgbot/internal/models/models_types"
)

type Database interface {
	List() map[models_types.TitleID]models.Title
}

type ListUseCase struct {
	db Database
}

func NewListUseCase(db Database) *ListUseCase {
	return &ListUseCase{db: db}
}

func (s *ListUseCase) List() map[models_types.TitleID]models.Title {
	return nil
}
