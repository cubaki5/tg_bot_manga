package add_use_case

type Database interface {
	Add() error
}

type AddUseCase struct {
	db Database
}

func NewAddUseCase(db Database) *AddUseCase {
	return &AddUseCase{db: db}
}

func (s *AddUseCase) Add() error {
	return nil
}
