package currency

import "sort"

func SortByValueDesc(list []Currency) {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].Value > list[j].Value
	})
}
