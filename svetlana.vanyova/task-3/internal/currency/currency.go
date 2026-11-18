package currency

import (
	"fmt"
	"strconv"
	"strings"
)

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func ConvertCurrencies(source []CurrencyXML) ([]Currency, error) {
	currencies := make([]Currency, 0, len(source))

	for _, item := range source {
		value := strings.ReplaceAll(item.Value, ",", ".")

		parsedValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse value: %w", err)
		}

		currencies = append(currencies, Currency{
			NumCode:  item.NumCode,
			CharCode: item.CharCode,
			Value:    parsedValue,
		})
	}

	return currencies, nil
}
