package currency

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

var floatBitSize = 32

type CommaFloat float32

func (commaFloat *CommaFloat) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var raw string
	if err := decoder.DecodeElement(&raw, &start); err != nil {
		return fmt.Errorf("Decoding error: %s", err)
	}

	cleaned := strings.ReplaceAll(strings.TrimSpace(raw), ",", ".")

	val, err := strconv.ParseFloat(cleaned, floatBitSize)
	if err != nil {
		return fmt.Errorf("Could not parse float: %s", err)
	}

	*commaFloat = CommaFloat(val)

	return nil
}
