package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func Parse(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("yAML Parser could not read file: %w", err)
	}

	var config Config

	if err = yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("yAML Unmarshaling error: %w", err)
	}

	return &config, nil
}
