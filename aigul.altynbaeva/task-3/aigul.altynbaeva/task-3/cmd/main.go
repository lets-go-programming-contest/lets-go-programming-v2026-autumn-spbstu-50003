package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"flag"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
	"go.yaml.in/yaml/v3"
)

type AppConfig struct {
	Input  string `yaml:"input-file"`
	Output string `yaml:"output-file"`
}

type ValCursXML struct {
	XMLName xml.Name    `xml:"ValCurs"`
	Valutes []ValuteXML `xml:"Valute"`
}

type ValuteXML struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type FinalValute struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	cfgPath := flag.String("config", "config.yaml", "path to configuration file")
	flag.Parse()

	cfgBytes, err := os.ReadFile(*cfgPath)
	must(err)

	var conf AppConfig
	must(yaml.Unmarshal(cfgBytes, &conf))

	xmlBytes, err := os.ReadFile(conf.Input)
	must(err)

	decoder := xml.NewDecoder(bytes.NewReader(xmlBytes))
	decoder.CharsetReader = charset.NewReaderLabel

	var parsedXML ValCursXML
	must(decoder.Decode(&parsedXML))

	var items []FinalValute
	for _, item := range parsedXML.Valutes {
		codeInt, err := strconv.Atoi(strings.TrimSpace(item.NumCode))
		must(err)

		valFloat, err := strconv.ParseFloat(strings.ReplaceAll(strings.TrimSpace(item.Value), ",", "."), 64)
		must(err)

		items = append(items, FinalValute{
			NumCode:  codeInt,
			CharCode: strings.TrimSpace(item.CharCode),
			Value:    valFloat,
		})
	}

	sort.SliceStable(items, func(i, j int) bool {
		return items[i].Value > items[j].Value
	})

	output, err := json.MarshalIndent(items, "", "  ")
	must(err)

	dir := filepath.Dir(conf.Output)
	must(os.MkdirAll(dir, 0755))

	must(os.WriteFile(conf.Output, output, 0644))
}
