package models

import (
	"tgbot/internal/models/models_types"
	"time"
)

type Title struct {
	TitleName  models_types.TitleName
	URL        models_types.URL
	LastUpdate time.Time
}
