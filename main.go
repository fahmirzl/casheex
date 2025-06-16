package main

import (
	"casheex/configs"
	"casheex/migrations"
	"casheex/routers"
)

func main() {
	configs.DBConnection()
	migrations.DBMigrate()
	defer configs.DB.Close()

	routers.StartServer().Run(":8080")
}