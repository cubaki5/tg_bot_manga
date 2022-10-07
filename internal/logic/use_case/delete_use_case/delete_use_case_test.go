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
		msg, err := deleteUC.Delete(testName)

		assert.NoError(t, err)
		assert.Equal(t, "Title <TestName> is deleted", msg)
	})
	t.Run("Delete func returns error", func(t *testing.T) {
		t.Run("Delete func returns the same error as method of interface", func(t *testing.T) {
			testName := models_types.TitleName("TestName")
			mockDB := initMock(t)
			mockDB.EXPECT().Delete(gomock.Any()).Return(errors.New("test error"))

			deleteUC := NewDeleteUseCase(mockDB)
			msg, err := deleteUC.Delete(testName)

			assert.EqualError(t, err, "test error")
			assert.Equal(t, "Doesn't know title <TestName>, try another name", msg)
		})
	})
}
