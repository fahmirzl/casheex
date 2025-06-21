package main

import (
	"casheex/configs"
	"casheex/migrations"
	"casheex/routers"
)

// @title Casheex
// @version 1.0
// @description Casheex is a simple RESTful API for a casheer system, designed to support basic cashier operations in retail or small business environments.
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	configs.DBConnection()
	migrations.DBMigrate()
	defer configs.DB.Close()

	routers.StartServer().Run(":8080")
}