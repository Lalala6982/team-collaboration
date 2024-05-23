package handler

import "net/http"

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

	db, err := mysql.OpenDB()
	if err != nil{
		log.Fatal(err)
		http.Error(w, "Datbase connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare(`INSERT INTO orders (id, shipper, from_address, from_zip_code, from-city, from_country, from _phone, from_email, consigee, to_address, to _zip_cod, to_city, to_county, to_phone, to_email, total_weight) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil{
		log.Fatal(err)
		http.Error(w, "Database preparation error", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err =stmt.Exec(order.Id, order.Shipper, order.FromAddress, order.FromZipCode, order.FromCity, order.FromCounty, order.FromPhone, order.FromEmail, order.Consigee, order.ToAddress, order.ToZipCode, order.ToCity, order.ToCounty, order.ToPhone, order.ToEmail, order.TotalWeight)
	if err !=nil{
		log.Fatal(err)
		http.Error(w, "Database execution error", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Order has been sucessfully uploaded with ID: %s", order.Id)
}

// provide shipping options based on the order information for users
func recommendHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recommendation Generated")
	

}

// user confirm final decision and dispatching centers deal with the shipping
func dispatchHandler(w http.ResponseWriter, r *http.Request) {}

// user payment
func checkoutHandler(w http.ResponseWriter, r *http.Request) {}

// track the status of an order
func trackHandler(w http.ResponseWriter, r *http.Request) {}
