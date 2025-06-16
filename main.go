package main

import (
	"casheex/configs"
	"casheex/migrations"
)

func main() {
	configs.DBConnection()
	migrations.DBMigrate()
	defer configs.DB.Close()
}