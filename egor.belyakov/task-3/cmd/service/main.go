package main

import (
	"flag"

	"github.com/Dadnsall/task-3/internal/config"
	"github.com/Dadnsall/task-3/internal/currency"
	"github.com/Dadnsall/task-3/internal/writer"
)

func main() {
	path := flag.String("config", "config.yaml", "path to yaml config")
	flag.Parse()

	cfg, err := config.Read(*path)
	if err != nil {
		panic(err)
	}

	valutes, err := currency.ReadXML(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	result, err := currency.ToCurrencies(valutes)
	if err != nil {
		panic(err)
	}

	currency.SortDesc(result)

	err = writer.WriteJSON(cfg.OutputFile, result)
	if err != nil {
		panic(err)
	}
}
