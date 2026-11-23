package parser

import (
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kbnzhpp/task-3/src/currency"
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

func ParseXML(path string) ([]currency.Currency, error) {
	var rawXML CurrencyRatesXML

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read XML file: %w", err)
	}

	defer func() {
		_ = file.Close()
	}()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	if err := decoder.Decode(&rawXML); err != nil {
		return nil, fmt.Errorf("failed to decode XML: %w", err)
	}

	rates := make([]currency.Currency, 0, len(rawXML.Currencies))

	for _, item := range rawXML.Currencies {
		parsed, err := ConvertXMLValueToFloat(item.Value)
		if err != nil {
			return nil, fmt.Errorf("failed to convert string to float: %w", err)
		}

		rates = append(rates, currency.Currency{
			NumCode:  item.NumCode,
			CharCode: item.CharCode,
			Value:    parsed,
		})
	}

	return rates, nil
}

func ConvertXMLValueToFloat(value string) (float64, error) {
	raw := strings.ReplaceAll(strings.TrimSpace(value), ",", ".")

	res, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse value: %w", err)
	}

	return res, nil
}
