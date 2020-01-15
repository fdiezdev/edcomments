package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/urfave/negroni"

	"github.com/fdiezdev/edcomments/commons"
	"github.com/fdiezdev/edcomments/migrations"
	"github.com/fdiezdev/edcomments/routes"
)

func main() {
	var migrate string

	flag.StringVar(&migrate, "migrate", "no", "Generates the DB migration")
	flag.IntVar(&commons.Port, "port", 8080, "Puerto en el que va a correr el servidor web")
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
		Addr:    fmt.Sprintf(":%d", commons.Port),
		Handler: n,
	}

	log.Printf("Servidor iniciado en http://localhost:%d", commons.Port)

	log.Println(server.ListenAndServe())

	log.Println("End")
}
