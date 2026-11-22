package writer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/daniil-kabishev/task-3/internal/models"
)

const (
	dirPermission  = 0o755
	filePermission = 0o644
)

func SaveCurrencies(path string, items []models.Currency) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, dirPermission); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}

	data, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal output json: %w", err)
	}

	if err := os.WriteFile(path, data, filePermission); err != nil {
		return fmt.Errorf("write output file: %w", err)
	}

	return nil
}
