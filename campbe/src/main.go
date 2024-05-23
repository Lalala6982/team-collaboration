package main

import (
	// "campbe/handler"
	"campbe/gateway"
	"fmt"
	"log"
	// "net/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
    fmt.Println("started-service")
	origin := "New York, NY"
    destination := "Washington, DC"
	gateway.GetRobotRoute(origin, destination)
	gateway.GetDroneRoute(origin, destination)
    // log.Fatal(http.ListenAndServe(":8080", handler.InitRouter()))
}