package model

type Order struct {
	Id          string `json:"id"`
	User        string `json:"user"`
	FromAddress string `json:"from_address"`
	ToAddress   string `json:"to_address"`
	Status      string `json:"status"`
	OrderTime   string `json:"order_time"`
	Price       int    `json:"price"`
	ProductID   string `json:"product_id"`
	PriceID     string `json:"price_id"`
}

type User struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	UserAddress string `json:"user_address"`
}

type Dispatcher struct {
	Id string `json:"id"`
	Address string `json:"address"`
	NumOfRobots int `json:"num_of_robots"`
	NumOfDrones int `json:"num_of_drones"`
}