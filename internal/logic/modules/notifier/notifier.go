package notifier

import (
	"tgbot/internal/models"
	"tgbot/internal/models/models_types"
)

type (
	WebClient interface {
		IsUpdated() bool
	}

	TGClient interface {
		PostMsg() error
	}

	Database interface {
		List() map[models_types.TitleID]models.Title
	}
)

type Notifier struct {
	client   WebClient
	db       Database
	tgClient TGClient
}

func NewNotifier(cl WebClient, db Database, tgClient TGClient) *Notifier {
	return &Notifier{
		client:   cl,
		db:       db,
		tgClient: tgClient,
	}
}

func (n *Notifier) CheckUpdates() {
	//каждый час я проверяю есть обновления во всём списке и если встречаю тру, то вызываю метод пост у телеграмклиента
}
