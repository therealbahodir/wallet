package handlers


import (
	"github.com/gin-gonic/gin"
	"github.com/therealbahodir/wallet/database"
	"log"
)

type TopUp struct {
	Amount float64  `json:"amount"`
}


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


func TopUpBalanceApi (ctx *gin.Context) {

	var user TopUp
	ctx.BindJSON(&user)

	userId := ctx.Query("X-UserId")
	digest := ctx.Query("X-Digest")

	err := database.TopUpBalance(userId, digest, user.Amount)
	if err != nil {
		log.Print(err)
		ctx.JSON(400, gin.H{
			"message" : err.Error(),
		},)
		return
	}

		ctx.JSON(202, gin.H{
			"message" : "Replenishment was successfully completed",
		},)

}


func MonthlyReplenishments (ctx *gin.Context) {

	userId := ctx.Query("X-UserId")
	digest := ctx.Query("X-Digest")

	err := database.IsExisting(userId, digest)

	if err != nil {
		ctx.Writer.WriteHeader(401)
		log.Print(err)
		return
	}

	count, sum := database.ReplenishmentsInfo(userId)
	ctx.JSON(200, gin.H{
		"count" : count,
		"sum"	: sum,
	},)

}