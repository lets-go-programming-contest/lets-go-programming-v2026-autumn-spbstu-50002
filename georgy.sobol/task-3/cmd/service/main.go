package main

import (
	"flag"

	"github.com/g-sobolevich/task-3/internal/config"
	"github.com/g-sobolevich/task-3/internal/valute"
)

func main() {
	configPath := flag.String(
		"config",
		"config.yaml",
		"path to configuration file",
	)

	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		panic(err)
	}

	costs, err := valute.LoadXML(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	valutes, err := valute.ConvertValutes(costs.Valutes)
	if err != nil {
		panic(err)
	}

	valute.SortInDecreasingOrder(valutes)

	if err := valute.ValutesToJSON(cfg.OutputFile, valutes); err != nil {
		panic(err)
	}
}
