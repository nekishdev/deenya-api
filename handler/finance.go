package handler

import (
	"deenya-api/database"
	"deenya-api/models"
	"deenya-api/services"
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
)

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	customerID := chi.URLParam(r, "customerID")

	customer, err := database.GetCustomerByID(customerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	js, err := json.Marshal(customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	WriteAsJSON(w, js)

}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	var data models.StripeCustomer

	customerID := chi.URLParam(r, "customerID")

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// get customer
	customer, err := database.GetCustomerByID(customerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// update customer
	data.ID = customer.ID
	err = database.UpdateCustomer(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {

	customerID := chi.URLParam(r, "customerID")

	err := database.DeleteCustomer(customerID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	msg := models.JsonResultMessage{
		Message: "Success",
	}

	js, err := json.Marshal(msg)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func NewCustomer(w http.ResponseWriter, r *http.Request) {

	customer, err := services.NewStripeCustomerForClient(GetAuthID(r))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	js, err := json.Marshal(customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	WriteAsJSON(w, js)

}

func ListCustomers(w http.ResponseWriter, r *http.Request) {
}

func GetCharge(w http.ResponseWriter, r *http.Request) {

}

func UpdateCharge(w http.ResponseWriter, r *http.Request) {

}

func DeleteCharge(w http.ResponseWriter, r *http.Request) {

}

func NewCharge(w http.ResponseWriter, r *http.Request) {

}

func ListCharges(w http.ResponseWriter, r *http.Request) {
}

func GetSubscription(w http.ResponseWriter, r *http.Request) {

}

func UpdateSubscription(w http.ResponseWriter, r *http.Request) {

}

func DeleteSubscription(w http.ResponseWriter, r *http.Request) {

}

func NewSubscription(w http.ResponseWriter, r *http.Request) {

}

func ListSubscriptions(w http.ResponseWriter, r *http.Request) {
}
