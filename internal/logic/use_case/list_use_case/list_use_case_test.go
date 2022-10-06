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

		expMsg := `Your title list is empty.
You can add title via command
/add <title url>`

		assert.Equal(t, expMsg, actMsg)
	})
	t.Run("List func gets normal map", func(t *testing.T) {
		testMangaDB := map[models_types.TitleID]models.Title{
			1: {
				ID:   models_types.TitleID(1),
				Name: models_types.TitleName("TestName1"),
				URL:  models_types.URL("TestUrl1"),
			},
			2: {
				ID:   models_types.TitleID(2),
				Name: models_types.TitleName("TestName2"),
				URL:  models_types.URL("TestUrl2"),
			},
			3: {
				ID:   models_types.TitleID(3),
				Name: models_types.TitleName("TestName3"),
				URL:  models_types.URL("TestUrl3"),
			},
		}
		expMsg := `Your titles:
1) TestName1 - TestUrl1
2) TestName2 - TestUrl2
3) TestName3 - TestUrl3
`

		mockDB := initMock(t)
		mockDB.EXPECT().List().Return(testMangaDB)

		listUC := NewListUseCase(mockDB)
		actMsg := listUC.List()

		assert.Equal(t, expMsg, actMsg)
	})
}
