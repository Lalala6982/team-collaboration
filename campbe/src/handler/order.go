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

type ShippingInfoRequest struct {
	FromAddress string `json:"from_address"`
	FromZipCode string `json:"from_zip_code"`
	FromCity    string `json:"from_city"`
	ToAddress   string `json:"to_address"`
	ToZipCode   string `json:"to_zip_code"`
	ToCity      string `json:"to_city"`
	TotalWeight int    `json:"total_weight"`
}

func getShippingOptionsHandler(w http.ResponseWriter, r *http.Request) {
	var req ShippingInfoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Fetch dispatching options
	fromFields := fmt.Sprintf("%s, %s, %s", req.FromAddress, req.FromCity, req.FromZipCode)
	toFields := fmt.Sprintf("%s, %s, %s", req.ToAddress, req.ToCity, req.ToZipCode)
	options, optionsID, err := service.GetDispatchingOptions(fromFields, toFields)
	if err != nil {
		http.Error(w, "Failed to get dispatching options", http.StatusInternalServerError)
		return
	}

	// Return the shipping options to the client
	response := map[string]interface{}{
		"options":    options,
		"options_id": optionsID,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

type CreateOrderRequest struct {
	Shipper        string `json:"shipper"`
	FromAddress    string `json:"from_address"`
	FromZipCode    string `json:"from_zip_code"`
	FromCity       string `json:"from_city"`
	FromCounty     string `json:"from_county"`
	FromPhone      string `json:"from_phone"`
	FromEmail      string `json:"from_email"`
	Consignee      string `json:"consignee"`
	ToAddress      string `json:"to_address"`
	ToZipCode      string `json:"to_zip_code"`
	ToCity         string `json:"to_city"`	
	ToCounty       string `json:"to_county"`
	ToPhone        string `json:"to_phone"`
	ToEmail        string `json:"to_email"`
	TotalWeight    int    `json:"total_weight"`
	SelectedOption string `json:"selected_option"`
	OptionsID      string `json:"options_id"`
	Status         string `json:"status"`
	OrderTime      string `json:"order_time"`
	ProductID      string `json:"product_id"`
	Price          float64`json:"price"`
	PriceID        string `json:"price_id"`
	Deliver        string `json:"deliver"`
	Duration       string `json:"duration"`
	Distance       float64 `json:"distance"`
}

func createOrderHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	// Retrieve stored options using OptionsID
	optionsStore, exists := service.OptionsCache[req.OptionsID]
	if !exists || len(optionsStore.Options) == 0 {
		http.Error(w, "Invalid or expired options ID", http.StatusBadRequest)
		return
	}

	// Find the selected option based on the OptionID
	var selectedOption service.Option
	for _, option := range optionsStore.Options {
		if option.OptionID == req.SelectedOption {
			selectedOption = option
			break
		}
	}

	if selectedOption.OptionID == "" {
		http.Error(w, "Invalid selected option", http.StatusBadRequest)
		return
	}

	// Create a new order based on the selected option and provided info
	order := model.Order{
		Id:             uuid.New().String(),
		Shipper:        req.Shipper,
		FromAddress:    req.FromAddress,
		FromZipCode:    req.FromZipCode,
		FromCity:       req.FromCity,
		FromCounty:     req.FromCounty,
		FromPhone:      req.FromPhone,
		FromEmail:      req.FromEmail,
		Consignee:      req.Consignee,
		ToAddress:      req.ToAddress,
		ToZipCode:      req.ToZipCode,
		ToCity:         req.ToCity,
		ToCounty:       req.ToCounty,
		ToPhone:        req.ToPhone,
		ToEmail:        req.ToEmail,
		TotalWeight:    req.TotalWeight,
		Status:         req.Status,
		OrderTime:      req.OrderTime,
		ProductID:      req.ProductID,
		Price:          req.Price,
		PriceID:        req.PriceID,
		Deliver:        selectedOption.Transportation,
		Duration:       strconv.Itoa(selectedOption.Duration),
		Distance:       selectedOption.Distance,
	}
	if err := database.SaveToDB(order); err != nil {
		http.Error(w, "Failed to save order to database", http.StatusInternalServerError)
		return
	}

	// Respond to the client
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Order created and saved successfully"))
}

// user payment
func checkoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one checkout request")
	w.Header().Set("Content-Type", "text/plain")
	orderID := r.FormValue("orderID")
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
// func trackHandler(w http.ResponseWriter, r *http.Request) {}

func orderHistoryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received order history request")
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


// searchOrderHandler handles the request to search for an order by ID
func searchOrderHandler(w http.ResponseWriter, r *http.Request) {
	orderID := r.URL.Query().Get("order_id")
	if orderID == "" {
		http.Error(w, "order_id is required", http.StatusBadRequest)
		return
	}

	order, err := service.SearchOrderByID(orderID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}