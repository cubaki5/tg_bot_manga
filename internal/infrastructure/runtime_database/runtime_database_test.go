package runtime_database

import (
	"fmt"
	"testing"
	"tgbot/internal/models"
	"tgbot/internal/models/models_types"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRuntimeDatabase(t *testing.T) {
	t.Run("Testing of List func", func(t *testing.T) {
		database := NewDatabase()
		testTitle := models.Title{
			TitleName:  "TestName",
			URL:        "TestURL",
			LastUpdate: time.Now(),
		}

		database.Add(testTitle)

		expMap := map[models_types.TitleID]models.Title{
			1: testTitle,
		}
		actMap := database.List()

		require.Equal(t, expMap, actMap)
	})

	t.Run("Testing Add func", func(t *testing.T) {
		database := NewDatabase()
		testTitle := models.Title{
			TitleName:  "TestName",
			URL:        "TestURL",
			LastUpdate: time.Now(),
		}
		expIncr := database.incr + 1
		database.Add(testTitle)

		require.Equal(t, expIncr, database.incr)
		require.Equal(t, testTitle, database.db[expIncr-1])
	})

	t.Run("Testing Delete func", func(t *testing.T) {
		t.Run("Delete title with correct name", func(t *testing.T) {
			database := NewDatabase()
			testTitle := models.Title{
				TitleName:  "TestName",
				URL:        "TestURL",
				LastUpdate: time.Now(),
			}
			database.Add(testTitle)
			expLen := len(database.db) - 1

			err := database.Delete(testTitle.TitleName)
			fmt.Println(database.db)
			require.NoError(t, err)
			require.Equal(t, expLen, len(database.db))
		})
		t.Run("Delete title with wrong name", func(t *testing.T) {
			database := NewDatabase()
			testTitle := models.Title{
				TitleName:  "TestName",
				URL:        "TestURL",
				LastUpdate: time.Now(),
			}
			database.Add(testTitle)
			expLen := len(database.db)
			var wrongTestName models_types.TitleName = "wrongTestName"

			err := database.Delete(wrongTestName)

			require.EqualError(t, err, "title with such name doesn't exist in database")
			require.Equal(t, expLen, len(database.db))
		})
	})
}
