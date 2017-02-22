package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewSessionMysql() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:1234@/encurtar_url?charset=utf8")
	if err != nil {
		return nil, err
	}
	return db, nil
}
