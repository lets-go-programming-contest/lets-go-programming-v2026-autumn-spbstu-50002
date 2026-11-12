package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to configuration file")
	flag.Parse()

	if err := run(*configPath); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(configPath string) error {
	inputFile, outputFile, err := loadConfig(configPath)
	if err != nil {
		return err
	}

	currencies, err := loadCurrencies(inputFile)
	if err != nil {
		return err
	}

	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].value > currencies[j].value
	})

	return saveCurrencies(outputFile, currencies)
}
