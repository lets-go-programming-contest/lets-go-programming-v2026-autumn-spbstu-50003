package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/xmdnx/task-3/internal/config"
	"github.com/xmdnx/task-3/internal/valute"
)

func run(configPath string) error {
	cfg, err := config.Load(configPath)
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	valutes, err := valute.Load(cfg.InputFile)
	if err != nil {
		return fmt.Errorf("load valutes: %w", err)
	}

	valute.SortByValueDesc(valutes)

	if err := valute.SaveJSON(cfg.OutputFile, valutes); err != nil {
		return fmt.Errorf("save json: %w", err)
	}

	return nil
}

func main() {
	configPath := flag.String("config", "config.yaml", "config file path")

	flag.Parse()

	if err := run(*configPath); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
