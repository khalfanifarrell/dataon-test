package main

import (
	"dataon-test/app"
	"dataon-test/db"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	database, err := db.CreateDatabase()
	if err != nil {
		log.Printf("Database connection failed: %s", err.Error())
		log.Fatal("Exiting...")
	}

	app := &app.App{
		Router:   mux.NewRouter().StrictSlash(true),
		Database: database,
	}

	app.SetupRouter()

	log.Fatal(http.ListenAndServe(":8080", app.Router))
}
