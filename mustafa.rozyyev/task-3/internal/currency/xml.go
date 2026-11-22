package currency

import (
	"encoding/xml"
	"errors"
	"fmt"
	"os"

	"golang.org/x/net/html/charset"
)

type xmlDocument struct {
	Rates []XMLRate `xml:"Valute"`
}

func ReadXML(path string) ([]XMLRate, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open XML file %q: %w", path, err)
	}

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	var document xmlDocument
	decodeErr := decoder.Decode(&document)
	closeErr := file.Close()

	if decodeErr != nil {
		decodeErr = fmt.Errorf("decode XML file %q: %w", path, decodeErr)
	}

	if closeErr != nil {
		closeErr = fmt.Errorf("close XML file %q: %w", path, closeErr)
	}

	if err := errors.Join(decodeErr, closeErr); err != nil {
		return nil, err
	}

	return document.Rates, nil
}
