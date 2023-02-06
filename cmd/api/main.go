package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string
}

func main() {
	//set application config
	var app application
	//read from cmdline

	//connect to db
	app.Domain = "localhost"
	//start the application server
	log.Println("starting application on port", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
