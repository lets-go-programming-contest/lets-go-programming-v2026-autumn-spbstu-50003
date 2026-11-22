package parser

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"

	"golang.org/x/net/html/charset"

	"golabs/internal/models"
)

func ParseCurrencies(path string) ([]models.Currency, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read input xml: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	var curs models.ValCurs

	if err := decoder.Decode(&curs); err != nil {
		return nil, fmt.Errorf("decode xml: %w", err)
	}

	if curs.Valutes == nil {
		return []models.Currency{}, nil
	}

	return curs.Valutes, nil
}
