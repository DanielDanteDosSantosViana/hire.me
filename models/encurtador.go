package models

import (
	"database/sql"
	"log"

	"github.com/DanielDanteDosSantosViana/hire.me/db"
	"github.com/DanielDanteDosSantosViana/hire.me/util"
)

type Encurtador struct {
	writeDB *sql.DB
}

func NewEncurtador() *Encurtador {
	session, err := db.NewSessionMysql()
	if err != nil {
		log.Panicf("Ocorreu um erro ao tentar abrir conexão com o db . %v", err)
	}

	err = session.Ping()
	if err != nil {
		log.Panicf("Ocorreu um erro ao tentar verificar conexão db . %v", err)
	}
	return &Encurtador{session}
}

func (e *Encurtador) Encurtar(url string) (UrlEncurtada, error) {
	log.Println(util.InteiroParaString(8360115174570405475))
	return UrlEncurtada{}, nil
}
func (e *Encurtador) EncontrarAlias(alias string) (UrlEncurtada, error) {
	log.Println(alias)
	return UrlEncurtada{}, nil
}
