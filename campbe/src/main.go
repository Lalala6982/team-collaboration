package main

import (
	// "campbe/handler"
	"fmt"
	"log"
	// "net/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err !=nil{
		log.Fatal("Error loading .env file")
	}
	fmt.Println("started-service")
    // log.Fatal(http.ListenAndServe(":8080", handler.InitRouter()))
}