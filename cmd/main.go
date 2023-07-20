package main

import (
	"github.com/codingwithik/simple-blog-backend-with-go/cmd/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"

	"github.com/codingwithik/simple-blog-backend-with-go/src/config"
)

func main() {
	// Initialize Database
	conf, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}
	config.ConnectDB(&conf)
	err = config.Migrate()
	if err != nil {
		log.Fatal("? Could not migrate:", err)
	}

	// Initialize Router

	docs.SwaggerInfo.BasePath = "/api/v1"

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Fatal(router.Run(":8080"))

}
