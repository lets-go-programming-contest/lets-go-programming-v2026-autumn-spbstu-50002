package writer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Ministrius/task-3/internal/currency"
)

const (
	dirP  os.FileMode = 0o755
	fileP os.FileMode = 0o644
)

func WriteJSON(path string, currencies []currency.Currency) error {
	dir := filepath.Dir(path)

	err := os.MkdirAll(dir, dirP)
	if err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	data, err := json.MarshalIndent(currencies, "", "\t")
	if err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	data = append(data, '\n')

	err = os.WriteFile(path, data, fileP)
	if err != nil {
		return fmt.Errorf("failed to write JSON file: %w", err)
	}

	return nil
}
