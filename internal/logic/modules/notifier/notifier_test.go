package notifier

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"tgbot/internal/logic/modules/mint_information/models"
	"tgbot/internal/logic/modules/notifier/mocks"
	models2 "tgbot/internal/models"
	"tgbot/internal/models/models_types"
)

func initMockWebClient(t *testing.T) *mocks.MockWebClient {
	ctr := gomock.NewController(t)
	return mocks.NewMockWebClient(ctr)
}

func initMockTGClient(t *testing.T) *mocks.MockTGClient {
	ctr := gomock.NewController(t)
	return mocks.NewMockTGClient(ctr)
}

func initMockParser(t *testing.T) *mocks.MockParser {
	ctr := gomock.NewController(t)
	return mocks.NewMockParser(ctr)
}

func initMockDatabase(t *testing.T) *mocks.MockDatabase {
	ctr := gomock.NewController(t)
	return mocks.NewMockDatabase(ctr)
}

func TestNotifier_IsUpdated(t *testing.T) {
	t.Run("Happy path", func(t *testing.T) {
		happyPathTests := []struct {
			testName         string
			expTitleName     models_types.TitleName
			url              models_types.URL
			lastUPD          time.Time
			updatedTitleName models_types.TitleName
			updatedLastUPD   time.Time
			b                []byte
			errorTGClient    error
			errorParser      error
			errorIsUPD       error
			expIsUPD         bool
		}{
			{
				"Is not updated",
				"TestName",
				"TestURL",
				time.Unix(1, 0),
				"TestName",
				time.Unix(1, 0),
				[]byte("TestBytes"),
				nil,
				nil,
				nil,
				false,
			},
			{
				"Is updated",
				"TestName",
				"TestURL",
				time.Unix(1, 0),
				"TestName",
				time.Unix(3, 0),
				[]byte("TestBytes"),
				nil,
				nil,
				nil,
				true,
			},
		}
		for _, happyPathTest := range happyPathTests {
			t.Run(happyPathTest.testName, func(t *testing.T) {
				mockWebClient := initMockWebClient(t)
				mockTGClient := initMockTGClient(t)
				mockParser := initMockParser(t)
				mockDB := initMockDatabase(t)

				mockWebClient.EXPECT().DoGetRequest(happyPathTest.url).Return(happyPathTest.b, happyPathTest.errorTGClient)
				mockParser.EXPECT().Parse(happyPathTest.b).Return(models.TitleParams{
					Name:    happyPathTest.updatedTitleName,
					LastUPD: happyPathTest.updatedLastUPD,
				}, happyPathTest.errorParser)

				not := NewNotifier(mockWebClient, mockDB, mockTGClient, mockParser)

				isUPD, err := not.IsUpdated(models2.Title{
					ID:         1,
					Name:       happyPathTest.expTitleName,
					URL:        happyPathTest.url,
					LastUpdate: happyPathTest.lastUPD,
				})

				assert.NoError(t, err)
				assert.Equal(t, happyPathTest.expIsUPD, isUPD)

			})
		}

	})
	t.Run("IsUpdated returns error", func(t *testing.T) {
		tests := []struct {
			testName         string
			expTitleName     models_types.TitleName
			url              models_types.URL
			lastUPD          time.Time
			updatedTitleName models_types.TitleName
			updatedLastUPD   time.Time
			b                []byte
			errorTGClient    error
			errorParser      error
			errorIsUPD       error
		}{
			{
				"TGClient returns error",
				"TestName",
				"TestURL",
				time.Unix(3, 0),
				"TestName",
				time.Unix(1, 0),
				[]byte("TestBytes"),
				errors.New("TestError"),
				nil,
				errors.New("TestError"),
			},
			{
				"Parser returns error",
				"TestName",
				"TestURL",
				time.Unix(3, 0),
				"TestName",
				time.Unix(1, 0),
				[]byte("TestBytes"),
				nil,
				errors.New("TestError"),
				errors.New("TestError"),
			},
			{
				"Different names",
				"TestName",
				"TestURL",
				time.Unix(3, 0),
				"WrongTestName",
				time.Unix(1, 0),
				[]byte("TestBytes"),
				nil,
				nil,
				errors.New("names mismatching"),
			},
		}

		for _, test := range tests {
			t.Run(test.testName, func(t *testing.T) {
				mockWebClient := initMockWebClient(t)
				mockTGClient := initMockTGClient(t)
				mockParser := initMockParser(t)
				mockDB := initMockDatabase(t)

				mockWebClient.EXPECT().DoGetRequest(test.url).Return(test.b, test.errorTGClient)
				if test.errorTGClient == nil {
					mockParser.EXPECT().Parse(test.b).Return(models.TitleParams{
						Name:    test.updatedTitleName,
						LastUPD: test.updatedLastUPD,
					}, test.errorParser)
				}
				not := NewNotifier(mockWebClient, mockDB, mockTGClient, mockParser)

				isUPD, err := not.IsUpdated(models2.Title{
					ID:         1,
					Name:       test.expTitleName,
					URL:        test.url,
					LastUpdate: test.lastUPD,
				})

				assert.EqualError(t, err, test.errorIsUPD.Error())
				assert.True(t, !isUPD)
			})
		}
	})
}
