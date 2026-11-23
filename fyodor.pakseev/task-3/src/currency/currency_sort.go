package currency

import (
	"cmp"
	"slices"
)

func SortCurrenciesByValue(currencies []Currency) {
	slices.SortFunc(currencies, func(a, b Currency) int {
		return cmp.Compare(b.Value, a.Value)
	})
}
