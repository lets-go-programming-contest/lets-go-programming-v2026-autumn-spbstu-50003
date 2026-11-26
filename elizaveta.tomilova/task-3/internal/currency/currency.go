package currency

import (
	"bytes"
	"cmp"
	"encoding/xml"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

const floatBitSize = 64

// Rate — курс валюты. В XML ЦБ записан с запятой ("43,8254").
type Rate float64

func (r *Rate) UnmarshalText(text []byte) error {
	clean := strings.ReplaceAll(strings.TrimSpace(string(text)), ",", ".")

	number, err := strconv.ParseFloat(clean, floatBitSize)
	if err != nil {
		return fmt.Errorf("parse rate %q: %w", text, err)
	}

	*r = Rate(number)

	return nil
}

// Valute — единая структура для чтения XML и записи JSON.
type Valute struct {
	NumCode  int    `json:"num_code"  xml:"NumCode"`
	CharCode string `json:"char_code" xml:"CharCode"`
	Value    Rate   `json:"value"     xml:"Value"`
}

type valCurs struct {
	Valutes []Valute `xml:"Valute"`
}

func ReadXML(path string) ([]Valute, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read input file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	var curs valCurs

	err = decoder.Decode(&curs)
	if err != nil {
		return nil, fmt.Errorf("decode xml: %w", err)
	}

	return curs.Valutes, nil
}

// SortByValueDesc сортирует валюты по убыванию курса.
func SortByValueDesc(valutes []Valute) {
	slices.SortFunc(valutes, func(first, second Valute) int {
		return cmp.Compare(second.Value, first.Value)
	})
}