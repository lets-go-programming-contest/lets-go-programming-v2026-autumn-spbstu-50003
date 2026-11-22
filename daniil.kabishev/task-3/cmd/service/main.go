package main

import (
	"flag"
	"sort"

	"golabs/internal/config"
	"golabs/internal/parser"
	"golabs/internal/writer"
)

const defaultConfigPath = "config.yaml"

func main() {
	var configPath string

	flag.StringVar(&configPath, "config", defaultConfigPath, "path to YAML configuration file")
	flag.Parse()

	cfg, err := config.Load(configPath)
	if err != nil {
		panic(err)
	}

	currencies, err := parser.ParseCurrencies(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	if err := writer.SaveCurrencies(cfg.OutputFile, currencies); err != nil {
		panic(err)
	}
}
