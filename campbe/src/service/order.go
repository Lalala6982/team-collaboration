package service

import (
	"campbe/database"
	"campbe/model"
	"fmt"
	"strings"
)

func GetOrderHistory(username string) ([]model.Order, error) {
	user, err := database.GetUser(username)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	if len(user.Index) == 0 {
		return nil, fmt.Errorf("no orders found for user: %v", username)
	}
	// Create a query with the order IDs
	query := fmt.Sprintf(`SELECT id, shipper, from_address, from_zip_code, from_city, from_county, from_phone, from_email, consigee, to_address, to_zip_code, to_city, to_county, to_phone, to_email, total_weight, status, order_time, price, price_id, deliver_id 
                          FROM orders 
                          WHERE id IN ('%s')`, strings.Join(model.user.Index, "','"))
	rows, err := database.ReadFromDB(query)
	if err != nil {
		return nil, fmt.Errorf("query error: %v", err)
	}
	defer rows.Close()

	var orders []model.Order

	for rows.Next() {
		var order model.Order
		err := rows.Scan(
			&order.Id, &order.Shipper, &order.FromAddress, &order.FromZipCode, &order.FromCity, &order.FromCounty,
			&order.FromPhone, &order.FromEmail, &order.Consigee, &order.ToAddress, &order.ToZipCode, &order.ToCity,
			&order.ToCounty, &order.ToPhone, &order.ToEmail, &order.TotalWeight,
		)
		if err != nil {
			return nil, fmt.Errorf("scan error: %v", err)
		}
		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}

	return orders, nil
}

func SaveOption(option *model.Deliver) error {}


func CheckoutApp(domain string, orderID string) (string, error) {}