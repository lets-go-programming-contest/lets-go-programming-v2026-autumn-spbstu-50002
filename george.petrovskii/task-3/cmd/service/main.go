package main

import (
	"errors"
	"flag"
	"fmt"

	"george.petrovskii/task-3/internal/config"
	"george.petrovskii/task-3/internal/currency"
	"george.petrovskii/task-3/internal/writer"
)

var errConfigPathNotSet = errors.New("config path is not set, use --config")

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	if *configPath == "" {
		return errConfigPathNotSet
	}

	cfg, err := config.Load(*configPath)
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	currencies, err := currency.Load(cfg.InputFile)
	if err != nil {
		return fmt.Errorf("load currencies: %w", err)
	}

	currency.SortByValueDesc(currencies)

	if err := writer.WriteJSON(cfg.OutputFile, currencies); err != nil {
		return fmt.Errorf("write result: %w", err)
	}

	return nil
}
