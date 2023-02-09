package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	DSN    string
	Domain string
	DB     *sql.DB
}

func main() {
	//set application config
	var app application
	//read from cmdline
	flag.StringVar(&app.DSN,
		"dsn",
		"host=localhost port=5432 user=postgres password=postgres dbname=movies sslmode=disable timezone=UTC connect_timeout=5",
		"Postgres connection string")
	flag.Parse()
	//connect to db
	conn, err := app.connectToDb()
	if err != nil {
		log.Fatal(err)
	}
	app.DB = conn
	defer app.DB.Close()
	app.Domain = "localhost"
	//start the application server
	log.Println("starting application on port", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
