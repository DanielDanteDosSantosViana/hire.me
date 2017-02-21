package main

import (
"net/http"
"log"
"github.com/gorilla/mux"
)
func main(){
	port := flag.String("port", "", "porta não foi definida para o serviço.")
	flag.Parse()

	if *port == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	
	db, err := sqlite.NewConn()
	if err != nil {
		log.Fatal(err)
	}
	if err := sqlite.CreateTables(db, sqlite.QuerysToMigrate); err != nil {
		log.Fatal(err.Error())
	}

	router := mux.NewRouter()
	recuperadorHandler := handlers.NewRecuperador()
	encurtadorHandler := handlers.NewEncurtador()

	router.HandleFunc("/url", encurtadorHandler.EncurtarURL).Methods("PUT")
  	router.HandleFunc("/url/{id}", recuperadorHandler.BuscarURL).Methods("GET")

	log.Println("Iniciou o serviço na porta ", port)
	http.ListenAndServe(port, nil)
}