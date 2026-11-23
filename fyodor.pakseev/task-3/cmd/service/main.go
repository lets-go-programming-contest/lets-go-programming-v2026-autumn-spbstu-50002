package main

import (
	"flag"

	"github.com/kbnzhpp/task-3/config"
	"github.com/kbnzhpp/task-3/src/currency"
	"github.com/kbnzhpp/task-3/src/parser"
	"github.com/kbnzhpp/task-3/src/writer"
)

func main() {
	configPath := flag.String("config", "config.yaml", "Path to the configuration file")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	currencies, err := parser.ParseXML(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	currency.SortCurrenciesByValue(currencies)

	if err := writer.SaveJSON(cfg.OutputFile, currencies); err != nil {
		panic(err)
	}
}
