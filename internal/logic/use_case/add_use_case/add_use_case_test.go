package add_use_case

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"tgbot/internal/logic/use_case/add_use_case/mocks"
	"tgbot/internal/models"
	"tgbot/internal/models/models_types"
)

func initMockDB(t *testing.T) *mocks.MockDatabase {
	ctr := gomock.NewController(t)
	return mocks.NewMockDatabase(ctr)
}

func initMockGetTitle(t *testing.T) *mocks.MockGetTitleModules {
	ctr := gomock.NewController(t)
	return mocks.NewMockGetTitleModules(ctr)
}

func TestAddUseCase_Add(t *testing.T) {
	t.Run("Happy Path", func(t *testing.T) {
		testName := models_types.TitleName("Test Name")
		testURL := models_types.URL("TestURL")
		testTitle := models.Title{
			URL:  testURL,
			Name: testName,
		}

		mockDB := initMockDB(t)
		mockGetTitle := initMockGetTitle(t)

		mockGetTitle.EXPECT().GetTitle(testURL).Return(testTitle, nil)
		mockDB.EXPECT().Add(testTitle)

		addUC := NewAddUseCase(mockDB, mockGetTitle)
		msg, err := addUC.Add(testURL)
		assert.NoError(t, err)
		assert.Equal(t, "Title with name <Test Name> was added", msg)
	})
	t.Run("GetTitle returns err", func(t *testing.T) {

		mockDB := initMockDB(t)
		mockGetTitle := initMockGetTitle(t)

		mockGetTitle.EXPECT().GetTitle(gomock.Any()).Return(models.Title{}, errors.New("TestError"))

		addUC := NewAddUseCase(mockDB, mockGetTitle)
		msg, err := addUC.Add("testURL")
		assert.EqualError(t, err, "TestError")
		assert.Equal(t, notExistedURL, msg)
	})
}
