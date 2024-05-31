package database

import (
	"campbe/constants"
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

func ReadFromDB(query string,args ...interface{}) (*sql.Rows, error) {
	// Execute the query, which requires default SQL syntax
	fmt.Println("Executing query:", query)
	results, err := Dbsql.Query(query, args...)
	if err != nil {
		// log.Fatal(err)
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	if err := results.Err(); err != nil {
		// log.Fatal(err)
		return nil, fmt.Errorf("error in result set: %v", err)
	}
	fmt.Println("Database Read successful!")
	return results, nil
}

func SaveToDBs(i interface{}) error {
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


func SaveOrderToDB(order model.Order) error {

    query := `INSERT INTO orders (id, shipper, from_address, from_zip_code, from_city, from_county, from_phone, from_email, 
 	   consignee, to_address, to_zip_code, to_city, to_county, to_phone, to_email, total_weight, user_name, status, order_time, 
 	   product_id, price, price_id, deliver, duration, distance) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	// Log the query and parameters for debugging purposes
	fmt.Println("Executing query:", query)
	fmt.Printf("With parameters: %+v\n", order)

    _, err = Dbsql.Exec(query, order.Id, order.Shipper, order.FromAddress, order.FromZipCode, order.FromCity, order.FromCounty, order.FromPhone,
 	   order.FromEmail, order.Consignee, order.ToAddress, order.ToZipCode, order.ToCity, order.ToCounty, order.ToPhone, order.ToEmail,
 	   order.TotalWeight, order.UserName, order.Status, order.OrderTime, order.ProductID, order.Price, order.PriceID, order.Deliver,
 	   order.Duration, order.Distance)
    if err != nil {
 	   return fmt.Errorf("failed to insert order: %v", err)
    }

    return nil
}

