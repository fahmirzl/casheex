package configs

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB  *sql.DB
	err error
)

func DBConnection() {
	LoadENV()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", os.Getenv("MYSQLUSER"), os.Getenv("MYSQLPASSWORD"), os.Getenv("MYSQLHOST"), os.Getenv("MYSQLPORT"), os.Getenv("MYSQLDATABASE"))
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
