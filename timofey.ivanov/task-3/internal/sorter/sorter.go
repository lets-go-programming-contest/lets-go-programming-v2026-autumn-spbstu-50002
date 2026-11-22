package sorter

import (
	"sort"

	"github.com/Kabisdox/task-3/internal/domain"
)

func ByValueDesc(items []domain.Currency) {
	sort.SliceStable(items, func(i, j int) bool {
		return items[i].Value > items[j].Value
	})
}
