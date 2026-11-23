package writer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/kbnzhpp/task-3/src/currency"
)

const (
	dirPerm  os.FileMode = 0o755
	filePerm os.FileMode = 0o600
)

func SaveJSON(path string, currs []currency.Currency) error {
	if err := os.MkdirAll(filepath.Dir(path), dirPerm); err != nil {
		return fmt.Errorf("failed to create dir: %w", err)
	}

	res, err := json.MarshalIndent(currs, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to convert to JSON: %w", err)
	}

	if err := os.WriteFile(path, res, filePerm); err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}
