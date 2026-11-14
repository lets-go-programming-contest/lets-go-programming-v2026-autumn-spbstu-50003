package main

import (
	"encoding/json"
	"encoding/xml"
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

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func main() {
	configPath := flag.String("config", "", "путь до YAML-файла конфигурации")
	flag.Parse()

	if strings.TrimSpace(*configPath) == "" {
		panic("не указан путь до файла конфигурации (флаг -config)")
	}

	cfg := loadConfig(*configPath)
	currencies := loadCurrencies(cfg.InputFile)

	sort.SliceStable(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	saveResult(cfg.OutputFile, currencies)
}

func loadConfig(path string) Config {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("чтение файла конфигурации: %w", err))
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic(fmt.Errorf("разбор файла конфигурации: %w", err))
	}

	cfg.InputFile = strings.TrimSpace(cfg.InputFile)
	cfg.OutputFile = strings.TrimSpace(cfg.OutputFile)

	if cfg.InputFile == "" {
		panic("в конфигурации не задано поле input-file")
	}
	if cfg.OutputFile == "" {
		panic("в конфигурации не задано поле output-file")
	}
	return cfg
}

func loadCurrencies(path string) []Currency {
	f, err := os.Open(path)
	if err != nil {
		panic(fmt.Errorf("открытие входного файла: %w", err))
	}
	defer f.Close()

	dec := xml.NewDecoder(f)
	dec.CharsetReader = charset.NewReaderLabel

	var curs ValCurs
	if err := dec.Decode(&curs); err != nil {
		panic(fmt.Errorf("разбор XML: %w", err))
	}

	result := make([]Currency, 0, len(curs.Valutes))
	for _, v := range curs.Valutes {
		num, err := strconv.Atoi(strings.TrimSpace(v.NumCode))
		if err != nil {
			panic(fmt.Errorf("некорректный NumCode %q: %w", v.NumCode, err))
		}
		val, err := strconv.ParseFloat(strings.ReplaceAll(strings.TrimSpace(v.Value), ",", "."), 64)
		if err != nil {
			panic(fmt.Errorf("некорректное Value %q у %s: %w", v.Value, v.CharCode, err))
		}
		result = append(result, Currency{
			NumCode:  num,
			CharCode: strings.TrimSpace(v.CharCode),
			Value:    val,
		})
	}
	return result
}

func saveResult(path string, currencies []Currency) {
	if dir := filepath.Dir(path); dir != "." {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			panic(fmt.Errorf("создание директории %q: %w", dir, err))
		}
	}

	f, err := os.Create(path)
	if err != nil {
		panic(fmt.Errorf("создание выходного файла: %w", err))
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "    ")
	if err := enc.Encode(currencies); err != nil {
		panic(fmt.Errorf("запись JSON: %w", err))
	}
}