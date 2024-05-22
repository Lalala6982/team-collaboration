package main

import (
	// "campbe/handler"
	"campbe/gateway"
	"fmt"
	// "log"
	// "net/http"
)

func main() {
    fmt.Println("started-service")
	origin := "New York, NY"
    destination := "Washington, DC"
	gateway.GetRobotRoute(origin, destination)
	gateway.GetDroneRoute(origin, destination)
    // log.Fatal(http.ListenAndServe(":8080", handler.InitRouter()))
}