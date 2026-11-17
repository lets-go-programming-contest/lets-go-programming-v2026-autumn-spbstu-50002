package config

import (
	"errors"
	"fmt"
	"os"

	"go.yaml.in/yaml/v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse YAML config: %w", err)
	}

	if cfg.InputFile == "" {
		return nil, errors.New("input-file are required")
	}

	if cfg.OutputFile == "" {
		return nil, errors.New("output-file are required")
	}

	return &cfg, nil
}
