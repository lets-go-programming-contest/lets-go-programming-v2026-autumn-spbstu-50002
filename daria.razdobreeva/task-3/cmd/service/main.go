package main

import (
	"flag"

	"github.com/Dobriash113/task-3/internal/config"
	"github.com/Dobriash113/task-3/internal/currency"
	"github.com/Dobriash113/task-3/internal/writer"
)

const (
	outputDirPerm  = 0o755
	outputFilePerm = 0o644
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		panic(err)
	}

	currencies, err := currency.Load(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	currency.SortByValueDesc(currencies)

	err = writer.SaveJSON(cfg.OutputFile, currencies, outputDirPerm, outputFilePerm)
	if err != nil {
		panic(err)
	}
}
