package currency

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

const floatBitSize = 64

type Rate float64

func (r *Rate) UnmarshalText(text []byte) error {
	withDot := strings.ReplaceAll(strings.TrimSpace(string(text)), ",", ".")

	number, err := strconv.ParseFloat(withDot, floatBitSize)
	if err != nil {
		return fmt.Errorf("parse rate %q: %w", text, err)
	}

	*r = Rate(number)

	return nil
}

type Currency struct {
	NumCode  int    `xml:"NumCode" json:"num_code"`
	CharCode string `xml:"CharCode" json:"char_code"`
	Value    Rate   `xml:"Value" json:"value"`
}

type currencyList struct {
	Items []Currency `xml:"Valute"`
}

func Read(path string) ([]Currency, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read input file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	var list currencyList

	err = decoder.Decode(&list)
	if err != nil {
		return nil, fmt.Errorf("decode xml: %w", err)
	}

	return list.Items, nil
}

func Sort(currencies []Currency) {
	sort.SliceStable(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})
}
