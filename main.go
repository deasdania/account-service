package main

import (
	"account-metalit/api"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// func main() {
// 	db := config.InitDb()
// 	defer db.Close()
// }

func main() {
	if len(os.Args) > 1 && os.Args[1] == "local" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file not FOUND")
		}
	}
	r := gin.Default()

	api.Init(r)
}
