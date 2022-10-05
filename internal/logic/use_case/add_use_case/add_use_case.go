package add_use_case

import (
	"errors"

	"tgbot/internal/models"
)

type Database interface {
	Add(title models.Title)
}

type AddUseCase struct {
	db Database
}

func NewAddUseCase(db Database) *AddUseCase {
	return &AddUseCase{db: db}
}

func (s *AddUseCase) Add(title models.Title) error {

	switch {
	case title.Name == "":
		return errors.New("title Name is empty")
	case title.URL == "":
		return errors.New("title URL is empty")
	}

	s.db.Add(title)
	return nil
}
