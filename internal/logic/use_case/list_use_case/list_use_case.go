package list_use_case

import (
	"fmt"

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

func (s *ListUseCase) List() string {
	mangaDB := s.db.List()

	if len(mangaDB) == 0 {
		return `
			Your title list is empty.
			You can add title via command
			/add <title url>`
	}

	msg := "Your titles:\n"
	for _, title := range mangaDB {
		msg = msg + fmt.Sprintf("%d) %s - %s\n", title.ID, title.Name, title.URL)
	}
	return msg
}
