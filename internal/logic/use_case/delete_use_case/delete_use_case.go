package delete_use_case

import "tgbot/internal/models/models_types"

type Database interface {
	Delete(titleName models_types.TitleName) error
}

type DeleteUseCase struct {
	db Database
}

func NewDeleteUseCase(db Database) *DeleteUseCase {
	return &DeleteUseCase{db: db}
}

func (s *DeleteUseCase) Delete() error {
	return nil
}
