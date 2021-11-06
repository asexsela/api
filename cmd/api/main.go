package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/asexsela/standart_web_server/internal/app/api"
)

var (
	configPath string = "configs/api.toml"
)

func init() {
	//Скажем, что наше приложение будет, на этапе запуска, получать путь до конфиг файла из внешнего мира
	flag.StringVar(&configPath, "path", "configs/api.toml", "Path to config file in .toml format")
}

func main() {
	//Initialization variable configPath
	flag.Parse()

	// server instance inicialization
	config := api.NewConfig()

	_, err := toml.DecodeFile(configPath, config) //Десириализуем содержимое .toml файла

	if err != nil {
		log.Println("Can not find config file. Using default values", err)
	}

	server := api.New(config)

	// api server start
	log.Fatal(server.Start())
}
