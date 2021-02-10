package handler

import (
	"deenya-api/database"
	"deenya-api/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/clarketm/json"

	"github.com/go-chi/chi"
)

// GetProduct godoc
// @Summary GetProduct
// @Description GetProduct
// @Tags Products
// @ID get-product
// @Accept  json
// @Produce  json
// @Param productID path int true "Product ID"
// @Success 200 {object} models.ProductData
// @Failure 400 {object} interface{}
// @Router /products/{productID} [get]
func GetProduct(w http.ResponseWriter, r *http.Request) {
	var data models.Product
	var id int64
	var err error
	var q string

	q = chi.URLParam(r, "productID")

	id, err = strconv.ParseInt(q, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err = database.GetProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data.Model, err = database.GetProductModel(*data.ProductModelID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	medias, err := database.ArrayMedia(data.Model.MediaIDs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, media := range medias {
		data.Model.Medias = append(data.Model.Medias, &media)
	}

	js, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// UpdateProduct godoc
// @Summary UpdateProduct
// @Description UpdateProduct
// @Tags Products
// @ID update-product
// @Accept  json
// @Produce  json
// @Param productID path int true "Product ID"
// @Param body body models.ProductData true "Product Object"
// @Success 200 {object} models.ProductData
// @Failure 400 {object} interface{}
// @Router /products/{productID} [put]
// @Security ApiKeyAuth
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var data models.Product
	var err error
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = database.UpdateProduct(data)

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

// DeleteProduct godoc
// @Summary DeleteProduct
// @Description DeleteProduct
// @Tags Products
// @ID delete-product
// @Accept  json
// @Produce  json
// @Param productID path int true "Product ID"
// @Success 200 {object} models.JsonResultMessage
// @Failure 400 {object} interface{}
// @Router /products/{productID} [delete]
// @Security ApiKeyAuth
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var q string

	mid := GetAuthID(r)

	q = chi.URLParam(r, "productID")

	id, err = strconv.ParseInt(q, 10, 64)

	if err != nil {
		fmt.Println(err)
	}

	err = database.DeleteProduct(id, mid)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	msg := models.JsonResultMessage{
		Message: "Success",
	}

	js, err := json.Marshal(msg)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

// NewProduct godoc
// @Summary NewProduct
// @Description NewProduct
// @Tags Products
// @ID new-product
// @Accept  json
// @Produce  json
// @Param body body models.ProductData true "Product Object"
// @Success 200 {object} models.ProductData
// @Failure 400 {object} interface{}
// @Router /products [post]
func NewProduct(w http.ResponseWriter, r *http.Request) {
	var data models.Product

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.NewProduct(&data)

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// UserProducts godoc
// @Summary Get user products by user ID
// @Description Get user products by user ID
// @Tags User
// @ID get-user-products-by-id
// @Accept  json
// @Produce  json
// @Param userID path int true "User ID"
// @Success 200 {array} models.ProductData
// @Failure 400 {object} interface{}
// @Router /users/{userID}/products [get]
func UserProducts(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var q string

	q = chi.URLParam(r, "userID")

	id, err = strconv.ParseInt(q, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := database.UserProducts(id)

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

// MyProducts godoc
// @Summary MyProducts
// @Description MyProducts
// @Tags Products
// @ID my-products
// @Accept  json
// @Produce  json
// @Success 200 {array} models.ProductData
// @Failure 400 {object} interface{}
// @Router /products [get]
// @Security ApiKeyAuth
func MyProducts(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error

	id = GetAuthID(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := database.MyProducts(id)

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

func UserProductsPage(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var q string
	var data models.User

	q = chi.URLParam(r, "userID")

	id, err = strconv.ParseInt(q, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err = database.GetUser(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	products, err := database.UserProducts(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, product := range products {
		data.Products = append(data.Products, &product)
	}

	js, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// SuggestProductModels godoc
// @Summary SuggestProductModels
// @Description SuggestProductModels
// @Tags Products
// @ID suggest-product-models
// @Accept  json
// @Produce  json
// @Param body body models.ProductModelSuggestRequest true "Query"
// @Success 200 {array} models.ProductModelData
// @Failure 400 {object} interface{}
// @Router /products/models/suggest [get]

func SuggestProductModels(w http.ResponseWriter, r *http.Request) {
	q := models.ProductModelSuggestRequest{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&q); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := database.SuggestProductModels(q.Tags)
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

func NewProductModel(w http.ResponseWriter, r *http.Request) {
}

func UpdateProductModel(w http.ResponseWriter, r *http.Request) {
}

func GetProductModel(w http.ResponseWriter, r *http.Request) {
}
