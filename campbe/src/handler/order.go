package handler

import (
	"campbe/database"
	"campbe/model"
	"campbe/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/google/uuid"
)

// user submit order information and create a new entry
func uploadOrderHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Shipping Info Received")

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	totalWeight, err := strconv.Atoi(r.FormValue("total_weight"))
	if err != nil {
		http.Error(w, "Invalid total weight format", http.StatusBadRequest)
		return
	}
	price, err := strconv.Atoi(r.FormValue("price"))
	if err != nil {
		http.Error(w, "Invalid price format", http.StatusBadRequest)
		return
	}

	order := model.Order{
		Id:          uuid.New().String(),
		Shipper:     r.FormValue("shipper"),
		FromAddress: r.FormValue("from_address"),
		FromZipCode: r.FormValue("from_zip_code"),
		FromCity:    r.FormValue("from_city"),
		FromCounty:  r.FormValue("from_county"),
		FromPhone:   r.FormValue("from_phone"),
		FromEmail:   r.FormValue("from_email"),
		Consignee:   r.FormValue("consignee"),
		ToAddress:   r.FormValue("to_address"),
		ToZipCode:   r.FormValue("to_zip_code"),
		ToCity:      r.FormValue("to_city"),
		ToCounty:    r.FormValue("to_county"),
		ToPhone:     r.FormValue("to_phone"),
		ToEmail:     r.FormValue("to_email"),
		TotalWeight: totalWeight,
		Status:      r.FormValue("status"),
		OrderTime:   r.FormValue("order_time"),
		// ProductID:   r.FormValue("product_id"),
		Price:     price,
		PriceID:   r.FormValue("price_id"),
		DeliverID: r.FormValue("deliver_id"),
	}

	// Get recommendation
	fromFields := fmt.Sprintf("%s, %s, %s", order.FromAddress, order.FromCity, order.FromZipCode)
	toFields := fmt.Sprintf("%s, %s, %s", order.ToAddress, order.ToCity, order.ToZipCode)
	options, err := service.GetDispatchingOptions(fromFields, toFields)
	if err != nil {
		http.Error(w, "Failed to get dispatching options", http.StatusInternalServerError)
		return
	}

	// Print or handle the options
	decisionId := 0 // which option user chooses
	fmt.Println(options)

	// Save the order to the database
	deliver, err := service.SaveOption(options[decisionId], &order)
	err = database.SaveToDB(deliver)
	if err := database.SaveToDB(order); err != nil {
		http.Error(w, "Failed to save order to database", http.StatusInternalServerError)
		return
	}

	// Respond to the client
	fmt.Fprintf(w, "Order saved successfully")
}

// provide shipping options based on the order information for users
// func recommendHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Recieved Recommendation Request")

// }

// user confirm final decision and dispatching centers deal with the shipping
// func dispatchHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Received Dispatching Request")
// }

// user payment
func checkoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one checkout request")
	w.Header().Set("Content-Type", "text/plain")
	orderID := r.FormValue("orderID")
	url, err := service.CheckoutApp(r.Header.Get("Origin"), orderID)
	if err != nil {
		fmt.Println("Checkout failed.")
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(url))
	fmt.Println("Checkout process started!")
}

// track the status of an order
func trackHandler(w http.ResponseWriter, r *http.Request) {}

func orderHistoryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received Order History request")
	token := r.Context().Value("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	//1. process request: URL param -> string
	w.Header().Set("Content-Type", "application/json")
	// status := r.URL.Query().Get("status")
	// price := r.URL.Query().Get("price")
	// orderTime := r.URL.Query().Get("order_time")
	// deliverID := r.URL.Query().Get("deliver_id")

	// Fetch orders from service user
	orders, err := service.GetOrderHistory(username)
	if err != nil {
		http.Error(w, "Failed to read orders from backend", http.StatusInternalServerError)
		return
	}

	// Construct response
	js, err := json.Marshal(orders)
	if err != nil {
		http.Error(w, "Failed to parse orders into JSON format", http.StatusInternalServerError)
		return
	}

	// Set response headers and write response
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
