package currency

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	defaultDirPerms  = 0o755
	defaultFilePerms = 0o600
)

func WriteJSON(path string, rates []Rate) error {
	if err := os.MkdirAll(filepath.Dir(path), defaultDirPerms); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}

	data, err := json.MarshalIndent(rates, "", "  ")
	if err != nil {
		return fmt.Errorf("encode JSON: %w", err)
	}

	if err := os.WriteFile(path, data, defaultFilePerms); err != nil {
		return fmt.Errorf("write JSON file %q: %w", path, err)
	}

	return nil
}
