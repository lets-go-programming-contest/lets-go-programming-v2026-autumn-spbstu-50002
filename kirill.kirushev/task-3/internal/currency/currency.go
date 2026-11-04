package currency

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var errMissingCurrencyFields = errors.New("currency has missing required fields")

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func ConvertCurrency(currency CurrencyXML) (Currency, error) {
	if currency.NumCode == 0 || currency.CharCode == "" || currency.Value == "" {
		return Currency{}, errMissingCurrencyFields
	}

	value := strings.ReplaceAll(currency.Value, ",", ".")

	parsedValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return Currency{}, fmt.Errorf("failed to parse currency value: %w", err)
	}

	return Currency{
		NumCode:  currency.NumCode,
		CharCode: currency.CharCode,
		Value:    parsedValue,
	}, nil
}

func ConvertCurrencies(source []CurrencyXML) ([]Currency, error) {
	currencies := make([]Currency, 0, len(source))

	for _, currencyXML := range source {
		currency, err := ConvertCurrency(currencyXML)
		if err != nil {
			return nil, fmt.Errorf("failed to convert currency: %w", err)
		}

		currencies = append(currencies, currency)
	}

	return currencies, nil
}
