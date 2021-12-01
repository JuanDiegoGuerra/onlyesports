package main

import (
	"log"

	"github.com/JuanDiegoGuerra/onlyesports/bd"
	"github.com/JuanDiegoGuerra/onlyesports/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
