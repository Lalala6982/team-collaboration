package model

type Order struct {
	Id          string  `json:"id"`
	Shipper     string  `json:"shipper"`
	FromAddress string  `json:"from_address"`
	FromZipCode string  `json:"from_zip_code"`
	FromCity    string  `json:"from_city"`
	FromCounty  string  `json:"from_county"`
	FromPhone   string  `json:"from_phone"`
	FromEmail   string  `json:"from_email"`
	Consignee   string  `json:"consignee"`
	ToAddress   string  `json:"to_address"`
	ToZipCode   string  `json:"to_zip_code"`
	ToCity      string  `json:"to_city"`
	ToCounty    string  `json:"to_county"`
	ToPhone     string  `json:"to_phone"`
	ToEmail     string  `json:"to_email"`
	TotalWeight int     `json:"total_weight"`
	Status      string  `json:"status"`
	OrderTime   string  `json:"order_time"`
	ProductID   string  `json:"product_id"`
	Price       float64 `json:"price"`
	PriceID     string  `json:"price_id"`
	Deliver     string  `json:"deliver"`
	Duration    string  `json:"duration"`
	Distance    float64 `json:"distance"`
}

type User struct {
	ID       int      `jason:"id"`
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
	Id            int    `json:"id"`
	BaseId        int    `json:"base_id"`
	DeliverType   string `json:"deliver_type"`
	DeliverSpeed  int    `json:"deliver_speed"`
	DeliverStatus string `json:"deliver_status"`
}
