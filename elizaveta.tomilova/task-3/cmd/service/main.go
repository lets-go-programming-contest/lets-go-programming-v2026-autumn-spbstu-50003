package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/ezttmv/task-3/internal/config"
	"github.com/ezttmv/task-3/internal/currency"
	"github.com/ezttmv/task-3/internal/output"
)

const defaultConfigPath = "config.yaml"

var errEmptyConfigPath = errors.New("config path is not specified")

func run(configPath string) error {
	if configPath == "" {
		return errEmptyConfigPath
	}

	cfg, err := config.Load(configPath)
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	valutes, err := currency.ReadXML(cfg.InputFile)
	if err != nil {
		return fmt.Errorf("read xml: %w", err)
	}

	currency.SortByValueDesc(valutes)

	err = output.SaveJSON(cfg.OutputFile, valutes)
	if err != nil {
		return fmt.Errorf("save json: %w", err)
	}

	return nil
}

func main() {
	configPath := flag.String("config", defaultConfigPath, "path to config file")
	flag.Parse()

	err := run(*configPath)
	if err != nil {
		panic(err)
	}
}
