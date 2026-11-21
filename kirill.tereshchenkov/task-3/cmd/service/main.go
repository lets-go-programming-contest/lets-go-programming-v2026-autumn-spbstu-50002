package main

import (
	"flag"
	"fmt"

	"github.com/Ministrius/task-3/internal/config"
	"github.com/Ministrius/task-3/internal/currency"
	"github.com/Ministrius/task-3/internal/writer"
)

func main() {
	configPath := flag.String(
		"config",
		"config.yaml",
		"Path to config",
	)

	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		panic(err)
	}

	currencies, err := currency.ParseXML(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	currency.SortByValueDescending(currencies)

	err = writer.WriteJSON(cfg.OutputFile, currencies)
	if err != nil {
		panic(err)
	}

	fmt.Println("Currency processing completed successfully!")
}
