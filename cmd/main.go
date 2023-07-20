package main

import (
	"log"

	"github.com/codingwithik/simple-blog-backend-with-go/src/config"
)

func main() {
	// Initialize Database
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}
	config.ConnectDB(&config)
	config.Migrate()
	// Initialize Router
	//router := initRouter()
	//router.Run(":8080")
}
