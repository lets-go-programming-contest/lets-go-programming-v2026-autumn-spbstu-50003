package cbr

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"

	"golang.org/x/net/html/charset"
)

func UnmarshalXMLInput(inputFile string) (ValCurs, error) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return ValCurs{}, fmt.Errorf("error reading input file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	var valCurs ValCurs

	if err := decoder.Decode(&valCurs); err != nil {
		return ValCurs{}, fmt.Errorf("error parsing input file: %w", err)
	}

	return valCurs, nil
}
