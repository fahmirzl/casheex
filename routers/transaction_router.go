package routers

import (
	"casheex/handlers"
	"casheex/middlewares"

	"github.com/gin-gonic/gin"
)

func AddTransactionRouter(rg *gin.RouterGroup) {
	transaction := rg.Group("transactions")
	transaction.POST("/", middlewares.JWTAuthMiddleware(false), handlers.TransactionStore)
}