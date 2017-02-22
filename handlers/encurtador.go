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

	var urlEncurtada models.UrlEncurtada
	url := r.URL.Query().Get("url")
	aliasCustomizado := r.URL.Query().Get("CUSTOM_ALIAS")
	log.Println(url)

	if url == "" {
		errorJ, _ := json.Marshal(models.NewErrorUrlNaoInformada())
		util.ResponseInternalError(w, errorJ)
		return
	}

	if aliasCustomizado != "" {
		retorno, err := e.encurtador.ExisteAlias(aliasCustomizado)
		if err != nil {
			errorJ, _ := json.Marshal(models.NewErrorInterno())
			util.ResponseInternalError(w, errorJ)
			return
		}
		if retorno {
			errorJ, _ := json.Marshal(models.NewErrorAliasExiste())
			util.ResponseInternalError(w, errorJ)
			return
		} 

		urlEncurtada, err := e.encurtarComAlias(url,aliasCustomizado)
		if err != nil {
			errorJ, _ := json.Marshal(models.NewErrorInterno())
			util.ResponseInternalError(w, errorJ)
			return
		}		
		urlEncurtadaJ, _:=json.Marshal(urlEncurtada)
		util.ResponseOK(w, urlEncurtadaJ)
		return
	}

	urlEncurtada, err := e.encurtar(url)
	if err != nil {
	    errorJ, _ := json.Marshal(models.NewErrorInterno())
		util.ResponseInternalError(w, errorJ)
		return
	}		
	urlEncurtadaJ, _:=json.Marshal(urlEncurtada)
	util.ResponseOK(w, urlEncurtadaJ)	
	return
}

func(e *Encurtador) encurtarComAlias(url string ,alias string)(models.UrlEncurtada, error) {
	return e.encurtador.Encurtar(url,alias)
}

func(e *Encurtador)encurtar(url string)(models.UrlEncurtada, error){
	return e.encurtador.Encurtar(url,"")
}

func (e *Encurtador) BuscarURL(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	alias := params["alias"]
	urlEncurtada, errRetorno := e.encurtador.BuscarPorAlias(alias)
	if errRetorno!=nil{
		log.Println(errRetorno)
		errorJ, _ := json.Marshal(errRetorno)
		util.ResponseOK(w, errorJ)
		return
	}
	http.Redirect(w, r, urlEncurtada.URL, http.StatusTemporaryRedirect)
	return	
}
