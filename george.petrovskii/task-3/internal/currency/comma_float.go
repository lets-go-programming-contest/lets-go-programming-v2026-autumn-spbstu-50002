package currency

import (
	"fmt"
	"strconv"
	"strings"
)

type CommaFloat float64

func (c *CommaFloat) UnmarshalText(text []byte) error {
	s := strings.ReplaceAll(string(text), ",", ".")

	f, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
	if err != nil {
		return fmt.Errorf("parse value %q: %w", string(text), err)
	}

	*c = CommaFloat(f)

	return nil
}
