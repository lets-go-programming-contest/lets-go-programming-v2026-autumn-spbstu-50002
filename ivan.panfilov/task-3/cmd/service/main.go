package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Waffletsa1/task-3/internal/cbr"
	"github.com/Waffletsa1/task-3/internal/config"
	"github.com/Waffletsa1/task-3/internal/convert"
	sorting "github.com/Waffletsa1/task-3/internal/sort"
)

func main() {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	if *configPath == "" {
		fmt.Fprintln(os.Stderr, "config path is required")
		os.Exit(1)
	}

	cfg, err := config.Load(*configPath)
	if err != nil {
		panic(err)
	}

	currencies, err := cbr.ParseXML(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	sorting.ByValueDesc(currencies)

	if err := convert.SaveJSON(cfg.OutputFile, currencies); err != nil {
		panic(err)
	}
}
