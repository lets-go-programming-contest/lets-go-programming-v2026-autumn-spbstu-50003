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

	"golang.org/x/text/encoding/htmlindex"
	"gopkg.in/yaml.v3"
)

const (
	floatBitSize = 64
	dirPerm      = 0o750
	filePerm     = 0o600
)

var (
	errMissingConfigFields = errors.New("invalid config: missing required fields")
	errUnsupportedCharset  = errors.New("unsupported charset")
	errNoValutesFound      = errors.New("invalid xml: no valutes found")
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type Valute struct {
	NumCode  int        `xml:"NumCode" json:"num_code"`
	CharCode string     `xml:"CharCode" json:"char_code"`
	Value    FloatValue `xml:"Value" json:"value"`
}

type FloatValue float64

func (f *FloatValue) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	if err := d.DecodeElement(&v, &start); err != nil {
		return err
	}
	v = strings.ReplaceAll(v, ",", ".")
	parsed, err := strconv.ParseFloat(v, floatBitSize)
	if err != nil {
		return err
	}
	*f = FloatValue(parsed)
	return nil
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
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
		return Config{}, errMissingConfigFields
	}

	return cfg, nil
}

func loadValutes(path string) ([]Valute, error) {
	xmlBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read xml file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(xmlBytes))
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		enc, err := htmlindex.Get(charset)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", errUnsupportedCharset, charset)
		}
		return enc.NewDecoder().Reader(input), nil
	}

	var curs ValCurs
	if err := decoder.Decode(&curs); err != nil {
		return nil, fmt.Errorf("decode xml: %w", err)
	}

	if len(curs.Valutes) == 0 {
		return nil, errNoValutesFound
	}

	return curs.Valutes, nil
}

func saveJSON(path string, data []Valute) error {
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
	configPath := flag.String("config", "", "config file path")
	flag.Parse()

	if *configPath == "" {
		panic("config file path is required (use -config flag)")
	}

	cfg, err := loadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	valutes, err := loadValutes(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	sort.SliceStable(valutes, func(i, j int) bool {
		return valutes[i].Value > valutes[j].Value
	})

	if err := saveJSON(cfg.OutputFile, valutes); err != nil {
		panic(err)
	}

	fmt.Println("Успешно! Результат сохранен в", cfg.OutputFile)
}
