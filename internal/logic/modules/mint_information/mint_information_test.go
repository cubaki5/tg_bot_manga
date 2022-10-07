package mint_information

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"tgbot/internal/logic/modules/mint_information/mocks"
	"tgbot/internal/logic/modules/mint_information/models"
	"tgbot/internal/models/models_types"
)

func initClientMock(t *testing.T) *mocks.MockWebClient {
	ctr := gomock.NewController(t)
	return mocks.NewMockWebClient(ctr)
}

func initParserMock(t *testing.T) *mocks.MockParser {
	ctr := gomock.NewController(t)
	return mocks.NewMockParser(ctr)
}

func TestGetTitleModule_GetTitle(t *testing.T) {
	t.Run("Happy path", func(t *testing.T) {
		url := models_types.URL("TestUrl")
		b := []byte("TestBytes")
		expTitleName := models_types.TitleName("TestName")
		expLastUPD := time.Unix(0, 0)

		mockClient := initClientMock(t)
		mockParser := initParserMock(t)

		mockClient.EXPECT().DoGetRequest(url).Return(b, nil)
		mockParser.EXPECT().Parse(b).Return(models.TitleParams{
			Name:    expTitleName,
			LastUPD: expLastUPD,
		}, nil)

		gtModule := NewGetTitleModule(mockClient, mockParser)

		actParams, err := gtModule.GetTitle(url)

		assert.NoError(t, err)

		t.Run("GetTitle returns correct title name", func(t *testing.T) {
			require.Equal(t, expTitleName, actParams.Name)
		})
		t.Run("GetTitle returns correct lastUDP", func(t *testing.T) {
			require.Equal(t, expLastUPD, actParams.LastUpdate)
		})
	})

	t.Run("Client returns error", func(t *testing.T) {
		mockClient := initClientMock(t)
		mockParser := initParserMock(t)

		mockClient.EXPECT().DoGetRequest(gomock.Any()).Return([]byte("TestBytes"), errors.New("TestError"))

		gtModule := NewGetTitleModule(mockClient, mockParser)

		_, err := gtModule.GetTitle("TestUrl")

		assert.EqualError(t, err, "TestError")
	})

	t.Run("Parser returns error", func(t *testing.T) {
		mockClient := initClientMock(t)
		mockParser := initParserMock(t)

		mockClient.EXPECT().DoGetRequest(gomock.Any()).Return([]byte("TestBytes"), nil)
		mockParser.EXPECT().Parse([]byte("TestBytes")).Return(models.TitleParams{}, errors.New("TestError"))

		gtModule := NewGetTitleModule(mockClient, mockParser)

		_, err := gtModule.GetTitle("TestUrl")

		assert.EqualError(t, err, "TestError")
	})
}
