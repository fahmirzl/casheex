package routers

import (
	"casheex/handlers"
	"casheex/middlewares"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(rg *gin.RouterGroup) {
	user := rg.Group("/users")
	user.POST("/login", handlers.Login)
	user.GET("/", middlewares.JWTAuthMiddleware(true), handlers.UserIndex)
	user.POST("/", middlewares.JWTAuthMiddleware(true), handlers.UserStore)
	user.GET("/:id", middlewares.JWTAuthMiddleware(true), handlers.UserFind)
	user.PUT("/:id", middlewares.JWTAuthMiddleware(true), handlers.UserUpdate)
	user.DELETE("/:id", middlewares.JWTAuthMiddleware(true), handlers.UserDestroy)
}