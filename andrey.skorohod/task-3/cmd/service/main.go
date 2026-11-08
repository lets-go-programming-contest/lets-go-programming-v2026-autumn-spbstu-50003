package main

import (
	"cmp"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"slices"

	"go.yaml.in/yaml/v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type Valute struct {
	NumCode  int     `xml:"NumCode"`
	CharCode string  `xml:"CharCode"`
	Value    float64 `xml:"Value"`
}

type ValCursXML struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

type ValCursJson struct {
	Valutes []struct {
		NumCode  int     `json:"num_code"`
		CharCode string  `json:"char_code"`
		Value    float64 `json:"value"`
	}
}

func configParse(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("Error reading config file: %v", err)
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("Error parsing config file: %v", err)
	}

	return cfg, nil
}

func unmarchalXmlInput(config Config) (ValCursXML, error) {
	data, err := os.ReadFile(config.InputFile)
	if err != nil {
		return ValCursXML{}, fmt.Errorf("Error reading input file: %v", err)
	}

	var ValCurs *ValCursXML
	if err := xml.Unmarshal(data, ValCurs); err != nil {
		return ValCursXML{}, fmt.Errorf("Error parsing input file: %v", err)
	}

	return *ValCurs, nil
}

func (v *ValCursXML) processData() {
	slices.SortFunc(v.Valutes, func(a, b Valute) int {
		return cmp.Compare(b.Value, a.Value)
	})
}

func (v *ValCursXML) marshalJsonOutput(config Config) error {
	fileData, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return fmt.Errorf("Serialization error: %v", err)
	}

	err = os.WriteFile(config.OutputFile, fileData, 0644)
	if err != nil {
		return fmt.Errorf("Error writing output file: %v", err)
	}

	return nil
}

func run(configPath string) error {
	config, err := configParse(configPath)
	if err != nil {
		return err
	}

	valutes, err := unmarchalXmlInput(config)
	if err != nil {
		return err
	}

	valutes.processData()

	return valutes.marshalJsonOutput(config)
}

func main() {
	configPath := flag.String("config", "config.yaml", "Path to config file")
	flag.Parse()

	if err := run(*configPath); err != nil {
		panic(err)
	}
}
