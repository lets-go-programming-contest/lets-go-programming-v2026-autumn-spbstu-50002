package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	ErrEmptyInputFile  = errors.New("input-file is empty")
	ErrEmptyOutputFile = errors.New("output-file is empty")
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Load(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}

	defer func() {
		_ = file.Close()
	}()

	var cfg Config

	if err := yaml.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, fmt.Errorf("failed to decode config: %w", err)
	}

	if cfg.InputFile == "" {
		return nil, ErrEmptyInputFile
	}

	if cfg.OutputFile == "" {
		return nil, ErrEmptyOutputFile
	}

	return &cfg, nil
}
