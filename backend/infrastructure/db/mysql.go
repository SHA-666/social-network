package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDB() (*sql.DB, error) {
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MY_SQLPW")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DB")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func Close(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Printf("Erreur lors de la fermeture de la DB : %v", err)
	}
}
