package config

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type service struct {
	Port string `toml:"port"`
	Host string `toml:"host"`
}

type db struct {
	MysqlWrite string `toml:"mysqlwrite"`
	MysqlRead  string `toml:"mysqlread"`
}
type base struct {
	Alfabeto string `toml:"alfabeto"`
}
type config struct {
	Service service `toml:"service"`
	Db      db      `toml:"db"`
	Base    base    `toml:"base"`
}

var Conf config

func Load(config string) {
	if fileInfo, err := os.Stat(config); err != nil {
		if os.IsNotExist(err) {
			log.Panicf("Arquivo de configuração %v não existe.", config)
		} else {
			log.Panicf("Arquivo de configuração %v não pode iniciar. %v", config, err)
		}
	} else {
		if fileInfo.IsDir() {
			log.Panicf("%v é um diretório ", config)
		}
	}

	content, err := ioutil.ReadFile(config)
	if err != nil {
		log.Panicf("read configuration file error. %v", err)
	}
	content = bytes.TrimSpace(content)
	if err := toml.Unmarshal(content, &Conf); err != nil {
		log.Panicf("Erro falta ao tentar carregar o arquivo de configuração. %v", err)
	}
}
