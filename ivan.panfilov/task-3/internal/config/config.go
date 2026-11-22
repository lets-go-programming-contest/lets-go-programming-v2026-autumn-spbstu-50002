package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
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
		return nil, fmt.Errorf("input-file is empty")
	}

	if cfg.OutputFile == "" {
		return nil, fmt.Errorf("output-file is empty")
	}

	return &cfg, nil
}
