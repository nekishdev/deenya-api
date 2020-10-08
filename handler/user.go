package handler

import (
	"deenya-api/database"
	"deenya-api/models"
	"net/http"
	"strconv"

	"github.com/clarketm/json"
	"github.com/go-chi/chi"
)

// UserPublic godoc
// @Summary Get user public by user ID
// @Description Get user public by user ID
// @Tags User
// @ID get-user-public-by-id
// @Accept  json
// @Produce  json
// @Param userID path int true "User ID"
// @Success 200 {object} models.UserData
// @Failure 400 {object} interface{}
// @Router /users/{userID}/public [get]
func UserPublic(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var id int64
	var err error
	var q string

	q = chi.URLParam(r, "userID")

	id, err = strconv.ParseInt(q, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err = database.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	products, err := database.UserProducts(*user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, product := range products {
		user.Products = append(user.Products, &product)
	}

	portfolios, err := database.UserPortfolios(*user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, portfolio := range portfolios {
		user.Portfolios = append(user.Portfolios, &portfolio)
	}

	posts, err := database.UserPosts(*user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, post := range posts {
		user.Posts = append(user.Posts, &post)
	}

	js, err := json.Marshal(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func UserWithDetails(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var id int64
	var err error
	var q string

	q = chi.URLParam(r, "userID")

	id, err = strconv.ParseInt(q, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err = database.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	details, err := database.GetUserDetails(*user.ID)

	user.UserDetails = details

	js, err := json.Marshal(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// GetUser godoc
// @Summary Get user by ID
// @Description Get user by id
// @Tags User
// @ID get-user-by-id
// @Accept  json
// @Produce  json
// @Param userID path int true "User ID"
// @Success 200 {object} models.UserData
// @Failure 400 {object} interface{}
// @Router /users/{userID} [get]
func GetUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var id int64
	var err error
	var q string

	q = chi.URLParam(r, "userID")

	id, err = strconv.ParseInt(q, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err = database.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// products, err := database.UserProducts(*user.ID)

	// user.Products = products

	user.Email = nil

	js, err := json.Marshal(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// UpdateUser godoc
// @Summary Update user by ID
// @Description Update user by id
// @Tags User
// @ID update-user-by-id
// @Accept  json
// @Produce  json
// @Param userID path int true "User ID"
// @Param body body models.UserData true "User data"
// @Success 200 {object} models.UserData
// @Failure 400 {object} interface{}
// @Router /users/{userID} [put]
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var err error
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	q := chi.URLParam(r, "userID")

	uid, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mid := GetAuthID(r)

	if mid != uid {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user.ID = &uid

	err = database.UpdateUser(user.UserData)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	js, err := json.Marshal(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

// DeleteUser godoc
// @Summary Delete user by ID
// @Description Delete user by id
// @Tags User
// @ID delete-user-by-id
// @Accept  json
// @Produce  json
// @Param userID path int true "User ID"
// @Success 200 {object} models.JsonResultMessage
// @Failure 400 {object} interface{}
// @Router /users/{userID} [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var q string

	q = chi.URLParam(r, "userID")

	id, err = strconv.ParseInt(q, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = database.DeleteUser(id)

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

func SearchUsers(w http.ResponseWriter, r *http.Request) {

}

// MyAccount godoc
// @Summary My account
// @Description My account
// @Tags User
// @ID get-my-account
// @Accept  json
// @Produce  json
// @Param userID path int true "User ID"
// @Success 200 {object} models.UserData
// @Failure 400 {object} interface{}
// @Router /users/{userID}/me [get]
func MyAccount(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var err error

	mid := GetAuthID(r)
	user, err = database.GetUser(mid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	details, err := database.GetUserDetails(*user.ID)

	user.UserDetails = details

	js, err := json.Marshal(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
