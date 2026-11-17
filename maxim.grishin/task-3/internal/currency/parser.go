package currency

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

var errUnsupportedEncoding = errors.New("unsupported encoding")

type Parser struct {
	inputFilePath string
}

func NewParser(path string) *Parser {
	parser := &Parser{
		inputFilePath: path,
	}

	return parser
}

func (parser *Parser) Parse() (*ValCurs, error) {
	file, err := os.Open(parser.inputFilePath)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	defer func() {
		_ = file.Close()
	}() //

	decoder := xml.NewDecoder(file)

	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch strings.ToLower(charset) {
		case "windows-1251", "cp1251":
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		case "koi8-r":
			return charmap.KOI8R.NewDecoder().Reader(input), nil
		default:
			return nil, fmt.Errorf("%w: %s", errUnsupportedEncoding, charset)
		}
	}

	var result ValCurs

	if err = decoder.Decode(&result); err != nil {
		return nil, fmt.Errorf("decoding error: %w", err)
	}

	return &result, nil
}
