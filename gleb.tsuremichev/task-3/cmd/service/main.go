package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"flag"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type ValCurs struct {
	Valutes []ValuteXML `xml:"Valute"`
}

type ValuteXML struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type CurrencyJSON struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func readConfig(path string) Config {
	cfgBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var cfg Config

	err = yaml.Unmarshal(cfgBytes, &cfg)
	if err != nil {
		panic(err)
	}

	if cfg.InputFile == "" || cfg.OutputFile == "" {
		panic("empty input-file or output-file in config")
	}

	return cfg
}

func readXML(path string) ValCurs {
	xmlBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(xmlBytes))
	decoder.CharsetReader = charset.NewReaderLabel

	var valCurs ValCurs

	err = decoder.Decode(&valCurs)
	if err != nil {
		panic(err)
	}

	return valCurs
}

func convertCurrencies(valCurs ValCurs) []CurrencyJSON {
	currencies := make([]CurrencyJSON, 0, len(valCurs.Valutes))

	for _, valute := range valCurs.Valutes {
		valStr := strings.Replace(valute.Value, ",", ".", 1)

		valFloat, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			panic(err)
		}

		currencies = append(currencies, CurrencyJSON{
			NumCode:  valute.NumCode,
			CharCode: valute.CharCode,
			Value:    valFloat,
		})
	}

	return currencies
}

func sortCurrencies(currencies []CurrencyJSON) {
	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})
}

const (
	dirPerm  = 0o755
	filePerm = 0o600
)

func writeJSON(path string, currencies []CurrencyJSON) {
	err := os.MkdirAll(filepath.Dir(path), dirPerm)
	if err != nil {
		panic(err)
	}

	jsonBytes, err := json.MarshalIndent(currencies, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(path, jsonBytes, filePerm)
	if err != nil {
		panic(err)
	}
}

func main() {
	configPath := flag.String("config", "", "path to config")
	flag.Parse()

	if *configPath == "" {
		panic("config file is required")
	}

	cfg := readConfig(*configPath)
	valCurs := readXML(cfg.InputFile)
	currencies := convertCurrencies(valCurs)
	sortCurrencies(currencies)
	writeJSON(cfg.OutputFile, currencies)
}
