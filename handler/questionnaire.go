package handler

import (
	"deenya-api/database"
	"deenya-api/models"
	"net/http"
	"strconv"

	"github.com/clarketm/json"
	"github.com/go-chi/chi"
)

func GetQuestionnaire(w http.ResponseWriter, r *http.Request) {

}

func UpdateQuestionnaire(w http.ResponseWriter, r *http.Request) {

	var data models.Questionnaire
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := database.UpdateQuestionnaire(data)

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

func DeleteQuestionnaire(w http.ResponseWriter, r *http.Request) {

}

func NewQuestionnaire(w http.ResponseWriter, r *http.Request) {

}

func UserQuestionnaires(w http.ResponseWriter, r *http.Request) {

	q := chi.URLParam(r, "userID")

	id, err := strconv.ParseInt(q, 10, 64)

	mid := GetAuthID(r)
	mtype := GetAuthType(r)

	data, err := database.UserQuestionnaires(mid, mtype, id)

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

func MyQuestionnaires(w http.ResponseWriter, r *http.Request) {

	mid := GetAuthID(r)
	mtype := GetAuthType(r)

	data, err := database.MyQuestionnaires(mid, mtype)

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

func GetQuestion(w http.ResponseWriter, r *http.Request) {

}

func UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	var data models.Question

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	t := GetAuthType(r)

	if t == "client" {
		data.Question = nil
	}

	if t == "consultant" {
		data.Answer = nil
	}

}

func DeleteQuestion(w http.ResponseWriter, r *http.Request) {

}

func NewQuestion(w http.ResponseWriter, r *http.Request) {

}

// func ListQuestions(w http.ResponseWriter, r *http.Request) {
// 	q := chi.URLParam(r, "questionnaireID")

// 	id, err := strconv.ParseInt(q, 10, 64)

// 	data, err := database.ListQuestions(id)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	js, err := json.Marshal(data)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(js)

// }
