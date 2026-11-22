package valutes

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

var (
	ErrBadEncoding = errors.New("unsupported XML encoding")
	ErrExtraData   = errors.New("unexpected data found after document end")
	ErrInvalidRate = errors.New("invalid rate value")
)

const dirPermissions = 0o755

type rawValute struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type xmlDocument struct {
	XMLName xml.Name    `xml:"ValCurs"`
	Items   []rawValute `xml:"Valute"`
}

type Record struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func ProcessXML(path string) ([]Record, error) {
	fileData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(fileData))
	decoder.CharsetReader = charset.NewReaderLabel

	var doc xmlDocument

	if err := decoder.Decode(&doc); err != nil {
		return nil, fmt.Errorf("failed to decode XML: %w", err)
	}

	if err := ensureNoExtraData(decoder); err != nil {
		return nil, err
	}

	records := make([]Record, 0, len(doc.Items))

	for i, item := range doc.Items {
		rec, err := convertItem(item)
		if err != nil {
			return nil, fmt.Errorf("error at element %d: %w", i+1, err)
		}

		records = append(records, rec)
	}

	sort.SliceStable(records, func(i, j int) bool {
		return records[i].Value > records[j].Value
	})

	return records, nil
}

func ensureNoExtraData(dec *xml.Decoder) error {
	for {
		token, err := dec.Token()
		if errors.Is(err, io.EOF) {
			return nil
		}

		if err != nil {
			return fmt.Errorf("token parsing error: %w", err)
		}

		switch t := token.(type) {
		case xml.Comment, xml.ProcInst:
			continue
		case xml.CharData:
			if len(bytes.TrimSpace(t)) == 0 {
				continue
			}
		}

		return ErrExtraData
	}
}

func convertItem(raw rawValute) (Record, error) {
	codeStr := strings.TrimSpace(raw.NumCode)
	if codeStr == "" {
		codeStr = "0"
	}

	code, err := strconv.Atoi(codeStr)
	if err != nil {
		return Record{}, fmt.Errorf("invalid NumCode %q: %w", raw.NumCode, err)
	}

	charCode := strings.TrimSpace(raw.CharCode)

	valStr := strings.ReplaceAll(strings.TrimSpace(raw.Value), ",", ".")
	valFloat, err := strconv.ParseFloat(valStr, 64)
	if err != nil {
		return Record{}, fmt.Errorf("%w %q: %w", ErrInvalidRate, raw.Value, err)
	}

	if math.IsNaN(valFloat) || math.IsInf(valFloat, 0) {
		return Record{}, fmt.Errorf("%w %q", ErrInvalidRate, raw.Value)
	}

	return Record{
		NumCode:  code,
		CharCode: charCode,
		Value:    valFloat,
	}, nil
}

func SaveJSON(path string, records []Record) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, dirPermissions); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	defer func() {
		_ = file.Close()
	}()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")

	if err := enc.Encode(records); err != nil {
		return fmt.Errorf("JSON encoding error: %w", err)
	}

	return nil
}
