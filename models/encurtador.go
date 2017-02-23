package models

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/DanielDanteDosSantosViana/hire.me/db"
	"github.com/DanielDanteDosSantosViana/hire.me/util"
)

type Encurtador struct {
	writeDB *sql.DB
	readDB  *sql.DB
}

func NewEncurtador() *Encurtador {
	sessionW, err := db.NewSessionMysqlWriteDB()
	if err != nil {
		log.Panicf("Ocorreu um erro ao tentar abrir conexão com o dbWrite . %v", err)
	}

	err = sessionW.Ping()
	if err != nil {
		log.Panicf("Ocorreu um erro ao tentar verificar conexão dbWrite . %v", err)
	}

	sessionR, err := db.NewSessionMysqlReadDB()
	if err != nil {
		log.Panicf("Ocorreu um erro ao tentar abrir conexão com o dbRead . %v", err)
	}

	err = sessionR.Ping()
	if err != nil {
		log.Panicf("Ocorreu um erro ao tentar verificar conexão dbRead . %v", err)
	}
	return &Encurtador{sessionW, sessionR}
}

func (e *Encurtador) Encurtar(url string, aliasCustomizavel string) (UrlEncurtada, error) {
	inicio := time.Now()
	if aliasCustomizavel != "" {
		tempoOperacao, err := e.criarUrlEncurtada(aliasCustomizavel, url, inicio)
		if err != nil {
			return UrlEncurtada{}, err
		}
		return NewUrlEncurtada(aliasCustomizavel, url, tempoOperacao, 0), nil
	}

	sequence, err := e.nextSequence()
	if err != nil {
		return UrlEncurtada{}, err
	}

	alias := util.InteiroParaString(sequence)
	tempoOperacao, err := e.criarUrlEncurtada(alias, url, inicio)
	if err != nil {
		return UrlEncurtada{}, err
	}

	return NewUrlEncurtada(alias, url, tempoOperacao, 0), nil
}
func (e *Encurtador) ExisteAlias(alias string) (bool, error) {
	var urlLonga string
	var rows *sql.Rows
	rows, err := e.readDB.Query(`SELECT url_longa FROM url_encurtada WHERE alias=?`, alias)
	if err != nil {
		log.Printf("dbRead error ao executar a query. %v", err)
		return false, errors.New("dbRead error ao executar a query")
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&urlLonga)
		if err != nil {
			log.Printf("dbRead error ao ler as linhas retornadas da consulta. %v", err)
			return false, errors.New("dbRead error ao ler as linhas retornadas da consulta")
		}
	}

	err = rows.Err()
	if err != nil {
		log.Printf("dbRead error na interacao da linha. %v", err)
		return false, errors.New("dbRead error na interacao da linha")
	}

	if urlLonga != "" {
		return true, nil
	}
	return false, nil
}

func (e *Encurtador) BuscarPorAlias(alias string) (UrlEncurtada, *ErrorReturn) {
	var urlLonga string
	var aliasRetorno string
	var tempoOperacao string
	var acessos int

	var rows *sql.Rows
	rows, err := e.readDB.Query(`SELECT url_longa,alias,tempo_operacao,acessos FROM url_encurtada WHERE alias=?`, alias)
	if err != nil {
		log.Printf("dbRead error ao executar a query. %v", err)
		return UrlEncurtada{}, NewErrorInterno()
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&urlLonga, &aliasRetorno, &tempoOperacao, &acessos)
		if err != nil {
			log.Printf("dbRead error ao ler as linhas retornadas da consulta. %v", err)
			return UrlEncurtada{}, NewErrorInterno()
		}
	}
	err = rows.Err()
	if err != nil {
		log.Printf("dbRead error na interacao da linha. %v", err)
		return UrlEncurtada{}, NewErrorInterno()
	}

	if aliasRetorno != "" {
		err := e.atualizarNumeroDeAcessos(alias, acessos)
		if err != nil {
			log.Printf("dbWrite na atualização . %v", err)
			return UrlEncurtada{}, NewErrorInterno()
		}

		return NewUrlEncurtada(aliasRetorno, urlLonga, tempoOperacao, acessos), nil
	}
	return UrlEncurtada{}, NewErrorAliasNaoEncontrado()
}

