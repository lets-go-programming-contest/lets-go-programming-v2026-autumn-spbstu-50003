package main

import (
	"errors"
	"flag"

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
		return err
	}

	valutes, err := currency.ReadXML(cfg.InputFile)
	if err != nil {
		return err
	}

	currency.SortByValueDesc(valutes)

	return output.SaveJSON(cfg.OutputFile, valutes)
}

func main() {
	configPath := flag.String("config", defaultConfigPath, "path to config file")
	flag.Parse()

	err := run(*configPath)
	if err != nil {
		panic(err)
	}
}