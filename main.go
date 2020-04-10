package main

import (
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/go-redis/redis/v7"
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
	os.Setenv("PORT", "8002")

	// Setup Redis database
	addr := getEnv("DB_CONNECTION", "localhost:16379")
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	err := client.Set("key", "value", 0).Err()
	failOnError(err, "Error conecting to redis")
	defer client.Close()

	// Setup router
	r := controller.SetupRouter(client)
	err = r.Run() // listen and serve on 0.0.0.0:8002
	failOnError(err, "Failed starting server")
}
