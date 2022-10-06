package add_use_case

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"tgbot/internal/logic/use_case/add_use_case/mocks"
	"tgbot/internal/models"
	"tgbot/internal/models/models_types"
)

func initMock(t *testing.T) *mocks.MockDatabase {
	ctr := gomock.NewController(t)
	return mocks.NewMockDatabase(ctr)
}

func TestAddUseCase_Add(t *testing.T) {
	t.Run("Test add func returns no error", func(t *testing.T) {
		testName := models_types.TitleName("Test Name")
		testURL := models_types.URL("TestURL")
		testTitle := models.Title{
			URL:  testURL,
			Name: testName,
		}
		mockDB := initMock(t)
		mockDB.EXPECT().Add(testTitle)

		addUC := NewAddUseCase(mockDB)
		err := addUC.Add(testTitle)
		assert.NoError(t, err)
	})
	t.Run("Test add func returns error", func(t *testing.T) {
		t.Run("Add func get empty name", func(t *testing.T) {
			testName := models_types.TitleName("")
			testURL := models_types.URL("TestURL")
			testTitle := models.Title{
				URL:  testURL,
				Name: testName,
			}
			mockDB := initMock(t)

			addUC := NewAddUseCase(mockDB)
			err := addUC.Add(testTitle)

			assert.EqualError(t, err, "title Name is empty")
		})
		t.Run("Add func get empty url", func(t *testing.T) {
			testName := models_types.TitleName("TestName")
			testURL := models_types.URL("")
			testTitle := models.Title{
				URL:  testURL,
				Name: testName,
			}
			mockDB := initMock(t)

			addUC := NewAddUseCase(mockDB)
			err := addUC.Add(testTitle)

			assert.EqualError(t, err, "title URL is empty")
		})
	})
}
