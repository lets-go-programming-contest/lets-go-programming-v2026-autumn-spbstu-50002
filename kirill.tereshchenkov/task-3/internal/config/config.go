package config

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"

	"go.yaml.in/yaml/v3"
)

var (
	errInputFileEmpty  = errors.New("input-file is empty")
	errOutputFileEmpty = errors.New("output-file is empty")
)

// Config содержит настройки приложения.
type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

// Load загружает конфигурацию из YAML-файла.
func Load(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read config: %w", err)
	}

	var cfg Config

	decoder := yaml.NewDecoder(bytes.NewReader(data))
	decoder.KnownFields(true)

	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, fmt.Errorf("failed to decode config: %w", err)
	}

	cfg.InputFile = strings.TrimSpace(cfg.InputFile)
	cfg.OutputFile = strings.TrimSpace(cfg.OutputFile)

	if cfg.InputFile == "" {
		return Config{}, errInputFileEmpty
	}

	if cfg.OutputFile == "" {
		return Config{}, errOutputFileEmpty
	}

	return cfg, nil
}
