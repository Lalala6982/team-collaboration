package handler

import (
	"net/http"
	"encoding/json"
)

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