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
		return fmt.Errorf("ошибка чтения xml-элемента: %w", err)
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
	NumCode  int           `json:"num_code"  xml:"NumCode"`
	CharCode string        `json:"char_code" xml:"CharCode"`
	Value    CurrencyValue `json:"value"     xml:"Value"`
}

type ValCurs struct {
	XMLName    xml.Name   `xml:"ValCurs"`
	Currencies []Currency `xml:"Valute"`
}
