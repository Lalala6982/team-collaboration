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
	Consigee    string `json:"consignee"`
	ToAddress   string `json:"to_address"`
	ToZipCode   string `json:"to_zip_code"`
	ToCity      string `json:"to_city"`
	ToCounty    string `json:"to_county"`
	ToPhone     string `json:"to_phone"`
	ToEmail     string `json:"to_email"`
	TotalWeight string `json:"total_weight"`
	Status      string `json:"status"`
	OrderTime   string `json:"order_time"`
	Price       int    `json:"price"`
	//	ProductID   string `json:"product_id"`
	PriceID   int `json:"price_id"`
	DeliverID int `json:"deliver_id"`
}

type User struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Base struct {
	Id          int    `json:"id"`
	BaseAddress string `json:"base_address"`
	NumOfRobots int    `json:"num_of_robots"`
	NumOfDrones int    `json:"num_of_drones"`
}

type Deliver struct {
	Id            int    `json:"id"`
	BaseId        int    `json:"base_id"`
	DeliverType   string `json:"deliver_type"`
	DeliverSpeed  int    `json:"deliver_speed"`
	DeliverStatus string `json:"deliver_status"`
}
