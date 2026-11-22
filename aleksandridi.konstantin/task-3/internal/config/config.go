package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	ErrEmptyConfigPath     = errors.New("не указан путь до файла конфигурации")
	ErrMissingConfigFields = errors.New("в конфигурации должны быть заданы input-file и output-file")
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Load(filePath string) (*Config, error) {
	if filePath == "" {
		return nil, ErrEmptyConfigPath
	}

	rawBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла конфигурации: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(rawBytes, &cfg); err != nil {
		return nil, fmt.Errorf("ошибка разбора yaml: %w", err)
	}

	if cfg.InputFile == "" || cfg.OutputFile == "" {
		return nil, ErrMissingConfigFields
	}

	return &cfg, nil
}
