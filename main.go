package main 

import (
	"github.com/joho/godotenv"
	"os"
	"log"
	"github.com/gin-gonic/gin"
)


func main () {

	err := godotenv.Load()
	if err != nil {
		log.Print(err)
	}

	router := gin.Default()
	
	port := os.Getenv("APP_HTTP_PORT")
	router.Run(port)
}