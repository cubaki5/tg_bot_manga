package HTML

import (
	"errors"
	"io"
	"log"
	"os"
	"testing"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"tgbot/internal/models/models_types"
)

const (
	testFilesPathLastUPD = "tests_files/get_last_upd/"
	testFilesPathName    = "tests_files/get_name/"
	testFilePathParser   = "tests_files/parser/parser_happy_path.html"
)

func parser(path string) *goquery.Document {
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
	}

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		log.Println(err)
	}

	return doc
}

func TestParser_Parse(t *testing.T) {
	t.Run("Happy path", func(t *testing.T) {
		path := testFilePathParser
		file, err := os.Open(path)
		if err != nil {
			log.Println(err)
		}

		var b []byte
		b, err = io.ReadAll(file)
		if err != nil {
			log.Println(err)
		}

		expTitleName := models_types.TitleName("TestName")
		expLastUPD := time.Date(
			2022,         // год
			time.October, // месяц
			7,            // день
			12,           // часы
			41,           // минуты
			17,           // секунды
			0,            // наносекунды
			time.UTC,     // временная зона
		)

		p := NewParser()

		actTitleParams, err := p.Parse(b)

		assert.NoError(t, err)

		t.Run("Parser return correct name", func(t *testing.T) {
			require.Equal(t, expTitleName, actTitleParams.Name)
		})
		t.Run("Parser return correct last upd", func(t *testing.T) {
			require.Equal(t, expLastUPD, actTitleParams.LastUPD)
		})
	})
}

func Test_GetName(t *testing.T) {
	tests := []struct {
		name         string
		expTitleName models_types.TitleName
		expError     error
		fileName     string
	}{
		{
			name:         "Happy Path",
			expTitleName: models_types.TitleName("TestName"),
			expError:     nil,
			fileName:     "get_name_happy_path.html",
		},
		{
			name:         "HTML has empty name",
			expTitleName: "",
			expError:     errors.New("can not parse this page"),
			fileName:     "get_name_has_empty_name.html",
		},
		{
			name:         "HTML does not have class .name",
			expTitleName: "",
			expError:     errors.New("can not parse this page"),
			fileName:     "get_name_does_not_have_class_name.html",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			path := testFilesPathName + test.fileName

			doc := parser(path)

			actTitleName, err := getName(doc)
			assert.Equal(t, test.expTitleName, actTitleName)
			if test.expError == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, test.expError.Error())
			}
		})
	}
}

func Test_LastUPD(t *testing.T) {
	tests := []struct {
		name       string
		expLastUPD time.Time
		expError   error
		fileName   string
	}{
		{
			name: "Happy Path",
			expLastUPD: time.Date(
				2022,         // год
				time.October, // месяц
				7,            // день
				12,           // часы
				41,           // минуты
				17,           // секунды
				0,            // наносекунды
				time.UTC,     // временная зона
			),
			expError: nil,
			fileName: "get_last_upd_happy_path.html",
		},
		{
			name:       "HTML has empty date",
			expLastUPD: time.Time{},
			expError:   errors.New("parsing time \"\" as \"2006-01-02 15:04:05\": cannot parse \"\" as \"2006\""),
			fileName:   "get_last_upd_has_empty_time.html",
		},
		{
			name:       "HTML has date with wrong format",
			expLastUPD: time.Time{},
			expError:   errors.New("parsing time \"12:41:17.087 2022-10-07\" as \"2006-01-02 15:04:05\": cannot parse \"1:17.087 2022-10-07\" as \"2006\""),
			fileName:   "get_last_upd_has_wrong_time_format.html",
		},
		{
			name:       "Gets html with no such attribute",
			expLastUPD: time.Unix(0, 0),
			expError:   nil,
			fileName:   "get_last_upd_no_attr_date.html",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			path := testFilesPathLastUPD + test.fileName

			doc := parser(path)

			actLastUPD, err := getLastUPD(doc)
			assert.Equal(t, test.expLastUPD, actLastUPD)
			if test.expError == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, test.expError.Error())
			}
		})
	}
}
