package handler

import (
	"deenya-api/database"
	"net/http"

	"github.com/clarketm/json"
)

func SearchPosts(w http.ResponseWriter, r *http.Request) {

}

func SearchConsultants(w http.ResponseWriter, r *http.Request) {

}

func SearchClinics(w http.ResponseWriter, r *http.Request) {}

func SearchProducts(w http.ResponseWriter, r *http.Request) {

}

func SearchPortfolios(w http.ResponseWriter, r *http.Request) {

}

func SearchProductModels(w http.ResponseWriter, r *http.Request) {
	q := struct {
		query string
		tags  []string
	}{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&q); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := database.SearchProductModels(q.query, q.tags)
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
