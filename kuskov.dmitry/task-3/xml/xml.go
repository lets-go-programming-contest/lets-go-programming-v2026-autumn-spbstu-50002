package xml

import (
	encodingxml "encoding/xml"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

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
		return nil, fmt.Errorf("open %q: %w", path, err)
	}

	defer func() {
		_ = file.Close()
	}()

	decoder := encodingxml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	var source rates

	if err := decoder.Decode(&source); err != nil {
		return nil, fmt.Errorf("decode %q: %w", path, err)
	}

	currencies := make([]Currency, 0, len(source.Valutes))

	for _, valute := range source.Valutes {
		value, err := strconv.ParseFloat(strings.ReplaceAll(valute.Value, ",", "."), 64)
		if err != nil {
			return nil, fmt.Errorf("parse value %q: %w", valute.Value, err)
		}

		currencies = append(currencies, Currency{
			NumCode:  valute.NumCode,
			CharCode: valute.CharCode,
			Value:    value,
		})
	}

	return currencies, nil
}
