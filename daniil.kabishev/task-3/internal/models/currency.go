package models

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type Float float64

func (f *Float) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var raw string

	if err := d.DecodeElement(&raw, &start); err != nil {
		return fmt.Errorf("decode float element: %w", err)
	}

	sanitized := strings.ReplaceAll(raw, ",", ".")

	val, err := strconv.ParseFloat(sanitized, 64)
	if err != nil {
		return fmt.Errorf("parse float value: %w", err)
	}

	*f = Float(val)

	return nil
}

type ValCurs struct {
	XMLName xml.Name   `xml:"ValCurs"`
	Valutes []Currency `xml:"Valute"`
}

type Currency struct {
	NumCode  int    `json:"num_code"  xml:"NumCode"`
	CharCode string `json:"char_code" xml:"CharCode"`
	Value    Float  `json:"value"     xml:"Value"`
}
