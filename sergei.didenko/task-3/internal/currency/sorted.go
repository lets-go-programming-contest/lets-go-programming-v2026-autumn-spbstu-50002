package currency

import "sort"

func SortByValueDescending(currencies []Currency) {
	sort.SliceStable(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})
}
