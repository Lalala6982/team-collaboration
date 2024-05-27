package service

import (
	"campbe/database"
	"campbe/gateway"
	"campbe/model"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
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
			&order.ToCounty, &order.ToPhone, &order.ToEmail, &order.TotalWeight, &order.Status, &order.OrderTime,
			&order.Price, &order.PriceID, &order.DeliverID,
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

func SaveOption(option Option, order *model.Order) (model.Deliver, error) {
	deliver := model.Deliver{
		Id:              uuid.New().String(),
		BaseId:          option.baseId,
		DeliverType:     option.transportation,
		DeliverDuration: option.duration,
		DeliverStatus:   "Pending",
	}
	query := "INSERT INTO delivers (base_id, deliver_type, deliver_duration, deliver_status) VALUES (?, ?, ?, ?)"
	if _, err := database.Dbsql.Exec(query, deliver.BaseId, deliver.DeliverType, deliver.DeliverDuration, deliver.DeliverStatus); err != nil {
		panic(err)
	}

	order.DeliverID = deliver.Id
	order.Price = int(option.price)
	_, priceID, err := gateway.CreateOrderWithPrice(order.Id, int64(order.Price*100))
	if err != nil {
		panic(err)
	}
	// order.ProductID = productID
	order.PriceID = priceID

	return deliver, nil
}

// func CheckoutApp(domain string, orderID string) (string, error) {
func CheckoutApp(domain string, order *model.Order) (string, error) {
	// order, err := ReadFromDB(orderID)
	// if err != nil {
	// 	return "", err
	// }
	if order == nil {
		return "", errors.New("unable to find order in database")
	}
	return gateway.CreateCheckoutSession(domain, order.PriceID)
}
