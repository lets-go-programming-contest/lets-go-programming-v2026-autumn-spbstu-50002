package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type Valute struct {
	NumCode    string `xml:"NumCode"`
	CharCode   string `xml:"CharCode"`
	Nominal    string `xml:"Nominal"`
	Name       string `xml:"Name"`
	Value      string `xml:"Value"`
	ValueFloat float64
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

type OutputValute struct {
	NumCode  string  `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func main() {
	configPath := flag.String("config", "", "путь к файлу конфигурации")
	flag.Parse()

	data, err := os.ReadFile(*configPath)
	if err != nil {
		panic(err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	fmt.Println("Input file:", config.InputFile)
	fmt.Println("Output file:", config.OutputFile)

	xmlData, err := os.ReadFile(config.InputFile)
	if err != nil {
		panic(err)
	}

	var valCurs ValCurs
	decoder := xml.NewDecoder(bytes.NewReader(xmlData))
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	}
	err = decoder.Decode(&valCurs)
	if err != nil {
		panic(err)
	}

	fmt.Println("Прочитано валют:", len(valCurs.Valutes))

	for i := range valCurs.Valutes {
		strValue := strings.Replace(valCurs.Valutes[i].Value, ",", ".", 1)
		f, err := strconv.ParseFloat(strValue, 64)
		if err != nil {
			panic(err)
		}
		valCurs.Valutes[i].ValueFloat = f
	}

	sort.Slice(valCurs.Valutes, func(i, j int) bool {
		return valCurs.Valutes[i].ValueFloat > valCurs.Valutes[j].ValueFloat
	})

	var output []OutputValute
	for _, v := range valCurs.Valutes {
		output = append(output, OutputValute{
			NumCode:  v.NumCode,
			CharCode: v.CharCode,
			Value:    v.ValueFloat,
		})
	}

	jsonData, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		panic(err)
	}

	dir := filepath.Dir(config.OutputFile)
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(config.OutputFile, jsonData, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Результат сохранён в:", config.OutputFile)
}
