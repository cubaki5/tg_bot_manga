package models

import (
	"time"

	"tgbot/internal/models/models_types"
)

type TitleParams struct {
	Name    models_types.TitleName
	LastUPD time.Time
}
