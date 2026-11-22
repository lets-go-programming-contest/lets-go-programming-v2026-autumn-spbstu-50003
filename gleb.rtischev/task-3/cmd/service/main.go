package main

import (
	"flag"
	"fmt"

	"github.com/lets-go-programming-contest/lets-go-programming-v2026-autumn-spbstu-50003/gleb.rtischev/task-3/internal/config"
	"github.com/lets-go-programming-contest/lets-go-programming-v2026-autumn-spbstu-50003/gleb.rtischev/task-3/internal/valutes"
)

func main() {
	cfgPath := flag.String("config", "", "")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	if err != nil {
		panic(err)
	}

	records, err := valutes.ProcessXML(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	if err := valutes.SaveJSON(cfg.OutputFile, records); err != nil {
		panic(err)
	}

	fmt.Println("Готово")
}
