package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
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

const (
	floatBitSize = 64
	dirPerm      = 0o750
	filePerm     = 0o600
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

// Входные данные из XML
type ValCurs struct {
	XMLName xml.Name    `xml:"ValCurs"`
	Valutes []XMLValute `xml:"Valute"`
}

type XMLValute struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

// Выходные данные для JSON (ровно те поля и типы, что просят в Шаге 4)
type OutputValute struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func loadConfig(path string) (Config, error) {
	configBytes, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("read config file: %w", err)
	}

	var cfg Config

	if err := yaml.Unmarshal(configBytes, &cfg); err != nil {
		return Config{}, fmt.Errorf("unmarshal config: %w", err)
	}

	cfg.InputFile = strings.TrimSpace(cfg.InputFile)
	cfg.OutputFile = strings.TrimSpace(cfg.OutputFile)

	if cfg.InputFile == "" || cfg.OutputFile == "" {
		return Config{}, errors.New("invalid config: missing required fields")
	}

	return cfg, nil
}

func loadValutes(path string) ([]XMLValute, error) {
	xmlBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read xml file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(xmlBytes))
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch strings.ToLower(charset) {
		case "windows-1251", "cp1251":
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		case "utf-8", "":
			return input, nil
		default:
			return nil, fmt.Errorf("unsupported charset: %s", charset)
		}
	}

	var curs ValCurs

	if err := decoder.Decode(&curs); err != nil {
		return nil, fmt.Errorf("decode xml: %w", err)
	}

	if len(curs.Valutes) == 0 {
		return nil, errors.New("invalid xml: no valutes found")
	}

	return curs.Valutes, nil
}

func convertValutes(valutes []XMLValute) ([]OutputValute, error) {
	result := make([]OutputValute, 0, len(valutes))

	for _, v := range valutes {
		valStr := strings.ReplaceAll(v.Value, ",", ".")

		valFloat, err := strconv.ParseFloat(valStr, floatBitSize)
		if err != nil {
			return nil, fmt.Errorf("parse value %q: %w", v.Value, err)
		}

		result = append(result, OutputValute{
			NumCode:  v.NumCode,
			CharCode: v.CharCode,
			Value:    valFloat,
		})
	}

	return result, nil
}

func saveJSON(path string, data []OutputValute) error {
	outDir := filepath.Dir(path)
	if err := os.MkdirAll(outDir, dirPerm); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}

	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal json: %w", err)
	}

	if err := os.WriteFile(path, jsonBytes, filePerm); err != nil {
		return fmt.Errorf("write output file: %w", err)
	}

	return nil
}

func main() {
	configPath := flag.String("config", "config.yaml", "config file path")

	flag.Parse()

	cfg, err := loadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	valutes, err := loadValutes(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	result, err := convertValutes(valutes)
	if err != nil {
		panic(err)
	}

	sort.SliceStable(result, func(i, j int) bool {
		return result[i].Value > result[j].Value
	})

	if err := saveJSON(cfg.OutputFile, result); err != nil {
		panic(err)
	}
}
