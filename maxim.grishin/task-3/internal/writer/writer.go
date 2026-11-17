package writer

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	dirPerm  os.FileMode = 0o755
	filePerm os.FileMode = 0o644
)

func CreateWriteAll(path string, data []byte) error {
	directory := filepath.Dir(path)

	err := os.MkdirAll(directory, dirPerm)
	if err != nil {
		return fmt.Errorf("Could not create directories: %s", err)
	}
	err = os.WriteFile(path, data, 0o644)
	if err != nil {
		return fmt.Errorf("Could not create file: %s", err)
	}

	return nil
}
