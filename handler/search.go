package handler

import (
	"deenya-api/database"
	"deenya-api/models"
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

// SearchProductModels godoc
// @Summary SearchProductModels
// @Description SearchProductModels
// @Tags Products
// @ID search-product-models
// @Accept  json
// @Produce  json
// @Param body body models.ProductModelSearchRequest true "Query"
// @Success 200 {array} models.ProductModelData
// @Failure 400 {object} interface{}
// @Router /products/models/search [get]
func SearchProductModels(w http.ResponseWriter, r *http.Request) {
	q := models.ProductModelSearchRequest{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&q); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := database.SearchProductModels(q.Query, q.Tags)
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
