package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func ConnectDB() {
	var err error
	dsn := "root:@tcp(localhost:3306)/todo_list"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Database tidak merespons:", err)
	}

	fmt.Println("Koneksi database sukses!")
}

func GetDB() *sql.DB {
	return db
}
