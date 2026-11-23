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

func Convert(source []CurrencyXML) ([]Currency, error) {
	result := make([]Currency, 0, len(source))

	for _, item := range source {
		rawValue := strings.ReplaceAll(item.Value, ",", ".")

		parsedVal, err := strconv.ParseFloat(rawValue, 64)
		if err != nil {
			return nil, fmt.Errorf("parse %s: %w", item.CharCode, err)
		}

		resItem := Currency{
			NumCode:  item.NumCode,
			CharCode: item.CharCode,
			Value:    parsedVal,
		}
		result = append(result, resItem)
	}

	return result, nil
}
