package config

import (
	"errors"
	"fmt"
	"os"

	"go.yaml.in/yaml/v3"
)

var (
	errEmptyInput  = errors.New("config: input-file must not be empty")
	errEmptyOutput = errors.New("config: output-file must not be empty")
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Load(path string) (*Config, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	var cfg Config

	if err := yaml.Unmarshal(raw, &cfg); err != nil {
		return nil, fmt.Errorf("decode yaml: %w", err)
	}

	if cfg.InputFile == "" {
		return nil, errEmptyInput
	}

	if cfg.OutputFile == "" {
		return nil, errEmptyOutput
	}

	return &cfg, nil
}
