package main

import (
	"flag"
	"sort"

	"github.com/qwwwf/task-3/internal/config"
	"github.com/qwwwf/task-3/internal/currency"
)

func main() {
	configPath := flag.String(
		"config",
		"config.yaml",
		"path to configuration file",
	)

	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	rates, err := currency.LoadCurrencies(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	currencies, err := currency.ConvertCurrencies(rates.Currencies)
	if err != nil {
		panic(err)
	}

	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	if err := currency.SaveCurrencies(cfg.OutputFile, currencies); err != nil {
		panic(err)
	}
}
