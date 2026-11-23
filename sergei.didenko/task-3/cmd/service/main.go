package main

import (
	"flag"
	"log"

	"github.com/MountenY/task-3/internal/config"
	"github.com/MountenY/task-3/internal/currency"
	"github.com/MountenY/task-3/internal/writer"
)

func main() {
	cfgPath := flag.String("config", "config.yaml", "path to configuration file")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	ratesXML, err := currency.LoadXML(cfg.InputFile)
	if err != nil {
		log.Fatalf("cannot parse XML: %v", err)
	}

	valutes, err := currency.Convert(ratesXML.Currencies)
	if err != nil {
		log.Fatalf("cannot convert currencies: %v", err)
	}

	currency.SortByValueDescending(valutes)

	if err := writer.SaveJSON(cfg.OutputFile, valutes); err != nil {
		log.Fatalf("cannot save JSON: %v", err)
	}
}
