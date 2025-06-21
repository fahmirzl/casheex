package routers

import (
	"github.com/gin-gonic/gin"
	_ "casheex/docs" 
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)


func StartServer() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := r.Group("/api")

	AddUserRouter(api)
	AddProductRouter(api)
	AddCartRouter(api)
	AddTransactionRouter(api)

	return r
}