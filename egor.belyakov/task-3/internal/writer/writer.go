package writer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Dadnsall/task-3/internal/currency"
)

const (
	dirMode  = 0o755
	fileMode = 0o644
)

func WriteJSON(path string, currencies []currency.Currency) error {
	data, err := json.MarshalIndent(currencies, "", "  ")
	if err != nil {
		return fmt.Errorf("cannot encode json: %w", err)
	}

	err = os.MkdirAll(filepath.Dir(path), dirMode)
	if err != nil {
		return fmt.Errorf("cannot create output dir: %w", err)
	}

	err = os.WriteFile(path, data, fileMode)
	if err != nil {
		return fmt.Errorf("cannot write json: %w", err)
	}

	return nil
}
