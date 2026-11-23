package service

import (
	"sort"

	"github.com/neBr3aD/task-3/internal/model"
)

type CurrencyParser interface {
	Parse(filePath string) ([]model.Currency, error)
}

type CurrencyWriter interface {
	Write(filePath string, currencies []model.Currency) error
}

type Converter struct {
	parser CurrencyParser
	writer CurrencyWriter
}

func NewConverter(parser CurrencyParser, writer CurrencyWriter) *Converter {
	return &Converter{
		parser: parser,
		writer: writer,
	}
}

func (converter *Converter) Convert(inputPath string, outputPath string) error {
	currencies, err := converter.parser.Parse(inputPath)
	if err != nil {
		return err
	}

	if currencies == nil {
		currencies = []model.Currency{}
	}

	sort.Slice(currencies, func(firstIndex, secondIndex int) bool {
		return currencies[firstIndex].Value > currencies[secondIndex].Value
	})

	return converter.writer.Write(outputPath, currencies)
}
