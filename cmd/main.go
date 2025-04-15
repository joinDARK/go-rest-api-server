package main

import (
	"flag"
	"log"

	"go-rest-api/internal/config"
	"go-rest-api/internal/server"

	"github.com/BurntSushi/toml"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "../configs/server.toml", "Путь к конфигу")
}

func main() {
	flag.Parse()
	
	cfg := config.NewServerConfig()
	_, errCfg := toml.DecodeFile(configPath, cfg)
	if errCfg != nil {
		log.Fatal(errCfg)
	}
	
	s := server.New(cfg)
	err := s.Start()
	if err != nil {
		log.Fatal(err)
	}
}