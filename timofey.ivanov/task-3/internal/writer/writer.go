package writer

import "github.com/Kabisdox/task-3/internal/domain"

type Saver interface {
	Save(path string, items []domain.Currency) error
}
