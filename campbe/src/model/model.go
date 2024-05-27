package model

type Order struct {
	Id          string `json:"id"`
	Shipper     string `json:"shipper"`
	FromAddress string `json:"from_address"`
	FromZipCode string `json:"from_zip_code"`
	FromCity    string `json:"from_city"`
	FromCounty  string `json:"from_county"`
	FromPhone   string `json:"from_phone"`
	FromEmail   string `json:"from_email"`
	Consignee   string `json:"consignee"`
	ToAddress   string `json:"to_address"`
	ToZipCode   string `json:"to_zip_code"`
	ToCity      string `json:"to_city"`
	ToCounty    string `json:"to_county"`
	ToPhone     string `json:"to_phone"`
	ToEmail     string `json:"to_email"`
	TotalWeight int    `json:"total_weight"`
	Status      string `json:"status"`
	OrderTime   string `json:"order_time"`
	Price       int    `json:"price"`
	PriceID     string `json:"price_id"`
	DeliverID   string `json:"deliver_id"`
}

type User struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	Index    []string `json:"id"`
}

type Base struct {
	Id          int    `json:"id"`
	BaseAddress string `json:"base_address"`
	NumOfRobots int    `json:"num_of_robots"`
	NumOfDrones int    `json:"num_of_drones"`
}

type Deliver struct {
	Id              string `json:"id"`
	BaseId          int    `json:"base_id"`
	DeliverType     string `json:"deliver_type"`
	DeliverDuration int    `json:"deliver_duration"`
	DeliverStatus   string `json:"deliver_status"`
}
