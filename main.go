package main

import (
	"log"

	"github.com/gitcart/development-twittor-go/bd"
	"github.com/gitcart/development-twittor-go/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("sin conexion  a la bd")
		return
	}

	handlers.Manejadores()
}
