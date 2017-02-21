package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
)
type Recuperador struct {
	session *mgo.Session
}

func NewRecuperador() *Recuperador {
	return &Recuperador{db.NewSession()}
}


func (rec *Recuperador) BuscarURL(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(req)
	log.Println(params)
}