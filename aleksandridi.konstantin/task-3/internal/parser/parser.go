package parser

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/neBr3aD/task-3/internal/model"
	"golang.org/x/net/html/charset"
)

type XMLParser struct{}

func NewXMLParser() *XMLParser {
	return &XMLParser{}
}

func (xmlParser *XMLParser) Parse(filePath string) ([]model.Currency, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть xml-файл: %w", err)
	}

	defer func() {
		_ = file.Close()
	}()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	var data model.ValCurs
	if err := decoder.Decode(&data); err != nil {
		return nil, fmt.Errorf("ошибка декодирования xml: %w", err)
	}

	return data.Currencies, nil
}
