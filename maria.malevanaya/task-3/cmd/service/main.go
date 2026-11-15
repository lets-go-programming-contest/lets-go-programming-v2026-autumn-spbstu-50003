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

	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v3"
)

const (
	dirMode    = 0o755
	jsonIndent = "  "
	bitSize    = 64
)

var (
	errNoConfigFlag = errors.New("flag -config is not set")
	errEmptyPaths   = errors.New("input-file and output-file must be set")
)

type config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type rate float64

func (r *rate) UnmarshalText(text []byte) error {
	normalized := strings.ReplaceAll(strings.TrimSpace(string(text)), ",", ".")

	parsed, err := strconv.ParseFloat(normalized, bitSize)
	if err != nil {
		return fmt.Errorf("failed to parse rate %q: %w", string(text), err)
	}

	*r = rate(parsed)

	return nil
}

type currency struct {
	NumCode  int    `xml:"NumCode" json:"num_code"`
	CharCode string `xml:"CharCode" json:"char_code"`
	Value    rate   `xml:"Value" json:"value"`
}

type valCurs struct {
	Currencies []currency `xml:"Valute"`
}

func readConfig(path string) (config, error) {
	var cfg config

	file, err := os.Open(path)
	if err != nil {
		return cfg, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	if err := yaml.NewDecoder(file).Decode(&cfg); err != nil {
		return cfg, fmt.Errorf("failed to decode config file: %w", err)
	}

	cfg.InputFile = strings.TrimSpace(cfg.InputFile)
	cfg.OutputFile = strings.TrimSpace(cfg.OutputFile)

	if cfg.InputFile == "" || cfg.OutputFile == "" {
		return cfg, errEmptyPaths
	}

	return cfg, nil
}

func readCurrencies(path string) ([]currency, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open input file: %w", err)
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	var curs valCurs
	if err := decoder.Decode(&curs); err != nil {
		return nil, fmt.Errorf("failed to decode input file: %w", err)
	}

	if curs.Currencies == nil {
		return []currency{}, nil
	}

	return curs.Currencies, nil
}

func sortCurrencies(currencies []currency) {
	sort.SliceStable(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})
}

func writeCurrencies(path string, currencies []currency) error {
	if err := os.MkdirAll(filepath.Dir(path), dirMode); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", jsonIndent)

	if err := encoder.Encode(currencies); err != nil {
		return fmt.Errorf("failed to encode output file: %w", err)
	}

	return nil
}

func run(configPath string) error {
	if configPath == "" {
		return errNoConfigFlag
	}

	cfg, err := readConfig(configPath)
	if err != nil {
		return err
	}

	currencies, err := readCurrencies(cfg.InputFile)
	if err != nil {
		return err
	}

	sortCurrencies(currencies)

	return writeCurrencies(cfg.OutputFile, currencies)
}

func main() {
	configPath := flag.String("config", "", "path to yaml config file")
	flag.Parse()

	if err := run(*configPath); err != nil {
		panic(err)
	}
}
