package models

import "github.com/DanielDanteDosSantosViana/hire.me/config"

type UrlEncurtada struct {
	Alias       string      `json:"alias,omitempty"`
	URL         string      `json:"url,omitempty"`
	Estatistica Estatistica `json:"estatistica,omitempty"`
	Acessos     int         `json:"acessos,omitempty"`
}

type Estatistica struct {
	TempoOperacao string `json:"tempoOperacao,omitempty"`
}

func NewUrlEncurtada(alias string, url string, tempoOperacao string, acessos int) UrlEncurtada {
	var estatistica Estatistica
	if tempoOperacao != "" {
		estatistica = Estatistica{tempoOperacao}
	}
	alias = config.Conf.Service.Host + alias
	return UrlEncurtada{alias, url, estatistica, acessos}
}
