package main

import (
	"encoding/json"
	"flag"

	"github.com/ssealzz/task-3/config"
	"github.com/ssealzz/task-3/internal/currency"
	"github.com/ssealzz/task-3/internal/writer"
)

func main() {
	configPath := flag.String("config", "config.yml", "path to config file")

	flag.Parse()

	var (
		projectConfig *config.Config
		err           error
	)
	projectConfig, err = config.Parse(*configPath)
	if err != nil {
		panic(err)
	}

	sourceFilePath := projectConfig.InputFile
	resultFilePath := projectConfig.OutputFile
	parser := currency.NewParser(sourceFilePath)
	valCurs1, err := parser.Parse()
	if err != nil {
		panic(err)
	}

	currency.Sort(*valCurs1)
	output, err := json.MarshalIndent(valCurs1.Valutes, "", "  ")
	if err != nil {
		panic(err)
	}

	err = writer.CreateWriteAll(resultFilePath, output)
	if err != nil {
		panic(err)
	}
}
