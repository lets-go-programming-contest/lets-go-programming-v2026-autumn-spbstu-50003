package config

import (
	"fmt"
	"os"

	"go.yaml.in/yaml/v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func ConfigParse(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("error reading config file: %w", err)
	}

	var cfg Config

	if err = yaml.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("error parsing config file: %w", err)
	}

	return cfg, nil
}
