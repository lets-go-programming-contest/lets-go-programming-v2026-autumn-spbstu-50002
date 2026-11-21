package valute

import "sort"

func SortInDecreasingOrder(valutes []Valute) {
	sort.SliceStable(valutes, func(i, j int) bool {
		return valutes[i].Value > valutes[j].Value
	})
}
