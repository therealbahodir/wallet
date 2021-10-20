package main 

import (
	"github.com/joho/godotenv"
	"os"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/therealbahodir/wallet/database"
)


func main () {

	_, err := database.DBConnection()
	if err != nil {
		log.Print(err)
	}

	err = godotenv.Load()
	if err != nil {
		log.Print(err)
	}

	router := gin.Default()

	port := os.Getenv("APP_HTTP_PORT")
	router.Run(port)
}