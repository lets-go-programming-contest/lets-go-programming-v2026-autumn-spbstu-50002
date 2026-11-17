package main

import (
	"flag"

	"github.com/KarKar07/task-3/internal/config"
	"github.com/KarKar07/task-3/internal/currency"
	"github.com/KarKar07/task-3/internal/writer"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to configuration file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		panic(err)
	}

	rates, err := currency.LoadXML(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	currencies, err := currency.ConvertCurrencies(rates.Currencies)
	if err != nil {
		panic(err)
	}

	currency.SortByValueDescending(currencies)

	if err := writer.SaveJSON(cfg.OutputFile, currencies); err != nil {
		panic(err)
	}
}
