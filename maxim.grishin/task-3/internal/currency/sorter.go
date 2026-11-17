package currency

import (
	"cmp"
	"slices"
)

func Sort(valCurs ValCurs) {
	slices.SortFunc(valCurs.Valutes, func(a, b Valute) int {
		return cmp.Compare(b.Value, a.Value)
	})
}
