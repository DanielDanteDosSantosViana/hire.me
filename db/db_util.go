package db
import (
	"database/sql"
	"log"
	"errors"
)
type DbUtil struct {
	writeDB *sql.DB
	readDB  *sql.DB
}
func NewDBUtil() *DbUtil {
	sessionW, err := NewSessionMysqlWriteDB()
	if err != nil {
		log.Panicf("Ocorreu um erro ao tentar abrir conexão com o dbWrite . %v", err)
	}

	err = sessionW.Ping()
	if err != nil {
		log.Panicf("Ocorreu um erro ao tentar verificar conexão dbWrite . %v", err)
	}

	sessionR, err := NewSessionMysqlReadDB()
	if err != nil {
		log.Panicf("Ocorreu um erro ao tentar abrir conexão com o dbRead . %v", err)
	}

	err = sessionR.Ping()
	if err != nil {
		log.Panicf("Ocorreu um erro ao tentar verificar conexão dbRead . %v", err)
	}
	return &DbUtil{sessionW, sessionR}
}


func (e *DbUtil) DeletePorAlias(alias string) error {
	var stmt *sql.Stmt
	stmt, err := e.writeDB.Prepare(`DELETE FROM url_encurtada where alias = ? `)
	if err != nil {
		log.Printf("Error na sequence. %v", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(alias)
	if err != nil {
		log.Printf("dbWrite error ao deletar. %v", err)
		return errors.New("dbWrite error ao deletar")
	}

	return nil
}

func (e *DbUtil) DeleteAll() error {
	var stmt *sql.Stmt
	stmt, err := e.writeDB.Prepare(`DELETE FROM url_encurtada`)
	if err != nil {
		log.Printf("Error na sequence. %v", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		log.Printf("dbWrite error ao deletar. %v", err)
		return errors.New("dbWrite error ao deletar")
	}

	return nil
}

func (e *DbUtil) Close() {
	if e.writeDB != nil {
		e.writeDB.Close()
		e.writeDB = nil
	}

	if e.readDB != nil {
		e.readDB.Close()
		e.readDB = nil
	}
}