package handler

import (
	"deenya-api/database"
	"deenya-api/models"
	"deenya-api/stripe_functions"
	"deenya-api/util"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/clarketm/json"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
)

func GenerateState() string {
	random := util.RandomString(32)
	return random
}

func InitiateInvoicePayment(w http.ResponseWriter, r *http.Request) {

	q := chi.URLParam(r, "invoiceID")

	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mid := GetAuthID(r)
	mtype := GetAuthType(r)

	invoice, err := database.GetInvoice(id, mid, mtype)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fee := *invoice.Total / 5
	//app commission 20%

	customer, err := database.GetCustomer(mid)
	if err != nil {
		http.Error(w, fmt.Sprintf("error get customer info, %s", err.Error()), http.StatusBadRequest)
		return
	}

	stripe.Key = stripe_functions.STRIPE_TEST_SECRET
	params := &stripe.PaymentIntentParams{
		// PaymentMethodTypes: stripe.StringSlice([]string{
		// 	"card",
		// }),
		SetupFutureUsage:     stripe.String(string(stripe.PaymentIntentSetupFutureUsageOffSession)),
		Amount:               stripe.Int64(*invoice.Total),
		Currency:             stripe.String(string(stripe.CurrencyGBP)),
		ApplicationFeeAmount: stripe.Int64(fee),
		Customer:             stripe.String(*customer.CustomerToken),
	}

	params.AddMetadata("invoice_id", strconv.FormatInt(*invoice.ID, 10))
	if *invoice.IsBooking {
		params.AddMetadata("invoice_type", "booking")
	}
	if *invoice.IsOrder {
		params.AddMetadata("invoice_type", "order")
	}

	connect, err := database.GetStripeConnect(*invoice.ConsultantID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	params.SetStripeAccount(*connect.AccountToken)
	pi, _ := paymentintent.New(params)

	resp := struct {
		ClientSecret string `json:"client_secret"`
	}{ClientSecret: pi.ClientSecret}

	js, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	WriteAsJSON(w, js)
	//save to database invoice
}

func SavePaymentMethod(w http.ResponseWriter, r *http.Request) {
	//var method stripe.PaymentMethod
	//var data models.StripePaymentMethod

	// mid := GetAuthID(r)

	// data.UserID = mid

	//update invoice
}

// func GenerateConnectAuth() {
// 	state := GenerateState()

// 	url := `https://connect.stripe.com/express/oauth/authorize?client_id=%s&state=%s&suggested_capabilities[]=transfers&stripe_user[email]=%s`

// 	//fmt.Sprintf(url, clientID, state, email)
// }

func NewConnectAccount(w http.ResponseWriter, r *http.Request) {

	body := struct {
		code string
	}{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mid := GetAuthID(r)

	token, err := stripe_functions.GetAccountAuthToken(body.code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data := models.StripeConnect{
		AccountToken: &token,
		ConsultantID: &mid,
	}

	err = database.NewStripeConnect(data)
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
	//save to database customer
}

// func AddCard() {
// 	var card stripe.Card
// 	//save to database source
// }

// func RemoveCard() {
// 	var source_token string
// 	//remove from db
// }

// func SetDefaultCard() {
// 	var source_token string
// 	//save to db
// }

// func UpdateBankAccount() {
// 	var bank stripe.BankAccount
// 	//save to db and stripe
// }
