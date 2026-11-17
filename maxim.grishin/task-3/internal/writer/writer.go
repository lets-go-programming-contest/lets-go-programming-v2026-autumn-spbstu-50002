package writer

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateWriteAll(path string, data []byte) error {
	var (
		dirPerm  os.FileMode = 0o750
		filePerm os.FileMode = 0o600
	)

	directory := filepath.Dir(path)

	if err := os.MkdirAll(directory, dirPerm); err != nil {
		return fmt.Errorf("could not create directories: %w", err)
	}

	if err := os.WriteFile(path, data, filePerm); err != nil {
		return fmt.Errorf("could not create file: %w", err)
	}

	return nil
}
