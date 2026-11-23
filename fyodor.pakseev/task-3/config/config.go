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

var (
	errNoInputFileProvided  = errors.New("input-file is not provided in config")
	errNoOutputFileProvided = errors.New("output-file is not provided in config")
)

func LoadConfig(configPath string) (*Config, error) {
	var cfg Config

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	if cfg.InputFile == "" {
		return nil, errNoInputFileProvided
	}

	if cfg.OutputFile == "" {
		return nil, errNoOutputFileProvided
	}

	return &cfg, nil
}
