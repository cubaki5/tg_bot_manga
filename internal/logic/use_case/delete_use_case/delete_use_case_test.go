package delete_use_case

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"tgbot/internal/logic/use_case/delete_use_case/mocks"
	"tgbot/internal/models/models_types"
)

func initMock(t *testing.T) *mocks.MockDatabase {
	ctr := gomock.NewController(t)
	return mocks.NewMockDatabase(ctr)
}

func TestDeleteUseCase_Delete(t *testing.T) {
	t.Run("Delete func returns no error", func(t *testing.T) {
		testName := models_types.TitleName("TestName")
		mockDB := initMock(t)
		mockDB.EXPECT().Delete(testName).Return(nil)

		deleteUC := NewDeleteUseCase(mockDB)
		err := deleteUC.Delete(testName)

		assert.NoError(t, err)
	})
	t.Run("Delete func returns error", func(t *testing.T) {
		t.Run("Delete func gets empty name", func(t *testing.T) {
			testName := models_types.TitleName("")
			mockDB := initMock(t)

			deleteUC := NewDeleteUseCase(mockDB)
			err := deleteUC.Delete(testName)

			assert.EqualError(t, err, "title Name is empty")
		})
		t.Run("Delete func returns the same error as method of interface", func(t *testing.T) {
			testName := models_types.TitleName("TestName")
			mockDB := initMock(t)
			mockDB.EXPECT().Delete(gomock.Any()).Return(errors.New("test error"))

			deleteUC := NewDeleteUseCase(mockDB)
			err := deleteUC.Delete(testName)

			assert.EqualError(t, err, "test error")
		})
	})
}
