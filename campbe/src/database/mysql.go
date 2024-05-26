package database
package database

import (
	"campbe/model"
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"
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
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", constants.DB_USER, constants.DB_PASSWORD, constants.DB_HOST, constants.DB_PORT, constants.DB_NAME)
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
    v := reflect.ValueOf(i)
    t := reflect.TypeOf(i)
    
    if t.Kind() != reflect.Struct {
        return fmt.Errorf("SaveToDB: expected a struct, got %s", t.Kind())
    }
    
    tableName := strings.ToLower(t.Name())
    
    var columns []string
    var placeholders []string
    var values []interface{}
    
    for j := 0; j < t.NumField(); j++ {
        field := t.Field(j)
        columns = append(columns, field.Name)
        placeholders = append(placeholders, "?")
        values = append(values, v.Field(j).Interface())
    }
    
    query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, strings.Join(columns, ", "), strings.Join(placeholders, ", "))
    
    _, err := Dbsql.Exec(query, values...)
    if err != nil {
        return fmt.Errorf("SaveToDB: failed to execute query: %v", err)
    }
    
    fmt.Println("Saved to database successfully!")
    return nil
}

