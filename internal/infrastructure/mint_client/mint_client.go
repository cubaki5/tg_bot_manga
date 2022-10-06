package mint_client

import (
	"io"
	"net/http"

	"github.com/labstack/gommon/log"

	"tgbot/internal/models/models_types"
)

type MintClient struct {
}

func NewMintClient() *MintClient {
	return &MintClient{}
}

func (m *MintClient) IsUpdated() bool {
	//TODO implement me
	panic("implement me")
}

func (m *MintClient) DoGetRequest(url models_types.URL) ([]byte, error) {

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "TG_Bot")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = res.Body.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	if res.StatusCode != 200 {
		log.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	var b []byte

	b, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}
