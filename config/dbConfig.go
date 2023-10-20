package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DBInstance *sql.DB

func DBconfig() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("error: ", err)
	}

	config := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "recordings",
		AllowNativePasswords: true,
	}

	DB, err := sql.Open("mysql", config.FormatDSN())

	DBInstance = DB

	if err != nil {
		log.Fatal("db connection err:   ", err)
	}

	log.Println("DB connected...")

}
