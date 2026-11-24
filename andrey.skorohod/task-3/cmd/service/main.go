package main

import (
	"flag"
	"fmt"

	"github.com/a-sik/task-3/internal/cbr"
	"github.com/a-sik/task-3/internal/config"
)

const configName = "config.yaml"

func run(configPath string) error {
	config, err := config.ConfigParse(configPath)
	if err != nil {
		return fmt.Errorf("parse config %q: %w", configPath, err)
	}

	valutes, err := cbr.UnmarshalXMLInput(config.InputFile)
	if err != nil {
		return fmt.Errorf("parse input file %q: %w", config.InputFile, err)
	}

	valutes.ProcessData()

	if err := valutes.MarshalJSONOutput(config.OutputFile); err != nil {
		return fmt.Errorf("write output file %q: %w", config.OutputFile, err)
	}

	return nil
}

func main() {
	configPath := flag.String("config", configName, "Path to config file")
	flag.Parse()

	if err := run(*configPath); err != nil {
		panic(err)
	}
}
