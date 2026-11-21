package yaml

import (
	"errors"
	"fmt"
	"os"

	yamlv3 "go.yaml.in/yaml/v3"
)

var errMissingFields = errors.New("input-file and output-file are required")

type Config struct {
	Input  string `yaml:"input-file"`
	Output string `yaml:"output-file"`
}

func Load(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("read %q: %w", path, err)
	}

	var cfg Config

	if err := yamlv3.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("unmarshal %q: %w", path, err)
	}

	if cfg.Input == "" || cfg.Output == "" {
		return Config{}, errMissingFields
	}

	return cfg, nil
}
