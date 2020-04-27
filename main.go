package main

import (
	"log"
	"os"

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

// @host 159.203.183.166:8002
// @BasePath /

func main() {
	log.Println("Starting web server...")
	os.Setenv("PORT", "8010")

	// docker run --name tapcolors-html -v /root/tapcolors/html/dist:/usr/share/nginx/html:ro -p 8011:80 -d nginx

	// Setup router
	r := controller.SetupRouter()
	err := r.Run() // listen and serve on 0.0.0.0:8002
	failOnError(err, "Failed starting server")
}
