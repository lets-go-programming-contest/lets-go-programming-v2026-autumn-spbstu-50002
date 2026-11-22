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

func ToCurrencies(valutes []ValuteXML) ([]Currency, error) {
	result := make([]Currency, 0, len(valutes))

	for _, valute := range valutes {
		text := strings.TrimSpace(valute.Value)

		value, err := strconv.ParseFloat(strings.Replace(text, ",", ".", 1), 64)
		if err != nil {
			return nil, fmt.Errorf("cannot parse value of %s: %w", valute.CharCode, err)
		}

		result = append(result, Currency{
			NumCode:  valute.NumCode,
			CharCode: valute.CharCode,
			Value:    value,
		})
	}

	return result, nil
}
