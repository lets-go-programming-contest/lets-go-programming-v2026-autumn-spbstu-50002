package parser

import "github.com/Kabisdox/task-3/internal/domain"

type Parser interface {
	Parse(path string) ([]domain.Currency, error)
}
