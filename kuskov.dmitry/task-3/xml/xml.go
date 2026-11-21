package xml

import (
	encodingxml "encoding/xml"
	"errors"
	"os"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

var errMissingFields = errors.New("missing currency fields")

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

type valute struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type rates struct {
	Valutes []valute `xml:"Valute"`
}

func Load(path string) ([]Currency, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := encodingxml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	var source rates
	if err := decoder.Decode(&source); err != nil {
		return nil, err
	}

	currencies := make([]Currency, 0, len(source.Valutes))
	for _, valute := range source.Valutes {
		if valute.NumCode == 0 || valute.CharCode == "" || valute.Value == "" {
			return nil, errMissingFields
		}
		value, err := strconv.ParseFloat(strings.ReplaceAll(valute.Value, ",", "."), 64)
		if err != nil {
			return nil, err
		}
		currencies = append(currencies, Currency{
			NumCode:  valute.NumCode,
			CharCode: valute.CharCode,
			Value:    value,
		})
	}
	return currencies, nil
}
