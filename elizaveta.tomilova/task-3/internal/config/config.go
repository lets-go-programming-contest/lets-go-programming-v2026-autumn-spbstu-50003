package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var errEmptyField = errors.New("input-file and output-file must be set")

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Load(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("read config: %w", err)
	}

	var cfg Config

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("parse config: %w", err)
	}

	if cfg.InputFile == "" || cfg.OutputFile == "" {
		return Config{}, errEmptyField
	}

	return cfg, nil
}