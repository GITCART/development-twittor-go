package main

import (
	"log"

	"github.com/GITCART/development-twittor-go/bd"

	"github.com/GITCART/development-twittor-go/handlers"
)

func main() {
	if bd.ChequeConnection() == 0 {
		log.Fatal("sin conexion  a la bd")
		return
	}

	handlers.Manejadores()
}
