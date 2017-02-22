package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DanielDanteDosSantosViana/hire.me/models"
	"github.com/DanielDanteDosSantosViana/hire.me/util"
	"github.com/gorilla/mux"
)

type Encurtador struct {
	encurtador *models.Encurtador
}

func NewEncurtador(encurtador *models.Encurtador) *Encurtador {
	return &Encurtador{encurtador}
}

func (e *Encurtador) EncurtarURL(w http.ResponseWriter, r *http.Request) {

	url := r.URL.Query().Get("url")
	aliasCustomizado := r.URL.Query().Get("CUSTOM_ALIAS")

	if url == "" {
		errorJ, _ := json.Marshal(models.NewErrorUrlNaoInformada())
		util.ResponseInternalError(w, errorJ)
		return
	}

	if aliasCustomizado != "" {
		urlRetorno, err := e.encurtador.EncontrarAlias(aliasCustomizado)
		if err != nil {
			errorJ, _ := json.Marshal(models.NewErrorInterno())
			util.ResponseInternalError(w, errorJ)
			return
		}
		if urlRetorno.Alias == "" {
			errorJ, _ := json.Marshal(models.NewErrorAliasExiste())
			util.ResponseInternalError(w, errorJ)
			return
		}

	}

	e.encurtador.Encurtar(url)
}

/*
	urlRetorno, err := e.encurtador.Encurtar(url)
	if err != nil {
		log.Println("Error ao encurtar a URL")
		http.Error(w, "Error ao encurtar a URL", http.StatusInternalServerError)
		return
	}
	log.Print(urlRetorno)

*/

func (e *Encurtador) BuscarURL(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	log.Println(params)
}
