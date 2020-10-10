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

// GetBooking godoc
// @Summary Get booking by ID
// @Description Get booking by id
// @Tags Booking
// @ID get-booking-by-id
// @Accept  json
// @Produce  json
// @Param bookingID path int true "Booking ID"
// @Success 200 {object} models.BookingData
// @Failure 400 {object} interface{}
// @Router /bookings/{bookingID} [get]
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

// UpdateBooking godoc
// @Summary Update booking by ID
// @Description Update booking by id
// @Tags Booking
// @ID update-booking-by-id
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param bookingID path int true "Booking ID"
// @Param body body models.BookingData true "Booking data"
// @Success 200 {object} models.BookingData
// @Failure 400 {object} interface{}
// @Router /bookings/{bookingID} [put]
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

// DeleteBooking godoc
// @Summary Delete booking by ID
// @Description Delete booking by id
// @Tags Booking
// @ID delete-booking-by-id
// @Accept  json
// @Produce  json
// @Param bookingID path int true "Booking ID"
// @Success 200 {object} models.JsonResultMessage
// @Failure 400 {object} interface{}
// @Router /bookings/{bookingID} [delete]
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

	resp := models.JsonResultMessage{
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

// NewBooking godoc
// @Summary New booking
// @Description New booking
// @Tags Booking
// @ID new-booking
// @Accept  json
// @Produce  json
// @Param body body models.BookingData true "Booking"
// @Success 200 {object} models.BookingData
// @Failure 400 {object} interface{}
// @Router /bookings/ [post]
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

// MyBookings godoc
// @Summary My bookings
// @Description My bookings
// @Tags Booking
// @ID my-bookings
// @Accept  json
// @Produce  json
// @Success 200 {array} models.BookingData
// @Failure 400 {object} interface{}
// @Router /bookings/ [get]
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

// AcceptBooking godoc
// @Summary Allows a consultant to accept a booking request
// @Description Allows a consultant to accept a booking request
// @Tags Booking
// @ID accept-booking
// @Accept  json
// @Produce  json
// @Success 200 {array} models.BookingData
// @Failure 400 {object} interface{}
// @Router /booking/{bookingID}/accept [get]
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

// AvailableBookings godoc
// @Summary AvailableBookings
// @Description AvailableBookings
// @Tags User
// @ID available-bookings
// @Accept  json
// @Produce  json
// @Param userID path int true "User ID"
// @Success 200 {array} integer
// @Failure 400 {object} interface{}
// @Router /users/{userID}/available [get]
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
