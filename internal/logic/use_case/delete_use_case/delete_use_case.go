package delete_use_case

type Database interface {
	Delete() error
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
