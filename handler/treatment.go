package handler

import (
	"deenya-api/database"
	"deenya-api/models"
	"net/http"
	"strconv"

	"github.com/clarketm/json"
	"github.com/go-chi/chi"
)

// GetTreatment godoc
// @Summary Get treatment by ID
// @Description Get treatment by id
// @Tags Treatment
// @ID get-treatment-by-id
// @Accept  json
// @Produce  json
// @Param treatmentID path int true "Treatment ID"
// @Success 200 {object} models.TreatmentData
// @Failure 400 {object} interface{}
// @Router /treatments/{treatmentID} [get]
func GetTreatment(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var q string

	q = chi.URLParam(r, "treatmentID")

	id, err = strconv.ParseInt(q, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := database.GetTreatment(id)

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

// UpdateTreatment godoc
// @Summary Update Treatment by ID
// @Description Update Treatment by id
// @Tags Treatment
// @ID update-treatment-by-id
// @Accept  json
// @Produce  json
// @Param treatmentID path int true "Treatment ID"
// @Param body body models.TreatmentData true "Treatment data"
// @Success 200 {object} models.TreatmentData
// @Failure 400 {object} interface{}
// @Router /treatments/{treatmentID} [put]
func UpdateTreatment(w http.ResponseWriter, r *http.Request) {
	var data models.Treatment

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	q := chi.URLParam(r, "treatmentID")

	id, err := strconv.ParseInt(q, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data.ID = &id

	mid := GetAuthID(r)
	mtype := GetAuthType(r)

	if mtype == "client" {

		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return

	}

	if mtype == "consultant" {
		data.ConsultantID = &mid
		data.ClientID = nil
	}

	database.UpdateTreatment(data, mtype) //, uid

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// DeleteTreatment godoc
// @Summary Delete Treatment by ID
// @Description Delete Treatment by id
// @Tags Treatment
// @ID delete-treatment-by-id
// @Accept  json
// @Produce  json
// @Param treatmentID path int true "Treatment ID"
// @Success 200 {object} models.JsonResultMessage
// @Failure 400 {object} interface{}
// @Router /treatments/{treatmentID} [delete]
func DeleteTreatment(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var q string

	q = chi.URLParam(r, "treatmentID")

	mid := GetAuthID(r)
	mtype := GetAuthType(r)

	id, err = strconv.ParseInt(q, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = database.DeleteTreatment(id, mid, mtype)

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

func NewTreatment(w http.ResponseWriter, r *http.Request) {
	var data models.Treatment

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
			return
		}
	}

	if atype == "consultant" {
		if aid != *data.ConsultantID {
			http.Error(w, UnauthErr.Error(), http.StatusUnauthorized)
			return
		}
	}

	//verify client relation - todo

	database.NewTreatment(&data)

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// MyTreatments godoc
// @Summary My treatments
// @Description My treatments
// @Tags Treatment
// @ID my-treatments
// @Accept  json
// @Produce  json
// @Success 200 {array} models.TreatmentData
// @Failure 400 {object} interface{}
// @Router /treatments/ [get]
func MyTreatments(w http.ResponseWriter, r *http.Request) {

	aid := GetAuthID(r)
	t := GetAuthType(r)

	data, err := database.MyTreatments(aid, t)

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
