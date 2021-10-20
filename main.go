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
	router.POST("/topup", handlers.TopUpBalanceApi)
	router.POST("/repl", handlers.MonthlyReplenishments)


	port := os.Getenv("APP_HTTP_PORT")
	router.Run(port)
}