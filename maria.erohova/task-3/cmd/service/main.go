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
	dirPermission  = 0o755
	filePermission = 0o644
	jsonIndent     = "    "
	floatBitSize   = 64
)

var (
	ErrEmptyConfigPath = errors.New("config path is not set (use -config flag)")
	ErrEmptyInputFile  = errors.New("input-file is not set in config")
	ErrEmptyOutputFile = errors.New("output-file is not set in config")
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
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
	var cfg Config

	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, fmt.Errorf("read config: %w", err)
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return cfg, fmt.Errorf("parse config: %w", err)
	}

	cfg.InputFile = strings.TrimSpace(cfg.InputFile)
	cfg.OutputFile = strings.TrimSpace(cfg.OutputFile)

	if cfg.InputFile == "" {
		return cfg, ErrEmptyInputFile
	}

	if cfg.OutputFile == "" {
		return cfg, ErrEmptyOutputFile
	}

	return cfg, nil
}

func loadCurrencies(path string) ([]Currency, error) {
	data, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open input file: %w", err)
	}

	decoder := xml.NewDecoder(data)
	decoder.CharsetReader = charset.NewReaderLabel

	var curs ValCurs

	decodeErr := decoder.Decode(&curs)
	closeErr := data.Close()

	if decodeErr != nil {
		return nil, fmt.Errorf("decode xml: %w", decodeErr)
	}

	if closeErr != nil {
		return nil, fmt.Errorf("close input file: %w", closeErr)
	}

	result := make([]Currency, 0, len(curs.Valutes))

	for _, valute := range curs.Valutes {
		rawValue := strings.ReplaceAll(strings.TrimSpace(valute.Value), ",", ".")

		value, err := strconv.ParseFloat(rawValue, floatBitSize)
		if err != nil {
			return nil, fmt.Errorf("parse value %q: %w", valute.Value, err)
		}

		result = append(result, Currency{
			NumCode:  valute.NumCode,
			CharCode: strings.TrimSpace(valute.CharCode),
			Value:    value,
		})
	}

	return result, nil
}

func saveResult(path string, currencies []Currency) error {
	if err := os.MkdirAll(filepath.Dir(path), dirPermission); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}

	data, err := json.MarshalIndent(currencies, "", jsonIndent)
	if err != nil {
		return fmt.Errorf("encode json: %w", err)
	}

	if err := os.WriteFile(path, data, filePermission); err != nil {
		return fmt.Errorf("write output file: %w", err)
	}

	return nil
}

func run(configPath string) error {
	if strings.TrimSpace(configPath) == "" {
		return ErrEmptyConfigPath
	}

	cfg, err := loadConfig(configPath)
	if err != nil {
		return err
	}

	currencies, err := loadCurrencies(cfg.InputFile)
	if err != nil {
		return err
	}

	sort.SliceStable(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	return saveResult(cfg.OutputFile, currencies)
}

func main() {
	configPath := flag.String("config", "", "path to YAML configuration file")
	flag.Parse()

	if err := run(*configPath); err != nil {
		panic(err)
	}
}
