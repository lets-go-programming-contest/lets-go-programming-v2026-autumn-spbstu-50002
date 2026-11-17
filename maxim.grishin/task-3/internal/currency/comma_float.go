package currency

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type CommaFloat float32

func (commaFloat *CommaFloat) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var (
		raw          string
		floatBitSize = 32
	)

	if err := decoder.DecodeElement(&raw, &start); err != nil {
		return fmt.Errorf("decoding error: %w", err)
	}

	cleaned := strings.ReplaceAll(strings.TrimSpace(raw), ",", ".")

	val, err := strconv.ParseFloat(cleaned, floatBitSize)
	if err != nil {
		return fmt.Errorf("could not parse float: %w", err)
	}

	*commaFloat = CommaFloat(val)

	return nil
}
