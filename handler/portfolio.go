package handler

import (
	"deenya-api/database"
	"deenya-api/models"
	"net/http"
	"strconv"

	"github.com/clarketm/json"

	"github.com/go-chi/chi"
)

func GetPortfolio(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var q string

	q = chi.URLParam(r, "portfolioID")

	id, err = strconv.ParseInt(q, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	portfolio, err := database.GetPortfolio(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	js, err := json.Marshal(portfolio)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func UpdatePortfolio(w http.ResponseWriter, r *http.Request) {
	var data models.Portfolio

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mid := GetAuthID(r)
	mtype := GetAuthType(r)

	if mtype == "client" {
		data.ClientID = &mid
		data.ConsultantID = nil
	}

	if mtype == "consultant" {
		data.ConsultantID = &mid
		data.ClientID = nil
	}

	database.UpdatePortfolio(data, mtype) //, uid

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func DeletePortfolio(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var q string

	q = chi.URLParam(r, "portfolioID")

	mid := GetAuthID(r)
	mtype := GetAuthType(r)

	id, err = strconv.ParseInt(q, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = database.DeletePortfolio(id, mid, mtype)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := struct {
		Message string `json:"message"`
	}{
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

func NewPortfolio(w http.ResponseWriter, r *http.Request) {
	var data models.Portfolio

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

	database.NewPortfolio(&data)

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// UserPortfolios godoc
// @Summary Get user portfolios by user ID
// @Description Get user portfolios by user ID
// @Tags User
// @ID get-user-portfolios-by-id
// @Accept  json
// @Produce  json
// @Param userID path int true "User ID"
// @Success 200 {array} models.PortfolioData
// @Failure 400 {object} interface{}
// @Router /users/{userID}/portfolios [get]
func UserPortfolios(w http.ResponseWriter, r *http.Request) {
	var uid int64
	var err error
	var q string
	var data []models.Portfolio

	q = chi.URLParam(r, "userID")

	uid, err = strconv.ParseInt(q, 10, 64)
	aid := GetAuthID(r)
	t := GetAuthType(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	portfolios, err := database.UserPortfolios(uid)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if portfolios == nil {
		http.Error(w, "No data found", http.StatusBadRequest)
		return
	}

	if uid != aid {
		for _, p := range portfolios {
			if *p.IsPublished == true {
				data = append(data, p)
			}
		}
	}

	if uid == aid && t == "consultant" {
		data = portfolios
	}

	js, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func MyPortfolios(w http.ResponseWriter, r *http.Request) {

	aid := GetAuthID(r)
	t := GetAuthType(r)

	data, err := database.MyPortfolios(aid, t)

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
