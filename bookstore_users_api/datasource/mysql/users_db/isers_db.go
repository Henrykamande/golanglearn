package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client         *sql.DB
	userName       = os.Getenv("mysql_users_username")
	host           = os.Getenv("mysql_users_host")
	password       = os.Getenv("mysql_users_password")
	dabatabaseName = os.Getenv("mysql_users_dbName")
)

func init() {
	fmt.Println(userName)
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		"root", "", "127.0.0.1", "users_db")
	var err error
	Client, err = sql.Open("mysql", datasourceName)

	if err != nil {
		panic(err.Error())

	}
	if err = Client.Ping(); err != nil {
		panic(err.Error())
	}
	log.Println("Database Successfuly Connected")
}
