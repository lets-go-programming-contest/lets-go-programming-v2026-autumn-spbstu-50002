package parser

import (
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Kabisdox/task-3/internal/domain"
	"golang.org/x/net/html/charset"
)

type XMLParser struct{}

func (XMLParser) Parse(path string) ([]domain.Currency, error) {
	xmlFile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open xml: %w", err)
	}

	defer func() {
		_ = xmlFile.Close()
	}()

	decoder := xml.NewDecoder(xmlFile)
	decoder.CharsetReader = charset.NewReaderLabel

	var doc domain.ValCursXML

	if err := decoder.Decode(&doc); err != nil {
		return nil, fmt.Errorf("decode xml: %w", err)
	}

	result := make([]domain.Currency, 0, len(doc.Valutes))

	for _, valute := range doc.Valutes {
		normalized := strings.ReplaceAll(valute.Value, ",", ".")

		parsed, err := strconv.ParseFloat(normalized, 64)
		if err != nil {
			return nil, fmt.Errorf("parse value %q: %w", valute.Value, err)
		}

		result = append(result, domain.Currency{
			NumCode:  valute.NumCode,
			CharCode: valute.CharCode,
			Value:    parsed,
		})
	}

	return result, nil
}
