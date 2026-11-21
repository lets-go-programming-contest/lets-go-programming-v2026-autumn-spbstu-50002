package valute

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	defaultDirPerms  = 0o755
	defaultFilePerms = 0o644
)

func ValutesToJSON(path string, valutes []Valute) error {
	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, defaultDirPerms); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	data, err := json.MarshalIndent(valutes, "", "		")
	if err != nil {
		return fmt.Errorf("failed to encode JSON file: %w", err)
	}

	if err := os.WriteFile(path, data, defaultFilePerms); err != nil {
		return fmt.Errorf("failed to write JSON file: %w", err)
	}

	return nil
}
