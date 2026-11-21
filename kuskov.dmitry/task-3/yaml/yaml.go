package yaml

import (
	"errors"
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
		return Config{}, err
	}

	var cfg Config
	if err := yamlv3.Unmarshal(data, &cfg); err != nil {
		return Config{}, err
	}
	if cfg.Input == "" || cfg.Output == "" {
		return Config{}, errMissingFields
	}
	return cfg, nil
}
