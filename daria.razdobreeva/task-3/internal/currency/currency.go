package currency

import (
	"bytes"
	"cmp"
	"encoding/xml"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

type xmlValute struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type xmlValCurs struct {
	XMLName xml.Name    `xml:"ValCurs"`
	Valutes []xmlValute `xml:"Valute"`
}

func Load(path string) ([]Currency, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read input file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	var curs xmlValCurs

	err = decoder.Decode(&curs)
	if err != nil {
		return nil, fmt.Errorf("decode xml: %w", err)
	}

	return convert(curs.Valutes)
}

func convert(valutes []xmlValute) ([]Currency, error) {
	currencies := make([]Currency, 0, len(valutes))

	for _, valute := range valutes {
		text := strings.ReplaceAll(strings.TrimSpace(valute.Value), ",", ".")

		value, err := strconv.ParseFloat(text, 64)
		if err != nil {
			return nil, fmt.Errorf("parse value of %s: %w", valute.CharCode, err)
		}

		currencies = append(currencies, Currency{
			NumCode:  valute.NumCode,
			CharCode: valute.CharCode,
			Value:    value,
		})
	}

	return currencies, nil
}

func SortByValueDesc(currencies []Currency) {
	slices.SortFunc(currencies, func(a, b Currency) int {
		return cmp.Compare(b.Value, a.Value)
	})
}
