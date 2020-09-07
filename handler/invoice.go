package handler

import (
	"deenya-api/database"
	"deenya-api/models"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/clarketm/json"
)

func MyInvoices(w http.ResponseWriter, r *http.Request) {
	// err := GetAuthID(r)

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// }

	mid := GetAuthID(r)
	mtype := GetAuthType(r)

	data, err := database.MyInvoices(mid, mtype)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	js, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	WriteAsJSON(w, js)

}

func GetInvoice(w http.ResponseWriter, r *http.Request) {
	q := chi.URLParam(r, "invoiceID")

	id, err := strconv.ParseInt(q, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mid := GetAuthID(r)

	mtype := GetAuthType(r)

	data, err := database.GetInvoice(id, mid, mtype)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if *data.IsBooking {
		//select * from booking where invoice_id = $1
	}

	if *data.IsOrder {
		//select * from order where invoice_id = $1
	}

	js, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	WriteAsJSON(w, js)
}

func UpdateInvoice(w http.ResponseWriter, r *http.Request) {
	var data models.Invoice
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := database.UpdateInvoice(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	js, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	WriteAsJSON(w, js)
}

func NewInvoice(w http.ResponseWriter, r *http.Request) {
	var data models.Invoice
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := database.NewInvoice(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	js, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	WriteAsJSON(w, js)
}
