package handler

import ()

// user submit order information and create a new entry
func uploadOrderHandler(w http.ResponseWriter, r *http.Request) {}
// provide shipping options based on the order information for users
func recommendHandler(w http.ResponseWriter, r *http.Request) {}
// user confirm final decision and dispatching centers deal with the shipping
func dispatchHandler(w http.ResponseWriter, r *http.Request) {}
// user payment
func checkoutHandler(w http.ResponseWriter, r *http.Request) {}
// track the status of an order
func trackHandler(w http.ResponseWriter, r *http.Request) {}
