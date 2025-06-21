package main

import (
	"casheex/configs"
	"casheex/migrations"
	"casheex/routers"
	"os"
)

// @title Casheex
// @version 1.0
// @description Casheex is a simple RESTful API for a casheer system, designed to support basic cashier operations in retail or small business environments.
// @host https://casheex-production.up.railway.app
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	configs.DBConnection()
	migrations.DBMigrate()
	defer configs.DB.Close()

	routers.StartServer().Run(":" + os.Getenv("PORT"))
}