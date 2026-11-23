package main

import (
	"flag"

	"github.com/neBr3aD/task-3/internal/config"
	"github.com/neBr3aD/task-3/internal/parser"
	"github.com/neBr3aD/task-3/internal/service"
	"github.com/neBr3aD/task-3/internal/writer"
)

func main() {
	var configPath string

	flag.StringVar(&configPath, "config", "", "путь до файла конфигурации")
	flag.Parse()

	appConfig, err := config.Load(configPath)
	if err != nil {
		panic(err)
	}

	xmlParser := parser.NewXMLParser()
	jsonWriter := writer.NewJSONWriter()
	converter := service.NewConverter(xmlParser, jsonWriter)

	if err := converter.Convert(appConfig.InputFile, appConfig.OutputFile); err != nil {
		panic(err)
	}
}
