package db

import (
	"database/sql"
	"github.com/DanielDanteDosSantosViana/hire.me/config"

	_ "github.com/go-sql-driver/mysql"
)

func NewSessionMysqlWriteDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.Conf.Db.MysqlWrite)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewSessionMysqlReadDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.Conf.Db.MysqlRead)
	if err != nil {
		return nil, err
	}
	return db, nil
}

