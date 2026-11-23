package writer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/MountenY/task-3/internal/currency"
)

const (
	dirPerm  os.FileMode = 0o755
	filePerm os.FileMode = 0o644
)

func SaveJSON(path string, valutes []currency.Currency) error {
	if err := os.MkdirAll(filepath.Dir(path), dirPerm); err != nil {
		return fmt.Errorf("mkdir: %w", err)
	}

	data, err := json.MarshalIndent(valutes, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}

	if err := os.WriteFile(path, data, filePerm); err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	return nil
}
