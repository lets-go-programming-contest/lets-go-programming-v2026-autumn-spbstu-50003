package main

import (
 "bytes"
 "cmp"
 "encoding/json"
 "encoding/xml"
 "errors"
 "flag"
 "fmt"
 "io"
 "os"
 "path/filepath"
 "slices"
 "strconv"
 "strings"

 "golang.org/x/text/encoding/charmap"
 "gopkg.in/yaml.v3"
)

const (
 folderPermissions = 0o755
 filePermissions   = 0o600
)

var errEmptyConfigPaths = errors.New("empty paths in config")

type AppConfig struct {
 Input  string `yaml:"input-file"`
 Output string `yaml:"output-file"`
}

type RawValCurs struct {
 XMLName xml.Name    `xml:"ValCurs"`
 Items   []RawValute `xml:"Valute"`
}

type RawValute struct {
 NumCode  string `xml:"NumCode"`
 CharCode string `xml:"CharCode"`
 Value    string `xml:"Value"`
}

type ResultCurrency struct {
 NumCode  int     `json:"num_code"`
 CharCode string  `json:"char_code"`
 Value    float64 `json:"value"`
}

func loadConfig(path string) (AppConfig, error) {
 raw, err := os.ReadFile(path)
 if err != nil {
  return AppConfig{}, fmt.Errorf("read config: %w", err)
 }

 var cfg AppConfig
 if err := yaml.Unmarshal(raw, &cfg); err != nil {
  return AppConfig{}, fmt.Errorf("unmarshal yaml: %w", err)
 }

 if cfg.Input == "" || cfg.Output == "" {
  return AppConfig{}, errEmptyConfigPaths
 }

 return cfg, nil
}

func parseCurrenciesXML(path string) ([]ResultCurrency, error) {
 data, err := os.ReadFile(path)
 if err != nil {
  return nil, fmt.Errorf("read xml: %w", err)
 }

 decoder := xml.NewDecoder(bytes.NewReader(data))
 decoder.CharsetReader = func(charset string, r io.Reader) (io.Reader, error) {
  if strings.EqualFold(charset, "windows-1251") {
   return charmap.Windows1251.NewDecoder().Reader(r), nil
  }

  return r, nil
 }

 var doc RawValCurs
 if err := decoder.Decode(&doc); err != nil {
  return nil, fmt.Errorf("decode xml: %w", err)
 }

 result := make([]ResultCurrency, 0, len(doc.Items))

 for _, entry := range doc.Items {
  numCode := 0

  trimmedNum := strings.TrimSpace(entry.NumCode)
  if trimmedNum != "" {
   numCode, err = strconv.Atoi(trimmedNum)
   if err != nil {
    return nil, fmt.Errorf("parse num code: %w", err)
   }
  }

  normalizedVal := strings.ReplaceAll(strings.TrimSpace(entry.Value), ",", ".")

  rate, err := strconv.ParseFloat(normalizedVal, 64)
  if err != nil {
   return nil, fmt.Errorf("parse value: %w", err)
  }

  result = append(result, ResultCurrency{
   NumCode:  numCode,
   CharCode: strings.TrimSpace(entry.CharCode),
   Value:    rate,
  })
 }

 return result, nil
}

func saveJSON(path string, items []ResultCurrency) error {
 if err := os.MkdirAll(filepath.Dir(path), folderPermissions); err != nil {
  return fmt.Errorf("create output dir: %w", err)
 }

 encoded, err := json.MarshalIndent(items, "", "  ")
 if err != nil {
  return fmt.Errorf("marshal json: %w", err)
 }

 if err := os.WriteFile(path, encoded, filePermissions); err != nil {
  return fmt.Errorf("write output file: %w", err)
 }

 return nil
}

func main() {
 cfgFlag := flag.String("config", "config.yaml", "config file path")
 flag.Parse()

 cfg, err := loadConfig(*cfgFlag)
 if err != nil {
  panic(err)
 }

 currencies, err := parseCurrenciesXML(cfg.Input)
 if err != nil {
  panic(err)
 }

 slices.SortFunc(currencies, func(a, b ResultCurrency) int {
  return cmp.Compare(b.Value, a.Value)
 })

 if err := saveJSON(cfg.Output, currencies); err != nil {
  panic(err)
 }
}