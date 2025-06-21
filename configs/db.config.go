package configs

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB  *sql.DB
	err error
)

const (
	DB_USERNAME = "root"
	DB_PASSWORD = ""
	DB_HOST     = "127.0.0.1"
	DB_PORT     = "3306"
	DB_DATABASE = "casheex"
)

func DBConnection() {
	LoadENV()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_DATABASE)
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to DB!")
}
