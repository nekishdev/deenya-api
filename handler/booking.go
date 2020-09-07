package handler

import (
	"deenya-api/database"
	"deenya-api/models"
	"net/http"
	"strconv"
	"time"

	"github.com/clarketm/json"

	"github.com/go-chi/chi"
)

func GetBooking(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var q string

	aid := GetAuthID(r)
	atype := GetAuthType(r)

	q = chi.URLParam(r, "bookingID")

	id, err = strconv.ParseInt(q, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	booking, err := database.GetBooking(id, aid, atype)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if *booking.ClientID != aid && *booking.ConsultantID != aid {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	js, err := json.Marshal(booking)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func UpdateBooking(w http.ResponseWriter, r *http.Request) {
	var data models.Booking

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	aid := GetAuthID(r)
	atype := GetAuthType(r)

	if atype == "client" {
		if aid != *data.ClientID {
			http.Error(w, UnauthErr.Error(), http.StatusUnauthorized)
		}
	}

	if atype == "consultant" {
		if aid != *data.ConsultantID {
			http.Error(w, UnauthErr.Error(), http.StatusUnauthorized)
		}
	}

	err := database.UpdateBooking(data, atype)
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

func DeleteBooking(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var q string

	q = chi.URLParam(r, "bookingID")

	id, err = strconv.ParseInt(q, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	aid := GetAuthID(r)
	atype := GetAuthType(r)

	if !database.VerifyBooking(id, aid, atype) {
		http.Error(w, UnauthErr.Error(), http.StatusUnauthorized)
		return
	}

	err = database.DeleteBooking(id, aid, atype)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := struct {
		Message string `json:"message"`
	}{
		Message: "Success",
	}

	js, err := json.Marshal(resp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func NewBooking(w http.ResponseWriter, r *http.Request) {
	var data models.Booking

	aid := GetAuthID(r)

	atype := GetAuthType(r)

	if atype != "client" {
		http.Error(w, UnauthErr.Error(), http.StatusBadRequest)
		return
	}

	data.ClientID = &aid

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//check if any bookings exist within start-end time range

	database.NewBooking(&data)

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func UserBookings(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var q string

	aid := GetAuthID(r)
	atype := GetAuthType(r)

	id, err = strconv.ParseInt(q, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bookings, err := database.UserBookings(aid, atype, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	js, err := json.Marshal(bookings)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func MyBookings(w http.ResponseWriter, r *http.Request) {

	var err error

	aid := GetAuthID(r)
	atype := GetAuthType(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bookings, err := database.MyBookings(aid, atype)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	js, err := json.Marshal(bookings)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func AcceptBooking(w http.ResponseWriter, r *http.Request) {
	var data models.Booking

	q := chi.URLParam(r, "bookingID")

	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	aid := GetAuthID(r)
	atype := GetAuthType(r)

	if atype != "consultant" {
		http.Error(w, UnauthErr.Error(), http.StatusUnauthorized)
		return
	}

	data.ConsultantID = &aid
	data.ID = &id

	accepted := true
	data.IsAccepted = &accepted

	database.UpdateBooking(data, atype)
}

func AvailableBookings(w http.ResponseWriter, r *http.Request) {
	body := struct {
		tz   string
		date time.Time
	}{}
	q := chi.URLParam(r, "userID")
	uid, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// now := time.Now()
	// start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, nil)
	// end := start.Add(24 * time.Hour)

	data, err := database.AvailableBookings(uid, body.date, body.tz)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var list []int64
	for _, t := range data {
		loc, _ := time.LoadLocation(`Asia/Shanghai`)
		d := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, loc)
		list = append(list, d.Unix())
	}

	js, err := json.Marshal(list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	WriteAsJSON(w, js)
}
