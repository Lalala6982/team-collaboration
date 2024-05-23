package handler

import (
	"campbe/database"
	"campbe/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

// by Ying
// user submit order information and create a new entry
func uploadOrderHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Shipping Info Received")

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
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
		Consigee:    r.FormValue("consigee"),
		ToAddress:   r.FormValue("to_address"),
		ToZipCode:   r.FormValue("to_zip_code"),
		ToCity:      r.FormValue("to_city"),
		ToCounty:    r.FormValue("to_county"),
		ToPhone:     r.FormValue("to_phone"),
		ToEmail:     r.FormValue("to_email"),
		TotalWeight: r.FormValue("total_weight"),
	}

	db, err := database.OpenDB()
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Datbase connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare(`INSERT INTO orders (id, shipper, from_address, from_zip_code, from-city, from_country, from _phone, from_email, consigee, to_address, to _zip_cod, to_city, to_county, to_phone, to_email, total_weight) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Database preparation error", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(order.Id, order.Shipper, order.FromAddress, order.FromZipCode, order.FromCity, order.FromCounty, order.FromPhone, order.FromEmail, order.Consigee, order.ToAddress, order.ToZipCode, order.ToCity, order.ToCounty, order.ToPhone, order.ToEmail, order.TotalWeight)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Database execution error", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Order has been sucessfully uploaded with ID: %s", order.Id)
}

// by Lingyun
// user submit order information and create a new entry
func uploadOrderHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var order Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Save the order to the database
	err = saveOrder(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
}

// provide shipping options based on the order information for users
func recommendHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the order information from the request
	orderID := r.URL.Query().Get("order_id")

	// Get the recommended shipping options for the order
	options, err := getShippingOptions(orderID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the shipping options as JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(options)
	fmt.Println("Recommendation Generated")
}

// user confirm final decision and dispatching centers deal with the shipping
func dispatchHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the order information from the request
	orderID := r.URL.Query().Get("order_id")

	// Dispatch the order to the appropriate shipping center
	err := dispatchOrder(orderID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
}

// user payment
func checkoutHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the order information from the request
	orderID := r.URL.Query().Get("order_id")

	// Process the payment for the order
	err := processPayment(orderID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
}

// track the status of an order
func trackHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the order information from the request
	orderID := r.URL.Query().Get("order_id")

	// Get the status of the order
	status, err := getOrderStatus(orderID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the order status as JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}
