package routers

import (
	"casheex/handlers"
	"casheex/middlewares"

	"github.com/gin-gonic/gin"
)

func AddCartRouter(rg *gin.RouterGroup) {
	cart := rg.Group("/carts")
	cart.GET("/", middlewares.JWTAuthMiddleware(false), handlers.CartIndex)
	cart.POST("/add", middlewares.JWTAuthMiddleware(false), handlers.AddToCart)
	cart.DELETE("/remove/:id", middlewares.JWTAuthMiddleware(false), handlers.RemoveFromCart)
}