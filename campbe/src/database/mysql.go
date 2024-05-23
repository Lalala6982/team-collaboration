package database

import (
	"campbe/model"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDB() (*sql.DB, error) {
	// Replace the following with your actual database credentials

	user := os.Getenv(model.DB_USER)
	password := os.Getenv(model.DB_PASSWORD)
	host := os.Getenv(model.DB_HOST)
	port := os.Getenv(model.DB_HOST)
	database := os.Getenv(model.DB_NAME)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database!")
	return db, nil
}
