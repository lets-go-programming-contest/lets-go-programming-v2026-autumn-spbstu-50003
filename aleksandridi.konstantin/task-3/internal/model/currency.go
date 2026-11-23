package model

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type CurrencyValue float64

func (value *CurrencyValue) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var raw string
	if err := decoder.DecodeElement(&raw, &start); err != nil {
		return err
	}

	normalized := strings.Replace(strings.TrimSpace(raw), ",", ".", 1)
	parsed, err := strconv.ParseFloat(normalized, 64)
	if err != nil {
		return fmt.Errorf("не удалось преобразовать значение %q в число: %w", raw, err)
	}

	*value = CurrencyValue(parsed)
	return nil
}

type Currency struct {
	NumCode  int           `xml:"NumCode" json:"num_code"`
	CharCode string        `xml:"CharCode" json:"char_code"`
	Value    CurrencyValue `xml:"Value" json:"value"`
}

type ValCurs struct {
	XMLName    xml.Name   `xml:"ValCurs"`
	Currencies []Currency `xml:"Valute"`
}
