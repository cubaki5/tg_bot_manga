package add_use_case

import (
	"tgbot/internal/models"
)

//go:generate mockgen --source=add_use_case.go --destination=mocks/mock_add_use_case.go --package=mocks

type Database interface {
	Add(title models.Title)
}

type AddUseCase struct {
	db Database
}

func NewAddUseCase(db Database) *AddUseCase {
	return &AddUseCase{db: db}
}

func (s *AddUseCase) Add(title models.Title) {
	s.db.Add(title)
}
