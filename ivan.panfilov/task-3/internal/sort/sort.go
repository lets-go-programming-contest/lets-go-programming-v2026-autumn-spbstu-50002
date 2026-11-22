package sort

import (
	"sort"

	"github.com/Waffletsa1/task-3/internal/cbr"
)

func ByValueDesc(currencies []cbr.Currency) {
	sort.SliceStable(currencies, func(i, j int) bool {
		if currencies[i].Value == currencies[j].Value {
			return currencies[i].NumCode < currencies[j].NumCode
		}
		return currencies[i].Value > currencies[j].Value
	})
}
