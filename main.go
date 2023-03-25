package main

import (
	"log"
	"mercadofoto/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// CREA LAS RUTAS

	mux := mux.NewRouter()
	fs := http.FileServer(http.Dir("./public/"))
	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	mux.HandleFunc("/", handler.HomePage)
	mux.HandleFunc("/API_v1.0/", handler.Landing).Methods("GET")
	mux.HandleFunc("/API_v1.0/test/", handler.GetTest).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", mux))

}
