package route

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
	router.HandleFunc("/url", encurtadorHandler.EncurtarURL).Methods("PUT")
	router.HandleFunc("/url/{id}", encurtadorHandler.BuscarURL).Methods("GET")
	http.Handle("/", router)
}
