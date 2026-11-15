package main

import (
	"bytes"
	"cmp"
	"encoding/json"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v3"
)

const (
	floatBitSize = 64
	dirPerm      = 0o755
	filePerm     = 0o644
)

var (
	errEmptyConfigPath = errors.New("config path is not specified")
	errEmptyField      = errors.New("input-file and output-file must be set")
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type ValCurs struct {
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func loadConfig(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("read config: %w", err)
	}

	var cfg Config

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("parse config: %w", err)
	}

	if cfg.InputFile == "" || cfg.OutputFile == "" {
		return Config{}, errEmptyField
	}

	return cfg, nil
}

func readValutes(path string) ([]Valute, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read input file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	var curs ValCurs

	err = decoder.Decode(&curs)
	if err != nil {
		return nil, fmt.Errorf("decode xml: %w", err)
	}

	return curs.Valutes, nil
}

func convertValutes(valutes []Valute) ([]Currency, error) {
	result := make([]Currency, 0, len(valutes))

	for _, valute := range valutes {
		text := strings.ReplaceAll(strings.TrimSpace(valute.Value), ",", ".")

		number, err := strconv.ParseFloat(text, floatBitSize)
		if err != nil {
			return nil, fmt.Errorf("parse value %q: %w", valute.Value, err)
		}

		result = append(result, Currency{
			NumCode:  valute.NumCode,
			CharCode: strings.TrimSpace(valute.CharCode),
			Value:    number,
		})
	}

	return result, nil
}

func sortCurrencies(currencies []Currency) {
	slices.SortFunc(currencies, func(first, second Currency) int {
		return cmp.Compare(second.Value, first.Value)
	})
}

func saveJSON(path string, currencies []Currency) error {
	err := os.MkdirAll(filepath.Dir(path), dirPerm)
	if err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}

	data, err := json.MarshalIndent(currencies, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal json: %w", err)
	}

	err = os.WriteFile(path, data, filePerm)
	if err != nil {
		return fmt.Errorf("write output file: %w", err)
	}

	return nil
}

func run(configPath string) error {
	if configPath == "" {
		return errEmptyConfigPath
	}

	cfg, err := loadConfig(configPath)
	if err != nil {
		return err
	}

	valutes, err := readValutes(cfg.InputFile)
	if err != nil {
		return err
	}

	currencies, err := convertValutes(valutes)
	if err != nil {
		return err
	}

	sortCurrencies(currencies)

	return saveJSON(cfg.OutputFile, currencies)
}

func main() {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	err := run(*configPath)
	if err != nil {
		panic(err)
	}
}
