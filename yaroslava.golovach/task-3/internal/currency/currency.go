package currency

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

const dirPerm = 0o755

var ErrInvalidXML = errors.New("invalid xml signature or empty ValCurs")

type ValCurs struct {
	XMLName xml.Name `json:"-" xml:"ValCurs"`
	Valutes []Valute `json:"-" xml:"Valute"`
}

type Valute struct {
	NumCode  int     `json:"num_code" xml:"NumCode"`
	CharCode string  `json:"char_code" xml:"CharCode"`
	RawValue string  `json:"-"        xml:"Value"`
	Value    float64 `json:"value"    xml:"-"`
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

	if curs.XMLName.Local != "ValCurs" {
		return nil, fmt.Errorf("validate xml: %w", ErrInvalidXML)
	}

	result := make([]Valute, 0, len(curs.Valutes))

	for _, item := range curs.Valutes {
		normalized := strings.ReplaceAll(strings.TrimSpace(item.RawValue), ",", ".")
		parsedVal, err := strconv.ParseFloat(normalized, 64)
		if err != nil {
			return nil, fmt.Errorf("parse currency value %q: %w", item.RawValue, err)
		}

		item.CharCode = strings.TrimSpace(item.CharCode)
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
