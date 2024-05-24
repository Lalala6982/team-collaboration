package database

import (
	"campbe/model"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Dbsql *sql.DB
	err   error
)

func InitMysql() {
	// Replace the following with your actual database credentials
	// user := os.Getenv(model.DB_USER)
	// password := os.Getenv(model.DB_PASSWORD)
	// host := os.Getenv(model.DB_HOST)
	// port := os.Getenv(model.DB_PORT)
	// database := os.Getenv(model.DB_NAME)
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)

	// Open and connect to database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", model.DB_USER, model.DB_PASSWORD, model.DB_HOST, model.DB_PORT, model.DB_NAME)
	Dbsql, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	// defer Dbsql.Close()

	// Verify the connection
	err = Dbsql.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dsn, "Connected to the database!")

	// read SQL content
	sqlScript, err := os.ReadFile("database-init.sql")
	if err != nil {
		log.Fatal(err)
	}
	// Split the SQL script into individual statements
	sqlStatements := strings.Split(string(sqlScript), ";")
	// Execute each SQL statement
	for _, statement := range sqlStatements {
		// Skip empty statements
		if strings.TrimSpace(statement) == "" {
			continue
		}
		// Execute the SQL statement
		_, err := Dbsql.Exec(statement)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Database initialization successful!")
}

func ReadFromDB(query string) (*sql.Rows, error) {
	// Execute the query, which requires default SQL syntax
	results, err := Dbsql.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	if err := results.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database Read successful!")
	return results, nil
}

func SaveToDB(i interface{}) error {
	// Prepare SQL statement
	// query := "INSERT INTO tables () VALUES ()"
	// Execute the SQL statement
	// _, err := Dbsql.Exec(query, i.ID, i.Name, i.Email)
	// if err != nil {
	// 	return err
	// }
	fmt.Println("Saved to database successfully!")
	return nil
}
