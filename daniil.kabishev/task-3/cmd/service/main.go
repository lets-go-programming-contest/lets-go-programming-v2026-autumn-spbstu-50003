package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	defaultConfigPath = "config.yaml"
	dirPermission     = 0755
	filePermission    = 0644
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type ValCursXML struct {
	XMLName xml.Name    `xml:"ValCurs"`
	Valutes []ValuteXML `xml:"Valute"`
}

type ValuteXML struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type CurrencyResult struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func loadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func parseCurrencies(path string) ([]CurrencyResult, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var curs ValCursXML
	if err := xml.Unmarshal(data, &curs); err != nil {
		return nil, err
	}

	results := make([]CurrencyResult, 0, len(curs.Valutes))
	for _, valute := range curs.Valutes {
		sanitizedValue := strings.ReplaceAll(valute.Value, ",", ".")
		parsedValue, err := strconv.ParseFloat(sanitizedValue, 64)
		if err != nil {
			return nil, err
		}

		results = append(results, CurrencyResult{
			NumCode:  valute.NumCode,
			CharCode: valute.CharCode,
			Value:    parsedValue,
		})
	}

	return results, nil
}

func saveResults(path string, items []CurrencyResult) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, dirPermission); err != nil {
		return err
	}

	data, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, filePermission)
}

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", defaultConfigPath, "path to YAML configuration file")
	flag.Parse()

	cfg, err := loadConfig(configPath)
	if err != nil {
		panic(err)
	}

	currencies, err := parseCurrencies(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	if err := saveResults(cfg.OutputFile, currencies); err != nil {
		panic(err)
	}
}
