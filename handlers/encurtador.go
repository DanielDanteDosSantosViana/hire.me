package handlers

import (
	"net/http"
	"github.com/DanielDanteDosSantosViana/hire.me/db"
	"github.com/DanielDanteDosSantosViana/hire.me/model"
	"github.com/DanielDanteDosSantosViana/hire.me/util"

)

type Encurtador struct {
	session *mgo.Session
}

func NewEncurtador() *Encurtador {
	return &Encurtador{db.NewSession()}
}

func (e *Encurtador) EncurtarURL(w http.ResponseWriter, r *http.Request) {
	
	url := r.URL.Query().Get("url")
    aliasCustomizado := r.URL.Query().Get("CUSTOM_ALIAS")

   	if url == "" {
		log.Println("URL não informada")
		http.Error(w, "URL não informada ", http.StatusInternalServerError)
		return
	}

    recuperador := models.NewRecuperador()
    if aliasCustomizado != "" {
		urlRetorno := recuperador.buscarUrlPorAlias(aliasCustomizado)
		if urlRetorno!=nil{
			log.Println("Alias informado já existe")
			http.Error(w, "URL não informada ", http.StatusInternalServerError)
			return
		}	
	}

	err, alias := util.Encurtar(url)
	if err != nil {
		log.Println("Error ao encurtar a URL")
		http.Error(w, "Error ao encurtar a URL", http.StatusInternalServerError)
		return
	}
	encurtador := models.NewEncurtardor(alias,url);
	err, encurtador := encurtador.Criar()
	if err != nil {
		log.Println("Error ao salvar a URL")
		http.Error(w, "Error ao encurtar a URL", http.StatusInternalServerError)
		return
	}	
	//recuperador.buscarPorAlias(customizado)

}