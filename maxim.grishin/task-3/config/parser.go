package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func Parse(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("YAML Parser could not read file: %s", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("YAML Unmarshaling error: %s", err)
	}

	return &config, nil
}
