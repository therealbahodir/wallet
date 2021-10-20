package main 

import (
	"github.com/joho/godotenv"
	"os"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/therealbahodir/wallet/handlers"
)


func main () {

	err := godotenv.Load()
	if err != nil {
		log.Print(err)
	}

	router := gin.Default()

	router.POST("/check", handlers.CheckApi)


	port := os.Getenv("APP_HTTP_PORT")
	router.Run(port)
}