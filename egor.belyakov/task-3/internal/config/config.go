package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	errNoInput  = errors.New("config: input-file is empty")
	errNoOutput = errors.New("config: output-file is empty")
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Read(path string) (Config, error) {
	var cfg Config

	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, fmt.Errorf("cannot read config: %w", err)
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return cfg, fmt.Errorf("cannot parse config: %w", err)
	}

	switch {
	case cfg.InputFile == "":
		return cfg, errNoInput
	case cfg.OutputFile == "":
		return cfg, errNoOutput
	}

	return cfg, nil
}
