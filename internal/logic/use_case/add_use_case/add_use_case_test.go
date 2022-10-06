package add_use_case

import (
	"testing"

	"github.com/golang/mock/gomock"

	"tgbot/internal/logic/use_case/add_use_case/mocks"
	"tgbot/internal/models"
	"tgbot/internal/models/models_types"
)

func initMock(t *testing.T) *mocks.MockDatabase {
	ctr := gomock.NewController(t)
	return mocks.NewMockDatabase(ctr)
}

func TestAddUseCase_Add(t *testing.T) {
	t.Run("Test add func gets the same title", func(t *testing.T) {
		testName := models_types.TitleName("Test Name")
		testURL := models_types.URL("TestURL")
		testTitle := models.Title{
			URL:  testURL,
			Name: testName,
		}
		mockDB := initMock(t)
		mockDB.EXPECT().Add(testTitle)

		addUC := NewAddUseCase(mockDB)
		addUC.Add(testTitle)
	})
}
