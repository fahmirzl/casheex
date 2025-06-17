package routers

import "github.com/gin-gonic/gin"

func StartServer() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")

	AddUserRouter(api)
	AddProductRouter(api)

	return r
}