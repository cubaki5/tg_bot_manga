package delete_use_case

import (
	"fmt"

	"tgbot/internal/models/models_types"
)

//go:generate mockgen --source=delete_use_case.go --destination=mocks/mock_add_use_case.go --package=mocks

type Database interface {
	Delete(titleName models_types.TitleName) error
}

type DeleteUseCase struct {
	db Database
}

func NewDeleteUseCase(db Database) *DeleteUseCase {
	return &DeleteUseCase{db: db}
}

func (s *DeleteUseCase) Delete(titleName models_types.TitleName) (string, error) {
	err := s.db.Delete(titleName)
	if err != nil {
		return fmt.Sprintf(deleteErr, titleName), err
	}
	return fmt.Sprintf(successDel, titleName), nil
}
