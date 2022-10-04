package models

import (
	"time"

	"tgbot/internal/models/models_types"
)

type Title struct {
	ID         models_types.TitleID
	Name       models_types.TitleName
	URL        models_types.URL
	LastUpdate time.Time
}
