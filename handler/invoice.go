package handler

import (
	"deenya-api/database"
	"deenya-api/models"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/clarketm/json"
)

// @Summary MyInvoices
// @Description MyInvoices
// @Tags Finance
// @ID MyInvoices
// @Accept  json
// @Produce  json
// @Success 200 {array} models.InvoiceData
// @Failure 400 {object} interface{}
// @Router /finance/invoices [get]
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

// @Summary GetInvoice
// @Description GetInvoice
// @Tags Finance
// @ID GetInvoice
// @Accept  json
// @Produce  json
// @Param invoiceID path int true "Invoice ID"
// @Success 200 {object} models.InvoiceData
// @Failure 400 {object} interface{}
// @Router /finance/{invoiceID} [get]
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

// @Summary UpdateInvoice
// @Description UpdateInvoice
// @Tags Finance
// @ID UpdateInvoice
// @Accept  json
// @Produce  json
// @Param invoiceID path int true "Invoice ID"
// @Param body body models.InvoiceData true "Invoice data"
// @Success 200 {object} models.InvoiceData
// @Failure 400 {object} interface{}
// @Router /finance/{invoiceID} [put]
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

// @Summary NewInvoice
// @Description NewInvoice
// @Tags Finance
// @ID NewInvoice
// @Accept  json
// @Produce  json
// @Param body body models.InvoiceData true "Invoice data"
// @Success 200 {object} models.InvoiceData
// @Failure 400 {object} interface{}
// @Router /finance/invoices [post]
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
