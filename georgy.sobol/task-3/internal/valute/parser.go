package valute

import (
	"encoding/xml"
	"fmt"
	"os"

	"golang.org/x/net/html/charset"
)

type ValuteXML struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type ValuteCostsXML struct {
	Valutes []ValuteXML `xml:"Valute"`
}

func LoadXML(path string) (*ValuteCostsXML, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open XML file: %w", err)
	}

	defer func() {
		_ = file.Close()
	}()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	var costs ValuteCostsXML

	if err := decoder.Decode(&costs); err != nil {
		return nil, fmt.Errorf("failed to decode XML file: %w", err)
	}

	return &costs, nil
}
