package main

import (
	"encoding/json"
	"flag"
	"os"
	"path/filepath"
	"sort"

	"github.com/hjnhjnhjncf/task-2-1/xml"
	"github.com/hjnhjnhjncf/task-2-1/yaml"
)

const (
	dirPerm  = 0o700
	filePerm = 0o600
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to configuration file")
	flag.Parse()

	cfg, err := yaml.Load(*configPath)
	must(err)
	currencies, err := xml.Load(cfg.Input)
	must(err)

	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	data, err := json.MarshalIndent(currencies, "", "    ")
	must(err)
	must(os.MkdirAll(filepath.Dir(cfg.Output), dirPerm))
	must(os.WriteFile(cfg.Output, data, filePerm))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
