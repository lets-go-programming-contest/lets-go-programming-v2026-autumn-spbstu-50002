package currency

import (
	"cmp"
	"slices"
)

func SortDesc(currencies []Currency) {
	slices.SortStableFunc(currencies, func(first, second Currency) int {
		return cmp.Compare(second.Value, first.Value)
	})
}
