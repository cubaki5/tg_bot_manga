package runtime_database

import (
	"testing"
	"time"

	"tgbot/internal/models"
	"tgbot/internal/models/models_types"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMangaDatabase_List(t *testing.T) {
	t.Run("Testing of List func", func(t *testing.T) {
		t.Run("List return correct map", func(t *testing.T) {
			expID := models_types.TitleID(1)
			expName := models_types.TitleName("TestName")
			expURL := models_types.URL("TestURL")
			expLastUPD := time.Now()

			d := &MangaDatabase{
				db: map[models_types.TitleID]models.Title{
					expID: {
						ID:         expID,
						Name:       expName,
						URL:        expURL,
						LastUpdate: expLastUPD,
					},
				},
				incr: 2,
			}

			actMap := d.List()
			t.Run("List return map with one element", func(t *testing.T) {
				assert.Len(t, actMap, 1)
			})

			actObj, ok := actMap[expID]
			require.True(t, ok)
			t.Run("Correct ID", func(t *testing.T) {
				assert.Equal(t, actObj.ID, expID)
			})
			t.Run("Correct title", func(t *testing.T) {
				assert.Equal(t, actObj.Name, expName)
			})
			t.Run("Correct URL", func(t *testing.T) {
				assert.Equal(t, actObj.URL, expURL)
			})
			t.Run("Correct LastUpdate", func(t *testing.T) {
				assert.Equal(t, actObj.LastUpdate, expLastUPD)
			})
		})
		t.Run("List return empty map", func(t *testing.T) {
			d := &MangaDatabase{
				db:   map[models_types.TitleID]models.Title{},
				incr: 1,
			}

			actMap := d.List()
			t.Run("List return empty map", func(t *testing.T) {
				assert.Len(t, actMap, 0)
			})
		})
	})
}

func TestMangaDatabase_Add(t *testing.T) {
	t.Run("Testing of Add func", func(t *testing.T) {
		d := &MangaDatabase{
			db:   map[models_types.TitleID]models.Title{},
			incr: 1,
		}

		expID := models_types.TitleID(1)
		expName := models_types.TitleName("TestName")
		expURL := models_types.URL("TestURL")
		expLastUPD := time.Now()

		expTitle := models.Title{
			Name:       expName,
			URL:        expURL,
			LastUpdate: expLastUPD,
		}

		d.Add(expTitle)
		actMap := d.db
		t.Run("Add create one element in map", func(t *testing.T) {
			assert.Len(t, actMap, 1)
		})

		actTitle, ok := actMap[expID]
		require.True(t, ok)
		t.Run("Add create element with correct ID", func(t *testing.T) {
			assert.Equal(t, expID, actTitle.ID)
		})
		t.Run("Add create element with correct Name", func(t *testing.T) {
			assert.Equal(t, expName, actTitle.Name)
		})
		t.Run("Add create element with correct URL", func(t *testing.T) {
			assert.Equal(t, expURL, actTitle.URL)
		})
		t.Run("Add create element with correct LastUpdate", func(t *testing.T) {
			assert.Equal(t, expLastUPD, actTitle.LastUpdate)
		})
	})
}

func TestMangaDatabase_Delete(t *testing.T) {
	t.Run("Testing of Delete func", func(t *testing.T) {
		t.Run("Delete title with correct name", func(t *testing.T) {
			expID := models_types.TitleID(1)
			expName := models_types.TitleName("TestName")

			d := &MangaDatabase{
				db: map[models_types.TitleID]models.Title{
					expID: {
						ID:   expID,
						Name: expName,
					},
				},
				incr: 2,
			}

			err := d.Delete(expName)
			require.NoError(t, err)
			t.Run("Delete func remove one element in map", func(t *testing.T) {
				assert.Len(t, d.db, 0)
			})
		})
		t.Run("Delete title with incorrect name", func(t *testing.T) {
			expID := models_types.TitleID(1)
			expName := models_types.TitleName("TestName")
			wrongName := models_types.TitleName("WrongTestName")

			d := &MangaDatabase{
				db: map[models_types.TitleID]models.Title{
					expID: {
						ID:   expID,
						Name: expName,
					},
				},
				incr: 2,
			}

			err := d.Delete(wrongName)
			assert.Error(t, err, "title with such name doesn't exist in database")
			t.Run("Delete func doesn't remove element in map", func(t *testing.T) {
				assert.Len(t, d.db, 1)
			})
		})
	})
}
