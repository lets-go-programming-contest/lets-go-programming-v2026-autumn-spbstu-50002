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
	errMissingInputFile  = errors.New("input-file is missing")
	errMissingOutputFile = errors.New("output-file is missing")
)

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse YAML configuration file: %w", err)
	}

	if cfg.InputFile == "" {
		return nil, errMissingInputFile
	}

	if cfg.OutputFile == "" {
		return nil, errMissingOutputFile
	}

	return &cfg, nil
}
