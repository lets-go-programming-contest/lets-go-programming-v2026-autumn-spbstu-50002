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
	must(os.MkdirAll(filepath.Dir(cfg.Output), 0755))
	must(os.WriteFile(cfg.Output, data, 0644))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
