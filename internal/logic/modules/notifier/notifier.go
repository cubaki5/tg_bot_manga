package notifier

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/labstack/gommon/log"

	parser_models "tgbot/internal/logic/modules/mint_information/models"
	"tgbot/internal/models"
	"tgbot/internal/models/models_types"
)

//go:generate mockgen --source=notifier.go --destination=mocks/mock_notifier.go --package=mocks

type (
	WebClient interface {
		DoGetRequest(url models_types.URL) ([]byte, error)
	}

	TGClient interface {
		PostMsg(title models.Title, user models.User) error
	}

	Database interface {
		List() map[models_types.TitleID]models.Title
	}

	Parser interface {
		Parse(b []byte) (parser_models.TitleParams, error)
	}
)

type Notifier struct {
	client   WebClient
	db       Database
	tgClient TGClient
	parser   Parser
}

func NewNotifier(cl WebClient, db Database, tgClient TGClient, parser Parser) *Notifier {
	return &Notifier{
		client:   cl,
		db:       db,
		tgClient: tgClient,
		parser:   parser,
	}
}

func (n *Notifier) CheckUpdates() {
	//каждый час я проверяю есть обновления во всём списке и если встречаю тру, то вызываю метод пост у телеграмклиента
	go func() {
		tick := time.NewTicker(time.Hour)
		defer tick.Stop()

		for range tick.C {
			n.checkListFotUpdates()
		}
	}()
}

func (n *Notifier) checkListFotUpdates() {
	titles := n.db.List()

	chatID, err := strconv.Atoi(os.Getenv("CHAT_ID"))
	if err != nil {
		log.Error(err)
		return
	}
	user := models.User{ID: models_types.ChatID(chatID)}

	for _, title := range titles {

		go func(title models.Title) {
			isUPD, err := n.IsUpdated(title)
			if err != nil {
				log.Error(err)
				return
			}

			if isUPD {
				err = n.tgClient.PostMsg(title, user)
				if err != nil {
					log.Error(err)
					return
				}
			}
		}(title)

	}
}

func (n *Notifier) IsUpdated(title models.Title) (bool, error) {
	b, err := n.client.DoGetRequest(title.URL)
	if err != nil {
		return false, err
	}

	updatedTitle, err := n.parser.Parse(b)
	if err != nil {
		return false, err
	}

	if title.Name != updatedTitle.Name {
		return false, errors.New("names mismatching")
	}

	if updatedTitle.LastUPD.After(title.LastUpdate) {
		title.LastUpdate = updatedTitle.LastUPD
		return true, nil
	}

	return false, nil
}
