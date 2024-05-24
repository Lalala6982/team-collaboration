package model

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	ORDER_INDEX    = "order"
	USER_INDEX     = "user"
	STRIPE_API_KEY = ""
	GCS_BUCKET     = ""
	DB_USER        = "flagcamp"
	DB_PASSWORD    = "flagcamp"
	// DB_HOST        = "mysql-container"
	DB_HOST        = "localhost"
	DB_PORT        = "3306"
	DB_NAME        = "mydb"
	MAP_API_KEY    = "AIzaSyA6no3J1oLtfvKm8okja-D0kxcz47KzD3k"
	DRONE_VELOCITY = 100.0 // km/h
	ROBOT_CHARGE   = 0.02  // $/km
	DRONE_CHARGE   = 0.1   // $/km
)

var Client *http.Client

func ProxySet() {
	proxyURL, err := url.Parse("http://127.0.0.1:10808")
	if err != nil {
		fmt.Println("Invalid proxy URL:", err)
		os.Exit(1)
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	Client = &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	}
}
