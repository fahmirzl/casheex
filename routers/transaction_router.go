package routers

import (
	"casheex/handlers"
	"casheex/middlewares"

	"github.com/gin-gonic/gin"
)

func AddTransactionRouter(rg *gin.RouterGroup) {
	transaction := rg.Group("transactions")
	transaction.POST("/", middlewares.JWTAuthMiddleware(false), handlers.TransactionStore)
	transaction.GET("/", middlewares.JWTAuthMiddleware(false), handlers.TransactionList)
	transaction.GET("/all", middlewares.JWTAuthMiddleware(true), handlers.TransactionAll)
	transaction.GET("/profit", middlewares.JWTAuthMiddleware(true), handlers.Profit)
}