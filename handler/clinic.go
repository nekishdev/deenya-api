package handler

import (
	"deenya-api/database"
	"deenya-api/models"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/clarketm/json"
)

// NewClinic godoc
// @Summary NewClinic
// @Description NewClinic
// @Tags Clinic
// @ID NewClinic
// @Accept  json
// @Produce  json
// @Param body body models.ClinicData true "Clinic data"
// @Success 200 {object} models.ClinicData
// @Failure 400 {object} interface{}
// @Router /clinics [post]
func NewClinic(w http.ResponseWriter, r *http.Request) {
	var data models.Clinic
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := database.NewClinic(&data)

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

// GetClinic godoc
// @Summary GetClinic
// @Description GetClinic
// @Tags Clinic
// @ID GetClinic
// @Accept  json
// @Produce  json
// @Param clinicID path int true "Clinic ID"
// @Success 200 {object} models.ClinicData
// @Failure 400 {object} interface{}
// @Router /clinics/{clinicID} [get]
func GetClinic(w http.ResponseWriter, r *http.Request) {
	q := chi.URLParam(r, "clinicID")

	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := database.GetClinic(id)
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

// UpdateClinic godoc
// @Summary UpdateClinic
// @Description UpdateClinic
// @Tags Clinic
// @ID UpdateClinic
// @Accept  json
// @Produce  json
// @Param clinicID path int true "Clinic ID"
// @Param body body models.ClinicData true "Clinic data"
// @Success 200 {object} models.ClinicData
// @Failure 400 {object} interface{}
// @Router /clinics/{clinicID} [put]
func UpdateClinic(w http.ResponseWriter, r *http.Request) {
	var data models.Clinic
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mid := GetAuthID(r)

	err := database.UpdateClinic(data, mid)

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

// NewClinicRequest godoc
// @Summary NewClinicRequest
// @Description NewClinicRequest
// @Tags Clinic
// @ID NewClinicRequest
// @Accept  json
// @Produce  json
// @Param clinicID path int true "Clinic ID"
// @Success 200 {object} models.ClinicMemberData
// @Failure 400 {object} interface{}
// @Router /clinics/{clinicID} [post]
func NewClinicRequest(w http.ResponseWriter, r *http.Request) {
	var data models.ClinicMember
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mid := GetAuthID(r)

	data.ConsultantID = &mid
	data.IsAccepted = nil

	err := database.NewClinicRequest(&data)

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

// @Summary AcceptClinicRequest
// @Description AcceptClinicRequest
// @Tags Clinic
// @ID AcceptClinicRequest
// @Accept  json
// @Produce  json
// @Param clinicID path int true "Clinic ID"
// @Param consultantID path int true "Consultant ID"
// @Success 200 {object} models.JsonResultMessage
// @Failure 400 {object} interface{}
// @Router /clinics/{clinicID}/{consultantID}/accept [get]
func AcceptClinicRequest(w http.ResponseWriter, r *http.Request) {
	q := chi.URLParam(r, "clinicID")

	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	q2 := chi.URLParam(r, "consultantID")

	tid, err := strconv.ParseInt(q2, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mid := GetAuthID(r)

	err = database.AcceptClinicRequest(id, mid, tid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	message := models.JsonResultMessage{
		Message: "Success",
	}

	js, err := json.Marshal(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	WriteAsJSON(w, js)

}

// ListClinicRequests godoc
// @Summary ListClinicRequests
// @Description ListClinicRequests
// @Tags Clinic
// @ID ListClinicRequests
// @Accept  json
// @Produce  json
// @Param clinicID path int true "Clinic ID"
// @Success 200 {array} models.ClinicMemberData
// @Failure 400 {object} interface{}
// @Router /clinics/{clinicID} [get]
func ListClinicRequests(w http.ResponseWriter, r *http.Request) {
	q := chi.URLParam(r, "clinicID")

	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mid := GetAuthID(r)

	data, err := database.ListClinicRequests(id, mid)
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

// ListClinicConsultants godoc
// @Summary ListClinicConsultants
// @Description ListClinicConsultants
// @Tags Clinic
// @ID ListClinicConsultants
// @Accept  json
// @Produce  json
// @Param clinicID path int true "Clinic ID"
// @Success 200 {array} models.UserData
// @Failure 400 {object} interface{}
// @Router /clinics/{clinicID} [get]
func ListClinicConsultants(w http.ResponseWriter, r *http.Request) {
	q := chi.URLParam(r, "clinicID")

	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := database.ListClinicConsultants(id)
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

// @Summary RemoveClinicMember
// @Description RemoveClinicMember
// @Tags Clinic
// @ID RemoveClinicMember
// @Accept  json
// @Produce  json
// @Param clinicID path int true "Clinic ID"
// @Param consultantID path int true "Consultant ID"
// @Success 200 {object} models.JsonResultMessage
// @Failure 400 {object} interface{}
// @Router /clinics/{clinicID}/{consultantID}/remove [get]
func RemoveClinicMember(w http.ResponseWriter, r *http.Request) {
	q := chi.URLParam(r, "clinicID")

	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	q2 := chi.URLParam(r, "consultantID")

	tid, err := strconv.ParseInt(q2, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mid := GetAuthID(r)

	err = database.RemoveClinicMember(id, mid, tid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	message := models.JsonResultMessage{
		Message: "Success",
	}

	js, err := json.Marshal(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	WriteAsJSON(w, js)

}

// LeaveClinic godoc
// @Summary LeaveClinic
// @Description LeaveClinic
// @Tags Clinic
// @ID LeaveClinic
// @Accept  json
// @Produce  json
// @Param clinicID path int true "Clinic ID"
// @Success 200 {object} models.JsonResultMessage
// @Failure 400 {object} interface{}
// @Router /clinics/{clinicID} [delete]
func LeaveClinic(w http.ResponseWriter, r *http.Request) {
	q := chi.URLParam(r, "clinicID")

	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mid := GetAuthID(r)

	err = database.LeaveClinic(id, mid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	message := models.JsonResultMessage{
		Message: "Success",
	}

	js, err := json.Marshal(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	WriteAsJSON(w, js)
}

// MyClinic godoc
// @Summary MyClinic
// @Description MyClinic
// @Tags Clinic
// @ID MyClinic
// @Accept  json
// @Produce  json
// @Success 200 {object} models.ClinicData
// @Failure 400 {object} interface{}
// @Router /clinics [get]
func MyClinic(w http.ResponseWriter, r *http.Request) {

	mid := GetAuthID(r)

	data, err := database.MyClinic(mid)
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
