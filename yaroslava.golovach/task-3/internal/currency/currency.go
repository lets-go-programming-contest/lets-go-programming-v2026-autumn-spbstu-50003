package currency

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

const dirPerm = 0o755

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs" json:"-"`
	Valutes []Valute `xml:"Valute" json:"-"`
}

type Valute struct {
	NumCode  int     `xml:"NumCode" json:"num_code"`
	CharCode string  `xml:"CharCode" json:"char_code"`
	RawValue string  `xml:"Value" json:"-"`
	Value    float64 `xml:"-" json:"value"`
}

func LoadAndSort(inputPath string) ([]Valute, error) {
	rawData, err := os.ReadFile(inputPath)
	if err != nil {
		return nil, fmt.Errorf("read input xml file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(rawData))
	decoder.CharsetReader = charset.NewReaderLabel

	var curs ValCurs
	if err := decoder.Decode(&curs); err != nil {
		return nil, fmt.Errorf("decode xml data: %w", err)
	}

	if curs.XMLName.Local != "ValCurs" || len(curs.Valutes) == 0 {
		return nil, fmt.Errorf("invalid xml signature or empty ValCurs")
	}

	result := make([]Valute, 0, len(curs.Valutes))
	for _, item := range curs.Valutes {
		normalized := strings.ReplaceAll(strings.TrimSpace(item.RawValue), ",", ".")
		parsedVal, err := strconv.ParseFloat(normalized, 64)
		if err != nil {
			return nil, fmt.Errorf("parse currency value %q: %w", item.RawValue, err)
		}

		item.CharCode = strings.TrimSpace(item.CharCode)
		if item.NumCode <= 0 || item.CharCode == "" {
			return nil, fmt.Errorf("invalid currency fields: num_code=%d char_code=%q", item.NumCode, item.CharCode)
		}

		item.Value = parsedVal
		result = append(result, item)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Value > result[j].Value
	})

	return result, nil
}

func SaveJSON(outputPath string, items []Valute) error {
	dir := filepath.Dir(outputPath)
	if dir != "" && dir != "." {
		if err := os.MkdirAll(dir, dirPerm); err != nil {
			return fmt.Errorf("create output directory: %w", err)
		}
	}

	payload, err := json.MarshalIndent(items, "", "    ")
	if err != nil {
		return fmt.Errorf("marshal json: %w", err)
	}

	payload = append(payload, '\n')

	if err := os.WriteFile(outputPath, payload, dirPerm); err != nil {
		return fmt.Errorf("write output json file: %w", err)
	}

	return nil
}
