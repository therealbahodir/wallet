package handlers


import (
	"github.com/gin-gonic/gin"
	"github.com/therealbahodir/wallet/database"
	"log"
)


func CheckApi (ctx *gin.Context) {

	userId := ctx.Query("X-UserId")
	digest := ctx.Query("X-Digest")

	err := database.IsExisting(userId, digest)

	if err != nil {
		ctx.Writer.WriteHeader(401)
		log.Print(err)
		return
	}

	ctx.Writer.WriteHeader(200)

}