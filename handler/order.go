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

// GetOrder godoc
// @Summary Get order object
// @Description Get order object
// @Tags Order
// @ID order-get
// @Accept  json
// @Produce  json
// @Param orderID path int true "Order ID"
// @Success 200 {object} models.OrderData
// @Failure 400 {object} interface{}
// @Router /orders/{orderID} [get]
func GetOrder(w http.ResponseWriter, r *http.Request) {
	var data models.Order
	var id int64
	var err error
	var q string

	q = chi.URLParam(r, "orderID")

	id, err = strconv.ParseInt(q, 10, 64)

	if err != nil {
		fmt.Println(err)
	}

	mid := GetAuthID(r)
	mtype := GetAuthType(r)

	data, err = database.GetOrder(id, mid, mtype)
	if err != nil {
		fmt.Println(err)
	}

	data.OrderItems, err = database.ListOrderItems(id)
	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		fmt.Println(err)
	}

	js, err := json.Marshal(data)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// UpdateOrder godoc
// @Summary Update order object
// @Description Update order object
// @Tags Order
// @ID order-update
// @Accept  json
// @Produce  json
// @Param orderID path int true "Order ID"
// @Param order body models.OrderData true "Update order"
// @Success 200 {object} models.OrderData
// @Failure 400 {object} interface{}
// @Router /orders/{orderID} [put]
func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	var data models.Order
	var err error
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mid := GetAuthID(r)
	mtype := GetAuthType(r)

	err = database.UpdateOrder(data, mid, mtype)

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

// func GetDiscountSecret(w http.ResponseWriter, r *http.Request) {

// 	key, err := database.GetDiscountKey(uid, discount, consultant_id)
// }

// func ApplyDiscount(w http.ResponseWriter, r *http.Request) {
// 	body := struct {
// 	}{}
// 	resp := struct {
// 		discount_secret string
// 	}{}

// 	// /w.Write(js)
// }

// NewOrder godoc
// @Summary Update order object
// @Description Update order object
// @Tags Order
// @ID new-order
// @Accept  json
// @Produce  json
// @Param orderID path int true "Order ID"
// @Param order body models.OrderData true "Update order"
// @Success 200 {object} models.OrderData
// @Failure 400 {object} interface{}
// @Router /orders/{orderID} [put]
func NewOrder(w http.ResponseWriter, r *http.Request) {
	body := struct {
		items        []models.OrderItem
		discount_key string
	}{}

	mid := GetAuthID(r)

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sellers := make(map[int64][]models.OrderItem)

	// for _, item := range body.items {
	// 	_, ok := sellers[*item.SellerID]

	// 	if !ok {
	// 		sellers[*item.SellerID] = []models.OrderItem{item}
	// 		continue
	// 	}
	// 	sellers[*item.SellerID] = append(sellers[*item.SellerID], item)
	// }
	var orders []models.Order
	for consultant_id, items := range sellers {
		var data models.Order

		data.ClientID = &mid
		data.ConsultantID = &consultant_id
		var total int64
		total = 0
		for _, item := range items {
			total += *item.Product.Price
		}
		err := database.NewOrder(&data)
		if err != nil {
			for _, o := range orders {
				database.DeleteOrder(*o.ID)
			}
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		orders = append(orders, data)
	}

	js, err := json.Marshal(orders)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func UserOrders(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var q string

	q = chi.URLParam(r, "userID")

	id, err = strconv.ParseInt(q, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mid := GetAuthID(r)
	mtype := GetAuthType(r)

	data, err := database.UserOrders(mid, mtype, id)

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
