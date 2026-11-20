package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	dirPermissions  = 0o750
	filePermissions = 0o600
)

var errInvalidCurrencyCodes = errors.New("invalid currency codes")

type config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type currencyRates struct {
	XMLName xml.Name      `xml:"ValCurs"`
	Items   []xmlCurrency `xml:"Valute"`
}

type xmlCurrency struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func main() {
	configPath := flag.String("config", "config.yaml", "path to YAML config")
	flag.Parse()

	data, err := os.ReadFile(*configPath)
	if err != nil {
		panic(err)
	}

	var cfg config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic(err)
	}

	inputPath := strings.TrimSpace(cfg.InputFile)
	outputPath := strings.TrimSpace(cfg.OutputFile)

	if inputPath == "" || outputPath == "" {
		panic("input-file and output-file are required")
	}

	currencies, err := readCurrencies(inputPath)
	if err != nil {
		panic(err)
	}

	result, err := json.MarshalIndent(currencies, "", "  ")
	if err != nil {
		panic(err)
	}
	if err := os.MkdirAll(filepath.Dir(outputPath), dirPermissions); err != nil {
		panic(err)
	}
	if err := os.WriteFile(outputPath, result, filePermissions); err != nil {
		panic(err)
	}
}

func readCurrencies(path string) ([]currency, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read XML file %q: %w", path, err)
	}

	var rates currencyRates
	if err := xml.Unmarshal(data, &rates); err != nil {
		return nil, fmt.Errorf("decode XML file %q: %w", path, err)
	}

	currencies := make([]currency, len(rates.Items))
	for index, item := range rates.Items {
		code := strings.TrimSpace(item.CharCode)
		if item.NumCode <= 0 || code == "" {
			return nil, fmt.Errorf("%w: %d, %q", errInvalidCurrencyCodes, item.NumCode, code)
		}

		rawValue := strings.ReplaceAll(strings.TrimSpace(item.Value), ",", ".")

		value, err := strconv.ParseFloat(rawValue, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid Value for %s: %w", code, err)
		}

		currencies[index] = currency{
			NumCode:  item.NumCode,
			CharCode: code,
			Value:    value,
		}
	}

	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	return currencies, nil
}
