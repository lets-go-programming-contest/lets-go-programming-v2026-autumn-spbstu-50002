package config

import (
	"errors"
	"fmt"
	"os"

	"go.yaml.in/yaml/v3"
)

var (
	ErrInputFileMissing  = errors.New("input-file is required")
	ErrOutputFileMissing = errors.New("output-file is required")
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Load(path string) (*Config, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(raw, &cfg); err != nil {
		return nil, fmt.Errorf("parse YAML: %w", err)
	}

	if cfg.InputFile == "" {
		return nil, ErrInputFileMissing
	}

	if cfg.OutputFile == "" {
		return nil, ErrOutputFileMissing
	}

	return &cfg, nil
}
