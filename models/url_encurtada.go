package models

import "github.com/DanielDanteDosSantosViana/hire.me/config"

type UrlEncurtada struct {
	Alias       string      `json:"alias"`
	URL         string      `json:"url"`
	Estatistica Estatistica `json:"estatistica"`
}

type Estatistica struct {
	TempoOperacao            string `json:"tempoOperacao"`
}

func NewUrlEncurtada(alias string, url string, tempoOperacao string)(UrlEncurtada){
	estatistica := Estatistica{tempoOperacao}
	alias = config.Conf.Service.Host+alias
	return UrlEncurtada{alias,url,estatistica}
}
