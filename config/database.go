package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Konfigurasi Database Golang CRUD Web

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "root:Kurshiel123A!@/go_products")
	if err != nil {
		panic(err)
	}

	log.Println("Database Connected")
	DB = db
}
