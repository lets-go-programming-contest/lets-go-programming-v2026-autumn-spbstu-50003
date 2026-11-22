package main

import (
	"flag"

	"github.com/julybae/task-3/internal/config"
	"github.com/julybae/task-3/internal/currency"
	"github.com/julybae/task-3/internal/output"
)

const defaultConfigPath = "config.yaml"

func main() {
	configPath := flag.String("config", defaultConfigPath, "path to config file")

	flag.Parse()

	cfg, err := config.Read(*configPath)
	if err != nil {
		panic(err)
	}

	currencies, err := currency.Read(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	currency.Sort(currencies)

	err = output.Save(cfg.OutputFile, currencies)
	if err != nil {
		panic(err)
	}
}
