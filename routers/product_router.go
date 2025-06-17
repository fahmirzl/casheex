package routers

import (
	"casheex/handlers"
	"casheex/middlewares"

	"github.com/gin-gonic/gin"
)

func AddProductRouter(rg *gin.RouterGroup) {
	product := rg.Group("/products")
	product.GET("/", middlewares.JWTAuthMiddleware(false), handlers.ProductIndex)
	product.POST("/", middlewares.JWTAuthMiddleware(true), handlers.ProductStore)
	product.GET("/:id", middlewares.JWTAuthMiddleware(false), handlers.ProductFind)
	product.PUT("/:id", middlewares.JWTAuthMiddleware(true), handlers.ProductUpdate)
	product.DELETE("/:id", middlewares.JWTAuthMiddleware(true), handlers.ProductDestroy)
}