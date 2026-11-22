package writer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Kabisdox/task-3/internal/domain"
)

const (
	defaultDirMode  os.FileMode = 0o755
	defaultFileMode os.FileMode = 0o644
)

// JSONWriter реализует Saver для JSON-файлов.
type JSONWriter struct{}

// Save сохраняет валюты в JSON, создавая директории при необходимости.
func (JSONWriter) Save(path string, items []domain.Currency) error {
	if err := os.MkdirAll(filepath.Dir(path), defaultDirMode); err != nil {
		return fmt.Errorf("ensure output dir: %w", err)
	}

	data, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return fmt.Errorf("encode json: %w", err)
	}

	if err := os.WriteFile(path, data, defaultFileMode); err != nil {
		return fmt.Errorf("write result file: %w", err)
	}

	return nil
}
