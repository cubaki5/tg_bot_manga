package list_use_case

import (
	"fmt"

	"tgbot/internal/models"
	"tgbot/internal/models/models_types"
)

//go:generate mockgen --source=list_use_case.go --destination=mocks/mock_list_use_case.go --package=mocks

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
	return getMessage(mangaDB)
}

func getMessage(mangaDB map[models_types.TitleID]models.Title) string {
	if len(mangaDB) == 0 {
		return listIsEmpty
	}

	msg := fmt.Sprintf(headTemplate, msgHead)
	msg += getMessageBody(mangaDB)
	return msg
}

func getMessageBody(mangaDB map[models_types.TitleID]models.Title) string {
	var msg string
	for _, title := range mangaDB {
		msg = msg + fmt.Sprintf(bodyTemplate, title.Name, title.URL)
	}
	return msg
}
