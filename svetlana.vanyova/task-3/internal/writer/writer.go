package writer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/KarKar07/task-3/internal/currency"
)

const (
	dirPerm  os.FileMode = 0o755
	filePerm os.FileMode = 0o644
)

func SaveJSON(path string, currencies []currency.Currency) error {
	if err := os.MkdirAll(filepath.Dir(path), dirPerm); err != nil {
		return fmt.Errorf("create directory: %w", err)
	}

	content, err := json.MarshalIndent(currencies, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal currencies: %w", err)
	}

	if err := os.WriteFile(path, content, filePerm); err != nil {
		return fmt.Errorf("save JSON file: %w", err)
	}

	return nil
}
