package convert

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Waffletsa1/task-3/internal/cbr"
)

func SaveJSON(path string, currencies []cbr.Currency) error {
	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}

	defer func() {
		_ = file.Close()
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(currencies); err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	return nil
}
