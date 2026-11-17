package writer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	dirPerm  = 0755
	filePerm = 0644
)

func SaveJSON(path string, currencies []Currency) error {
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
