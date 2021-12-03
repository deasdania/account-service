package main

import (
	"account-metalit/api"
	"account-metalit/docs"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

func main() {
	if len(os.Args) > 1 && os.Args[1] == "local" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file not FOUND")
		}
	}
	r := gin.Default()

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Toko Bunga API"
	docs.SwaggerInfo.Description = "This is a Final Project of Camp"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = os.Getenv("PORT_RUN")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	api.Init(r)
}
