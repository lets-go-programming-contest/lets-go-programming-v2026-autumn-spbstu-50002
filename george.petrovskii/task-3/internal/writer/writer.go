package writer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"george.petrovskii/task-3/internal/currency"
)

const (
	dirPermissions  os.FileMode = 0o755
	filePermissions os.FileMode = 0o644
)

func WriteJSON(path string, currencies []currency.Currency) error {
	data, err := json.MarshalIndent(currencies, "", "    ")
	if err != nil {
		return fmt.Errorf("encode json: %w", err)
	}

	data = append(data, '\n')

	if err := os.MkdirAll(filepath.Dir(path), dirPermissions); err != nil {
		return fmt.Errorf("create output dir: %w", err)
	}

	if err := os.WriteFile(path, data, filePermissions); err != nil {
		return fmt.Errorf("write output file: %w", err)
	}

	return nil
}
