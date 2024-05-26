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
	MAP_API_KEY    = "AIzaSyA6no3J1oLtfvKm8okja-D0kxcz47KzD3k"
	GCS_BUCKET     = ""
	DRONE_VELOCITY = 100.0 // km/h
	ROBOT_CHARGE   = 0.02  // $/km
	DRONE_CHARGE   = 0.1   // $/km
	DB_USER        = "flagcamp"
	DB_PASSWORD    = "flagcamp"
	// DB_HOST        = "mysql-container"
	DB_HOST        = "localhost"
	DB_PORT        = "3306"
	DB_NAME        = "mydb"
)
