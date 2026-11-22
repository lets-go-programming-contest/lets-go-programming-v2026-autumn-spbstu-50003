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
	configPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	cfgBytes, err := os.ReadFile(*configPath)
	must(err)

	var conf Config

	must(yaml.Unmarshal(cfgBytes, &conf))

	xmlBytes, err := os.ReadFile(conf.InputFile)
	must(err)

	decoder := xml.NewDecoder(bytes.NewReader(xmlBytes))
	decoder.CharsetReader = charset.NewReaderLabel

	var parsedXML ValCurs

	must(decoder.Decode(&parsedXML))

	items := make([]FinalValute, 0, len(parsedXML.Valutes))

	for _, item := range parsedXML.Valutes {
		var numCode int

		cleanNum := strings.TrimSpace(item.NumCode)
		if cleanNum != "" {
			numCode, err = strconv.Atoi(cleanNum)
			must(err)
		}

		valStr := strings.ReplaceAll(strings.TrimSpace(item.Value), ",", ".")
		valFloat, err := strconv.ParseFloat(valStr, 64)
		must(err)

		items = append(items, FinalValute{
			NumCode:  numCode,
			CharCode: strings.TrimSpace(item.CharCode),
			Value:    valFloat,
		})
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].Value > items[j].Value
	})

	dir := filepath.Dir(conf.OutputFile)
	must(os.MkdirAll(dir, 0o755))

	output, err := json.MarshalIndent(items, "", "  ")
	must(err)

	must(os.WriteFile(conf.OutputFile, output, 0o600))
}
