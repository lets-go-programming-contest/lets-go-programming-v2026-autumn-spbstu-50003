package cbr

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	dirPerm  os.FileMode = 0o755
	filePerm os.FileMode = 0o644
)

func (val *ValCurs) MarshalJSONOutput(outputFile string) error {
	outputJSON, err := json.MarshalIndent(val.Valutes, "", "    ")
	if err != nil {
		return fmt.Errorf("serialization error: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(outputFile), dirPerm); err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}

	err = os.WriteFile(outputFile, outputJSON, filePerm)
	if err != nil {
		return fmt.Errorf("error writing output file: %w", err)
	}

	return nil
}
