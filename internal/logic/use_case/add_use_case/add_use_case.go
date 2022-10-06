package add_use_case

import (
	"fmt"

	"tgbot/internal/models"
	"tgbot/internal/models/models_types"
)

//go:generate mockgen --source=add_use_case.go --destination=mocks/mock_add_use_case.go --package=mocks

type Database interface {
	Add(title models.Title)
}

type GetTitleModules interface {
	GetTitle(URL models_types.URL) (models.Title, error)
}

type AddUseCase struct {
	db      Database
	modules GetTitleModules
}

func NewAddUseCase(db Database, modules GetTitleModules) *AddUseCase {
	return &AddUseCase{
		db:      db,
		modules: modules,
	}
}

func (s *AddUseCase) Add(URL models_types.URL) (string, error) {

	title, err := s.modules.GetTitle(URL)
	if err != nil {
		return notExistedURL, err
	}

	s.db.Add(title)
	return fmt.Sprintf(successAddition, title.Name), nil
}
