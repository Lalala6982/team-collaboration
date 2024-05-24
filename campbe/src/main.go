package main

import (
	// "campbe/handler"
	"campbe/database"
	"fmt"

	// "net/http"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("started-service")
	database.InitMysql()
	// log.Fatal(http.ListenAndServe(":8080", handler.InitRouter()))
}
