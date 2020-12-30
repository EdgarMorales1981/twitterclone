package main

import (
	"log"

	"github.com/EdgarMorales1981/twitterclone/bd"
	"github.com/EdgarMorales1981/twitterclone/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
