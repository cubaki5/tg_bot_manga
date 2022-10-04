package add_use_case

import "tgbot/internal/models"

type Database interface {
	Add(title models.Title)
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
