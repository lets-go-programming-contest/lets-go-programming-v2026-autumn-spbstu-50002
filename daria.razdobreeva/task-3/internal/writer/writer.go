package writer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func SaveJSON(path string, data any, dirPerm, filePerm os.FileMode) error {
	err := os.MkdirAll(filepath.Dir(path), dirPerm)
	if err != nil {
		return fmt.Errorf("create output dir: %w", err)
	}

	result, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return fmt.Errorf("encode json: %w", err)
	}

	err = os.WriteFile(path, result, filePerm)
	if err != nil {
		return fmt.Errorf("write output file: %w", err)
	}

	return nil
}
