package router

import (
	"net/http"

	"github.com/DanielDanteDosSantosViana/hire.me/handlers"
	"github.com/DanielDanteDosSantosViana/hire.me/models"
	"github.com/gorilla/mux"
)

func IniciarRotas() {
	encurtador := models.NewEncurtador()
	encurtadorHandler := handlers.NewEncurtador(encurtador)
	router := mux.NewRouter()
	router.HandleFunc("/criar", encurtadorHandler.EncurtarURL).Methods(http.MethodPut)
	router.HandleFunc("/url/{alias}", encurtadorHandler.BuscarURL).Methods(http.MethodGet)
	router.HandleFunc("/list", encurtadorHandler.ListarUrlsMaisAcessadas).Methods(http.MethodGet)

	http.Handle("/", router)
}
