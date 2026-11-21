package valute

import (
	"fmt"
	"strconv"
	"strings"
)

type Valute struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func ConvertValutes(source []ValuteXML) ([]Valute, error) {
	valutes := make([]Valute, 0, len(source))

	for _, current := range source {
		valueStr := strings.ReplaceAll(current.Value, ",", ".")

		valueFloat, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse value: %w", err)
		}

		valutes = append(valutes, Valute{
			NumCode:  current.NumCode,
			CharCode: current.CharCode,
			Value:    valueFloat,
		})
	}

	return valutes, nil
}
