package handler

import (
	"deenya-api/database"
	"deenya-api/models"
	"net/http"

	"github.com/clarketm/json"
)

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
