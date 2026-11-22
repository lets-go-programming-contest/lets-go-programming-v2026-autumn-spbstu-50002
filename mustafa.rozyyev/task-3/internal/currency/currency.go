package currency

import (
	"fmt"
	"strconv"
	"strings"
)

type XMLRate struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type Rate struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func Convert(source []XMLRate) ([]Rate, error) {
	rates := make([]Rate, 0, len(source))

	for _, item := range source {
		charCode := strings.TrimSpace(item.CharCode)
		valueText := strings.ReplaceAll(strings.TrimSpace(item.Value), ",", ".")

		value, err := strconv.ParseFloat(valueText, 64)
		if err != nil {
			return nil, fmt.Errorf("parse value for currency %s: %w", charCode, err)
		}

		rates = append(rates, Rate{
			NumCode:  item.NumCode,
			CharCode: charCode,
			Value:    value,
		})
	}

	return rates, nil
}
