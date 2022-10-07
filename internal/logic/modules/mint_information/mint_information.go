package mint_information

import (
	parser_models "tgbot/internal/logic/modules/mint_information/models"
	"tgbot/internal/models"
	"tgbot/internal/models/models_types"
)

//go:generate mockgen --source=mint_information.go --destination=mocks/mock_mint_information.go --package=mocks

type (
	Parser interface {
		Parse(b []byte) (parser_models.TitleParams, error)
	}
	WebClient interface {
		DoGetRequest(url models_types.URL) ([]byte, error)
	}
)

type GetTitleModule struct {
	client WebClient
	parser Parser
}

func NewGetTitleModule(client WebClient, parser Parser) *GetTitleModule {
	return &GetTitleModule{
		client: client,
		parser: parser,
	}
}

func (m *GetTitleModule) GetTitle(URL models_types.URL) (models.Title, error) {
	b, err := m.client.DoGetRequest(URL)
	if err != nil {
		return models.Title{}, err
	}

	titleParams, err := m.parser.Parse(b)
	if err != nil {
		return models.Title{}, err
	}

	title := models.Title{
		Name:       titleParams.Name,
		URL:        URL,
		LastUpdate: titleParams.LastUPD,
	}

	return title, nil
}
