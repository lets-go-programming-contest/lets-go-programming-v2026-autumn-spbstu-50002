package currency

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

type valCurs struct {
	Valutes []ValuteXML `xml:"Valute"`
}

func ReadXML(path string) ([]ValuteXML, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("cannot open xml: %w", err)
	}

	defer func() {
		_ = file.Close()
	}()

	var curs valCurs

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	err = decoder.Decode(&curs)
	if err != nil {
		return nil, fmt.Errorf("cannot decode xml: %w", err)
	}

	return curs.Valutes, nil
}
