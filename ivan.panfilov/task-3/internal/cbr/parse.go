package cbr

import (
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

func ParseXML(path string) ([]Currency, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open XML file: %w", err)
	}

	defer func() {
		_ = file.Close()
	}()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	var document xmlDocument

	if err := decoder.Decode(&document); err != nil {
		return nil, fmt.Errorf("failed to decode XML file: %w", err)
	}

	currencies := make([]Currency, 0, len(document.Valutes))

	for _, valute := range document.Valutes {
		charCode := strings.TrimSpace(valute.CharCode)

		rawValue := strings.ReplaceAll(
			strings.TrimSpace(valute.Value),
			",",
			".",
		)

		value, err := strconv.ParseFloat(rawValue, 64)
		if err != nil {
			return nil, fmt.Errorf(
				"failed to parse value for %s: %w",
				charCode,
				err,
			)
		}

		parsedCurrency := Currency{
			NumCode:  valute.NumCode,
			CharCode: charCode,
			Value:    value,
		}
		currencies = append(currencies, parsedCurrency)
	}

	return currencies, nil
}
