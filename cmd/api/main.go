package main

import (
	"flag"
	"fmt"
	"go-restapi/inernal/repository"
	"go-restapi/inernal/repository/dbrepo"
	"log"
	"net/http"
	"time"
)

const port = 8080

type application struct {
	DSN          string
	Domain       string
	DB           repository.DatabaseRepo
	auth         Auth
	JWTSecret    string
	JWTIssuer    string
	JWTAudience  string
	CookieDomain string
}

func main() {
	//set application config
	var app application
	//read from cmdline
	flag.StringVar(&app.DSN,
		"dsn",
		"host=localhost port=5432 user=postgres password=postgres dbname=movies sslmode=disable timezone=UTC connect_timeout=5",
		"Postgres connection string")
	flag.StringVar(&app.JWTSecret, "jwt-secret", "itssecret", "signing secret")
	flag.StringVar(&app.JWTIssuer, "jwt-issuer", "example.com", "signing issuer")
	flag.StringVar(&app.JWTAudience, "jwt-audience", "example.com", "signing audience")
	flag.StringVar(&app.CookieDomain, "cookie-domain", "localhost", "cookie domain")
	flag.StringVar(&app.Domain, "domain", "example.com", "domain")
	flag.Parse()
	//connect to db
	conn, err := app.connectToDb()
	if err != nil {
		log.Fatal(err)
	}
	app.DB = &dbrepo.PostgresDBRepo{Db: conn}
	defer app.DB.Connection().Close()
	app.auth = Auth{
		Issuer:        app.JWTIssuer,
		Audience:      app.JWTAudience,
		Secret:        app.JWTSecret,
		TokenExpiry:   time.Minute * 15,
		RefreshExpiry: time.Hour * 24,
		CookiePath:    "/",
		CookieDomain:  app.CookieDomain,
		CookieName:    "__Host-refresh_token",
	}
	//start the application server
	log.Println("starting application on port", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
