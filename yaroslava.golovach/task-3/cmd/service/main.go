package main

import (
	"flag"
	"fmt"

	"yaroslava.golovach/task-3/internal/config"
	"yaroslava.golovach/task-3/internal/currency"
)

const defaultConfigPath = "config.yaml"

func main() {
	configPath := flag.String("config", defaultConfigPath, "path to yaml config file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		panic(fmt.Sprintf("config error: %v", err))
	}

	valutes, err := currency.LoadAndSort(cfg.InputFile)
	if err != nil {
		panic(fmt.Sprintf("currency processing error: %v", err))
	}

	if err := currency.SaveJSON(cfg.OutputFile, valutes); err != nil {
		panic(fmt.Sprintf("save result error: %v", err))
	}
}
