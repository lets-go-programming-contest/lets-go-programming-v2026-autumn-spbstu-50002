package main

import (
	"errors"
	"fmt"
	"os"

	"go.yaml.in/yaml/v3"
)

var errMissingConfigFields = errors.New("input-file and output-file are required")

func loadConfig(path string) (string, string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", "", fmt.Errorf("read config file: %w", err)
	}

	var config map[string]string

	if err := yaml.Unmarshal(data, &config); err != nil {
		return "", "", fmt.Errorf("parse config file: %w", err)
	}

	inputFile := config["input-file"]
	outputFile := config["output-file"]

	if inputFile == "" || outputFile == "" {
		return "", "", errMissingConfigFields
	}

	return inputFile, outputFile, nil
}