func (e *Encurtador) BuscarDezUrlsMaisAcessadas() ([]UrlEncurtada, *ErrorReturn) {
	var urls []UrlEncurtada
	var rows *sql.Rows
	rows, err := e.readDB.Query(`SELECT url_longa,alias,tempo_operacao,acessos FROM url_encurtada order by acessos desc limit 10   `)
	if err != nil {
		log.Printf("dbRead error ao executar a query. %v", err)
		return urls, NewErrorInterno()
	}

	defer rows.Close()
	for rows.Next() {
		var urlLonga string
		var aliasRetorno string
		var tempoOperacao string
		var acessos int
		err = rows.Scan(&urlLonga, &aliasRetorno, &tempoOperacao, &acessos)
		if err != nil {
			log.Printf("dbRead error ao ler as linhas retornadas da consulta. %v", err)
			return urls, NewErrorInterno()
		}
		urls = append(urls, NewUrlEncurtada(aliasRetorno, urlLonga, tempoOperacao, acessos))
	}
	err = rows.Err()
	if err != nil {
		log.Printf("dbRead error na interacao da linha. %v", err)
		return urls, NewErrorInterno()
	}

	return urls, nil
}

func (e *Encurtador) nextSequence() (sequence uint64, err error) {
	var stmt *sql.Stmt
	stmt, err = e.writeDB.Prepare(`REPLACE INTO sequence(stub) VALUES ("a")`)
	if err != nil {
		log.Printf("Error na sequence. %v", err)
		return 0, err
	}
	defer stmt.Close()

	var res sql.Result
	res, err = stmt.Exec()
	if err != nil {
		log.Printf("Error na chamada do REPLACE . %v", err)
		return 0, err
	}

	var lastID int64
	lastID, err = res.LastInsertId()
	if err != nil {
		log.Printf("Error ao tentar pegar o ultimo id inserido. %v", err)
		return 0, err
	} else {
		sequence = uint64(lastID)
		sequence -= 1
		return sequence, nil
	}
}

func (e *Encurtador) criarUrlEncurtada(alias string, url string, inicio time.Time) (string, error) {
	var stmt *sql.Stmt
	stmt, err := e.writeDB.Prepare(`INSERT INTO url_encurtada(url_longa, alias,tempo_operacao) VALUES(?, ?, ?)`)
	if err != nil {
		log.Printf("Error na sequence. %v", err)
		return "", err
	}
	defer stmt.Close()
	fim := time.Since(inicio)
	fimMs := fim.Nanoseconds() / 1000000
	tempoOperacao := strconv.Itoa(int(fimMs))
	tempoOperacao += "ms"
	_, err = stmt.Exec(url, alias, tempoOperacao)
	if err != nil {
		log.Printf("dbWrite error ao inserir. %v", err)
		return "", errors.New("dbWrite error ao inserir")
	}

	return tempoOperacao, nil
}

func (e *Encurtador) atualizarNumeroDeAcessos(alias string, acessos int) error {
	var stmt *sql.Stmt
	stmt, err := e.writeDB.Prepare(`UPDATE url_encurtada set acessos=? where alias = ?`)
	if err != nil {
		log.Printf("Error na Query %v", err)
		return err
	}

	defer stmt.Close()
	acessos++
	_, err = stmt.Exec(acessos, alias)
	if err != nil {
		log.Printf("dbWrite error ao atualizar. %v", err)
		return errors.New("dbWrite error ao inserir")
	}

	return nil
}

func (e *Encurtador) close() {
	if e.writeDB != nil {
		e.writeDB.Close()
		e.writeDB = nil
	}

	if e.readDB != nil {
		e.readDB.Close()
		e.readDB = nil
	}
}

