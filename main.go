package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/DanielDanteDosSantosViana/hire.me/config"
	"github.com/DanielDanteDosSantosViana/hire.me/router"
)

func main() {
	configFile := flag.String("config", "conf.toml", "Path para o arquivo de configuração")
	flag.Parse()
	if *configFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	config.Load(*configFile)
	router.IniciarRotas()
	log.Println("Iniciou o serviço")
	http.ListenAndServe(config.Conf.Service.Port, nil)
}
