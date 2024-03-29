package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
)

type Database struct {
	db *sql.DB
}

func NewDatabase() *sql.DB {
	// Inisialisasi koneksi database MySQL
	connString := "root:password123#@tcp(localhost:3306)/db_local"
	db, err := sql.Open("mysql", connString)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	return db
}

func (db *Database) Close() {
	// Tutup koneksi database
	err := db.db.Close()
	if err != nil {
		log.Fatalf("Failed to close database connection: %v", err)
	}
}
