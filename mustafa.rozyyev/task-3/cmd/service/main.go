package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/r-mustafa-03/task-3/internal/config"
	"github.com/r-mustafa-03/task-3/internal/currency"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to YAML configuration")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Printf("load configuration: %v", err)
		os.Exit(1)
	}

	source, err := currency.ReadXML(cfg.InputFile)
	if err != nil {
		log.Printf("read currencies: %v", err)
		os.Exit(1)
	}

	rates, err := currency.Convert(source)
	if err != nil {
		log.Printf("convert currencies: %v", err)
		os.Exit(1)
	}

	sort.SliceStable(rates, func(i, j int) bool {
		return rates[i].Value > rates[j].Value
	})

	if err := currency.WriteJSON(cfg.OutputFile, rates); err != nil {
		log.Printf("write currencies: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Converted %d currencies to %s\n", len(rates), cfg.OutputFile)
}
