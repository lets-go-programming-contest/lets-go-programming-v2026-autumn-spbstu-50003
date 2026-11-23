package valute

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

const (
	floatBitSize = 64
	dirPerm      = 0o750
	filePerm     = 0o600
)

var errNoValutesFound = errors.New("invalid xml: no valutes found")

type Float float64

func (f *Float) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var valStr string

	if err := d.DecodeElement(&valStr, &start); err != nil {
		return fmt.Errorf("decode float element: %w", err)
	}

	valStr = strings.ReplaceAll(strings.TrimSpace(valStr), ",", ".")

	valFloat, err := strconv.ParseFloat(valStr, floatBitSize)
	if err != nil {
		return fmt.Errorf("parse float %q: %w", valStr, err)
	}

	*f = Float(valFloat)

	return nil
}

type Valute struct {
	NumCode  int    `json:"num_code"  xml:"NumCode"`
	CharCode string `json:"char_code" xml:"CharCode"`
	Value    Float  `json:"value"     xml:"Value"`
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

func Load(path string) ([]Valute, error) {
	xmlBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read xml file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(xmlBytes))
	decoder.CharsetReader = charset.NewReaderLabel

	var curs ValCurs

	if err := decoder.Decode(&curs); err != nil {
		return nil, fmt.Errorf("decode xml: %w", err)
	}

	if len(curs.Valutes) == 0 {
		return nil, errNoValutesFound
	}

	return curs.Valutes, nil
}

func SortByValueDesc(valutes []Valute) {
	sort.SliceStable(valutes, func(i, j int) bool {
		return valutes[i].Value > valutes[j].Value
	})
}

func SaveJSON(path string, valutes []Valute) error {
	outDir := filepath.Dir(path)

	if err := os.MkdirAll(outDir, dirPerm); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}

	jsonBytes, err := json.MarshalIndent(valutes, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal json: %w", err)
	}

	if err := os.WriteFile(path, jsonBytes, filePerm); err != nil {
		return fmt.Errorf("write output file: %w", err)
	}

	return nil
}
