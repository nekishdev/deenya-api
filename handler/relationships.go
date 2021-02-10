package handler

import (
	"deenya-api/database"
	"deenya-api/models"
	"net/http"

	"github.com/clarketm/json"
)

// MyClients godoc
// @Summary MyClients
// @Description MyClients
// @Tags Clients
// @ID my-client
// @Accept  json
// @Produce  json
// @Success 200 {array} models.UserData
// @Failure 400 {object} interface{}
// @Router /clients [get]
// @Security ApiKeyAuth
func MyClients(w http.ResponseWriter, r *http.Request) {
	var clients []models.User
	var err error

	mid := GetAuthID(r)

	if GetAuthType(r) == "client" {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	clients, err = database.MyClients(mid)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	js, err := json.Marshal(clients)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// MyConsultants godoc
// @Summary MyConsultants
// @Description MyConsultants
// @Tags Consultants
// @ID my-consultants
// @Accept  json
// @Produce  json
// @Success 200 {array} models.UserData
// @Failure 400 {object} interface{}
// @Router /consultants [get]
// @Security ApiKeyAuth
func MyConsultants(w http.ResponseWriter, r *http.Request) {
	var list []models.User
	var err error

	mid := GetAuthID(r)

	if GetAuthType(r) == "consultant" {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	list, err = database.MyConsultants(mid)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	js, err := json.Marshal(list)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
