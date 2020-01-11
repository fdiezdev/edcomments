package main

import (
	"flag"
	"log"

	"github.com/fdiezdev/edcomments/migrations"
)

func main() {
	var migrate string

	flag.StringVar(&migrate, "migrate", "no", "Generates the DB migration")
	flag.Parse()

	if migrate == "yes" {

		log.Println("Comenzó la migración, se crearán las tablas en la base de datos")
		migrations.Migrate()
		log.Println("¡Migración completada!")

	}
}
