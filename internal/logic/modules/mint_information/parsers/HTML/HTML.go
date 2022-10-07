package HTML

import (
	"bytes"
	"errors"
	"time"

	"github.com/PuerkitoBio/goquery"

	parser_models "tgbot/internal/logic/modules/mint_information/models"
	"tgbot/internal/models/models_types"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (h *Parser) Parse(b []byte) (parser_models.TitleParams, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(b))
	if err != nil {
		return parser_models.TitleParams{}, err
	}

	name, err := getName(doc)
	if err != nil {
		return parser_models.TitleParams{}, err
	}

	lastUPD, err := getLastUPD(doc)
	if err != nil {
		return parser_models.TitleParams{}, err
	}

	titleParams := parser_models.TitleParams{
		Name:    name,
		LastUPD: lastUPD,
	}

	return titleParams, nil
}

func getName(doc *goquery.Document) (models_types.TitleName, error) {
	titleName := doc.Find(mintName)
	name, err := titleName.Html()
	if err != nil {
		return "", err
	}
	if name == "" {
		return "", errors.New("can not parse this page")
	}
	return models_types.TitleName(name), err
}

func getLastUPD(doc *goquery.Document) (time.Time, error) {
	MintLastUPD := doc.Find(mintChapter).First().Find(chapterDate)
	StrLastUPD, exist := MintLastUPD.Attr(date)
	if !exist {
		return time.Unix(0, 0), nil
	}

	lastUPD, err := time.Parse("2006-01-02 15:04:05", StrLastUPD)
	if err != nil {
		return time.Time{}, err
	}

	return lastUPD, nil
}
