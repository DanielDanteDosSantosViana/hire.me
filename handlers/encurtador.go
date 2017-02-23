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
		util.ResponseBadRequest(w, errorJ)
		return
	}
	log.Println(util.IsUrl(url))
	if !util.IsUrl(url) {
		errorJ, _ := json.Marshal(models.NewErrorUrlInvalida(url))
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
			util.ResponseBadRequest(w, errorJ)
			return
		}

		urlEncurtada, err := e.encurtarComAlias(url, aliasCustomizado)
		if err != nil {
			errorJ, _ := json.Marshal(models.NewErrorInterno())
			util.ResponseInternalError(w, errorJ)
			return
		}
		urlEncurtadaJ, _ := json.Marshal(urlEncurtada)
		util.ResponseCreate(w, urlEncurtadaJ)
		return
	}

	urlEncurtada, err := e.encurtar(url)
	if err != nil {
		errorJ, _ := json.Marshal(models.NewErrorInterno())
		util.ResponseInternalError(w, errorJ)
		return
	}
	urlEncurtadaJ, _ := json.Marshal(urlEncurtada)
	util.ResponseCreate(w, urlEncurtadaJ)
	return
}

func (e *Encurtador) encurtarComAlias(url string, alias string) (models.UrlEncurtada, error) {
	return e.encurtador.Encurtar(url, alias)
}

func (e *Encurtador) encurtar(url string) (models.UrlEncurtada, error) {
	return e.encurtador.Encurtar(url, "")
}

func (e *Encurtador) BuscarURL(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	alias := params["alias"]
	urlEncurtada, errRetorno := e.encurtador.BuscarPorAlias(alias)
	if errRetorno != nil {
		errorJ, _ := json.Marshal(errRetorno)
		util.ResponseBadRequest(w, errorJ)
		return
	}
	http.Redirect(w, r, urlEncurtada.URL, http.StatusMovedPermanently)
	return
}

func (e *Encurtador) ListarUrlsMaisAcessadas(w http.ResponseWriter, r *http.Request) {
	urls, err := e.encurtador.BuscarDezUrlsMaisAcessadas()
	if err != nil {
		errorJ, _ := json.Marshal(err)
		util.ResponseBadRequest(w, errorJ)
		return
	}
	urlsJ, _ := json.Marshal(urls)
	util.ResponseOK(w, urlsJ)
	return
}
