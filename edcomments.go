package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/urfave/negroni"

	"github.com/fdiezdev/edcomments/migrations"
	"github.com/fdiezdev/edcomments/routes"
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

	// Iniciando enrutador
	router := routes.InitRoutes()

	// Iniciando middlewares
	n := negroni.Classic()
	n.UseHandler(router)

	// Iniciando servidor
	server := &http.Server{
		Addr:    ":8080",
		Handler: n,
	}

	log.Println("Servidor iniciado en http://localhost:8080")

	log.Println(server.ListenAndServe())

	log.Println("End")
}
