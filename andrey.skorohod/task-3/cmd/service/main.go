package main

import (
	"bytes"
	"cmp"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"go.yaml.in/yaml/v3"
	"golang.org/x/net/html/charset"
)

const (
	dirPerm  os.FileMode = 0o755
	filePerm os.FileMode = 0o644
)

type CBRFloat float64

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type Valute struct {
	NumCode  string   `xml:"NumCode"`
	CharCode string   `xml:"CharCode"`
	Value    CBRFloat `xml:"Value"`
}

type ValCursXML struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

type ValuteJSON struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func (float *CBRFloat) UnmarshalXML(decod *xml.Decoder, start xml.StartElement) error {
	var str string

	if err := decod.DecodeElement(&str, &start); err != nil {
		return fmt.Errorf("decode Value element: %w", err)
	}

	str = strings.TrimSpace(strings.ReplaceAll(str, ",", "."))

	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return fmt.Errorf("parse Value %q: %w", str, err)
	}

	*float = CBRFloat(value)

	return nil
}

func configParse(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("error reading config file: %w", err)
	}

	var cfg Config

	if err = yaml.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("error parsing config file: %w", err)
	}

	return cfg, nil
}

func unmarshalXMLInput(config Config) (ValCursXML, error) {
	data, err := os.ReadFile(config.InputFile)
	if err != nil {
		return ValCursXML{}, fmt.Errorf("error reading input file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	var valCurs ValCursXML

	if err := decoder.Decode(&valCurs); err != nil {
		return ValCursXML{}, fmt.Errorf("error parsing input file: %w", err)
	}

	return valCurs, nil
}

func (val *ValCursXML) processData() {
	slices.SortFunc(val.Valutes, func(a, b Valute) int {
		return cmp.Compare(b.Value, a.Value)
	})
}

func (val *ValCursXML) marshalJSONOutput(config Config) error {
	outputJSON := make([]ValuteJSON, 0, len(val.Valutes))

	for _, valute := range val.Valutes {
		numCodeStr := strings.TrimSpace(valute.NumCode)
		numCodeInt := 0

		if numCodeStr != "" {
			parsed, err := strconv.Atoi(numCodeStr)
			if err != nil {
				return fmt.Errorf("parse NumCode %q: %w", valute.NumCode, err)
			}

			numCodeInt = parsed
		}

		outputJSON = append(outputJSON, ValuteJSON{
			NumCode:  numCodeInt,
			CharCode: strings.TrimSpace(valute.CharCode),
			Value:    float64(valute.Value),
		})
	}

	fileData, err := json.MarshalIndent(outputJSON, "", "    ")
	if err != nil {
		return fmt.Errorf("serialization error: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(config.OutputFile), dirPerm); err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}

	err = os.WriteFile(config.OutputFile, fileData, filePerm)
	if err != nil {
		return fmt.Errorf("error writing output file: %w", err)
	}

	return nil
}

func run(configPath string) error {
	config, err := configParse(configPath)
	if err != nil {
		return err
	}

	valutes, err := unmarshalXMLInput(config)
	if err != nil {
		return err
	}

	valutes.processData()

	return valutes.marshalJSONOutput(config)
}

func main() {
	configPath := flag.String("config", "config.yaml", "Path to config file")
	flag.Parse()

	if err := run(*configPath); err != nil {
		panic(err)
	}
}
