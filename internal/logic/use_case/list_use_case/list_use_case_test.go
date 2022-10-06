package list_use_case

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"tgbot/internal/logic/use_case/list_use_case/mocks"
	"tgbot/internal/models"
	"tgbot/internal/models/models_types"
)

func initMock(t *testing.T) *mocks.MockDatabase {
	ctr := gomock.NewController(t)
	return mocks.NewMockDatabase(ctr)
}

func TestListUseCase_List(t *testing.T) {
	t.Run("List func gets empty map", func(t *testing.T) {
		testMangaDB := map[models_types.TitleID]models.Title{}

		mockDB := initMock(t)
		mockDB.EXPECT().List().Return(testMangaDB)

		listUC := NewListUseCase(mockDB)
		actMsg := listUC.List()

		expMsg := listIsEmpty

		assert.Equal(t, expMsg, actMsg)
	})
	t.Run("List func gets normal map", func(t *testing.T) {
		testMangaDB := map[models_types.TitleID]models.Title{
			1: {
				ID:   models_types.TitleID(1),
				Name: models_types.TitleName("TestName"),
				URL:  models_types.URL("TestUrl"),
			},
		}
		expMsg := `Your titles:
— TestName - TestUrl
`

		mockDB := initMock(t)
		mockDB.EXPECT().List().Return(testMangaDB)

		listUC := NewListUseCase(mockDB)
		actMsg := listUC.List()

		assert.Equal(t, expMsg, actMsg)
	})
}
