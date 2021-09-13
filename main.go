package main

import (
	"log"

	"github.com/iurzaiz/similTwitter/bd"
	"github.com/iurzaiz/similTwitter/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Siin conexion a la base de datos")
		return
	}
	handlers.Manejadores()
}
