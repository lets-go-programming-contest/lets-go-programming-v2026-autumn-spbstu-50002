package currency

import (
	"encoding/xml"
	"fmt"
	"os"

	"golang.org/x/net/html/charset"
)

type CurrencyXML struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type CurrencyRatesXML struct {
	Currencies []CurrencyXML `xml:"Valute"`
}

func LoadCurrencies(path string) (CurrencyRatesXML, error) {
	file, err := os.Open(path)
	if err != nil {
		return CurrencyRatesXML{}, fmt.Errorf("failed to open XML file: %w", err)
	}

	defer func() {
		_ = file.Close()
	}()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	var rates CurrencyRatesXML

	if err := decoder.Decode(&rates); err != nil {
		return CurrencyRatesXML{}, fmt.Errorf("failed to decode XML: %w", err)
	}

	return rates, nil
}
