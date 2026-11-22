package output

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	dirPerm  = 0o750
	filePerm = 0o600
)

func Save(path string, data any) error {
	err := os.MkdirAll(filepath.Dir(path), dirPerm)
	if err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}

	content, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("encode json: %w", err)
	}

	err = os.WriteFile(path, content, filePerm)
	if err != nil {
		return fmt.Errorf("write output file: %w", err)
	}

	return nil
}
