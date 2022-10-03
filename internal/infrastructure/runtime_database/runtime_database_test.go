package runtime_database

import (
	"testing"
	"tgbot/internal/models"
	"tgbot/internal/models/models_types"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRuntimeDatabase(t *testing.T) {
	t.Run("Testing Add func", func(t *testing.T) {
		database := NewDatabase()
		testTitle := models.Title{
			TitleName:  "TestName",
			URL:        "TestURL",
			LastUpdate: time.Now(),
		}

		database.Add(testTitle)

		require.Equal(t, database.incr.Uint32(), uint32(2))
		require.Equal(t, database.db[1], testTitle)
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

			err := database.Delete(testTitle.TitleName)

			require.NoError(t, err)
			require.Equal(t, len(database.db), 0)
		})
		t.Run("Delete title with wrong name", func(t *testing.T) {
			database := NewDatabase()
			testTitle := models.Title{
				TitleName:  "TestName",
				URL:        "TestURL",
				LastUpdate: time.Now(),
			}
			database.Add(testTitle)
			var wrongTestName models_types.TitleName = "wrongTestName"

			err := database.Delete(wrongTestName)

			require.EqualError(t, err, "title with such name doesn't exist in database")
			require.Equal(t, len(database.db), 1)
		})
	})

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

		require.Equal(t, actMap, expMap)
	})
}
