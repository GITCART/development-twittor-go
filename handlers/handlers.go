package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/GITCART/development-twittor-go/middlew"
	"github.com/GITCART/development-twittor-go/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores seteo mi puerto, el Handler y pongo a escuchar al servidor*/
func Manejadores() {
	router := mux.NewRouter()

	//RUTAS
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8084"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
