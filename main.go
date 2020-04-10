package main

import (
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/hernanrocha/minesweeper/controller"
	_ "github.com/hernanrocha/minesweeper/docs"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// @title Swagger Minesweeper API
// @version 1.0
// @description API for minesweeper

// @contact.name Hernan Rocha
// @contact.email hernanrocha93(at)gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host http://localhost:8001
// @BasePath /

func main() {
	log.Println("Starting web server...")
	os.Setenv("PORT", "8002")

	// Setup Postgres database
	/*
		dbconn := getEnv("DB_CONNECTION", "host=localhost port=15432 user=postgres password=postgres dbname=finchat sslmode=disable")
		db, err := gorm.Open("postgres", dbconn)
		failOnError(err, "Error conecting to database")
		defer db.Close()

		// Run migration
		models.Setup(db)
	*/

	// Setup router
	r := controller.SetupRouter()
	err := r.Run() // listen and serve on 0.0.0.0:8001
	failOnError(err, "Failed starting server")
}
