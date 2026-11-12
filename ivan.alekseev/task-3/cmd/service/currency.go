package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

const (
	directoryPermissions = 0o755
	filePermissions      = 0o644
)

type currency struct {
	numCode  int
	charCode string
	value    float64
}

func loadCurrencies(path string) ([]currency, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open XML file: %w", err)
	}

	defer func() {
		_ = file.Close()
	}()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	var rates struct {
		Currencies []struct {
			NumCode  int    `xml:"NumCode"`
			CharCode string `xml:"CharCode"`
			Value    string `xml:"Value"`
		} `xml:"Valute"`
	}

	if err := decoder.Decode(&rates); err != nil {
		return nil, fmt.Errorf("decode XML file: %w", err)
	}

	currencies := make([]currency, 0, len(rates.Currencies))

	for _, item := range rates.Currencies {
		valueText := strings.ReplaceAll(
			strings.TrimSpace(item.Value),
			",",
			".",
		)

		value, err := strconv.ParseFloat(valueText, 64)
		if err != nil {
			return nil, fmt.Errorf("parse value for %s: %w", item.CharCode, err)
		}

		currencies = append(currencies, currency{
			numCode:  item.NumCode,
			charCode: item.CharCode,
			value:    value,
		})
	}

	return currencies, nil
}

func saveCurrencies(path string, currencies []currency) error {
	if err := os.MkdirAll(filepath.Dir(path), directoryPermissions); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}

	output := make([]map[string]any, 0, len(currencies))

	for _, item := range currencies {
		output = append(output, map[string]any{
			"num_code":  item.numCode,
			"char_code": item.charCode,
			"value":     item.value,
		})
	}

	data, err := json.MarshalIndent(output, "", "    ")
	if err != nil {
		return fmt.Errorf("encode JSON: %w", err)
	}

	if err := os.WriteFile(path, data, filePermissions); err != nil {
		return fmt.Errorf("write JSON file: %w", err)
	}

	return nil
}
