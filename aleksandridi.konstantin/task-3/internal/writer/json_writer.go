package writer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/neBr3aD/task-3/internal/model"
)

type JSONWriter struct{}

func NewJSONWriter() *JSONWriter {
	return &JSONWriter{}
}

func (jsonWriter *JSONWriter) Write(filePath string, currencies []model.Currency) error {
	outputDir := filepath.Dir(filePath)
	if outputDir != "" && outputDir != "." {
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return fmt.Errorf("не удалось создать директорию %s: %w", outputDir, err)
		}
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("не удалось создать выходной файл: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(currencies); err != nil {
		return fmt.Errorf("ошибка записи json: %w", err)
	}

	return nil
}
