package runtime_database

import (
	"errors"
	"sync"

	"tgbot/internal/models"
	"tgbot/internal/models/models_types"
)

type MangaDatabase struct {
	mx   sync.RWMutex
	db   map[models_types.TitleID]models.Title
	incr models_types.TitleID
}

var mangaDB *MangaDatabase

var once sync.Once

func NewDatabase() *MangaDatabase {
	once.Do(func() {
		mangaDB = &MangaDatabase{
			db:   make(map[models_types.TitleID]models.Title, 500),
			incr: 1,
		}
	})
	return mangaDB
}

func (d *MangaDatabase) List() map[models_types.TitleID]models.Title {
	d.mx.RLock()
	var titles = make(map[models_types.TitleID]models.Title, len(d.db))
	for key, title := range d.db {
		titles[key] = title
	}
	d.mx.RUnlock()
	return titles
}

func (d *MangaDatabase) Add(title models.Title) {
	d.mx.Lock()
	title.ID = d.incr
	d.db[d.incr] = title
	d.incr++
	d.mx.Unlock()
}

func (d *MangaDatabase) Delete(titleName models_types.TitleName) error {
	d.mx.Lock()
	defer d.mx.Unlock()
	for key := range d.db {
		if d.db[key].Name == titleName {
			delete(d.db, key)
			return nil
		}
	}
	return errors.New("title with such name doesn't exist in database")
}
