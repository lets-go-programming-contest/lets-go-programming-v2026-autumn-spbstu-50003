package output

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ezttmv/task-3/internal/currency"
)

const (
	dirPerm  = 0o755
	filePerm = 0o644
)

func SaveJSON(path string, valutes []currency.Valute) error {
	err := os.MkdirAll(filepath.Dir(path), dirPerm)
	if err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}

	data, err := json.MarshalIndent(valutes, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal json: %w", err)
	}

	err = os.WriteFile(path, data, filePerm)
	if err != nil {
		return fmt.Errorf("write output file: %w", err)
	}

	return nil
}
