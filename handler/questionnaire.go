package handler

import (
	"deenya-api/database"
	"deenya-api/models"
	"net/http"

	"github.com/clarketm/json"
)

// GetQuestionnaire godoc
// @Summary Get questionnaire object by ID
// @Description Get questionnaire object by ID
// @Tags User
// @ID questionnaire-get
// @Accept  json
// @Produce  json
// @Param questionnaireID path int true "Questionnaire ID"
// @Success 200 {object} models.QuestionnaireData
// @Failure 400 {object} interface{}
// @Router /questionnaires/{questionnaireID} [get]
// @Security ApiKeyAuth
func GetQuestionnaire(w http.ResponseWriter, r *http.Request) {

}

// UpdateQuestionnaire godoc
// @Summary Update questionnaire object
// @Description Update questionnaire object
// @Tags User
// @ID questionnaire-update
// @Accept  json
// @Produce  json
// @Param questionnaireID path int true "Questionnaire ID"
// @Param body body models.QuestionnaireData true "Questionnaire Object"
// @Success 200 {object} models.QuestionnaireData
// @Failure 400 {object} interface{}
// @Router /questionnaires/{questionnaireID} [post]
// @Security ApiKeyAuth
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

// DeleteQuestionnaire godoc
// @Summary Delete questionnaire object
// @Description Delete questionnaire object
// @Tags User
// @ID questionnaire-delete
// @Accept  json
// @Produce  json
// @Param questionnaireID path int true "Questionnaire ID"
// @Success 200 {object} models.JsonResultMessage
// @Failure 400 {object} interface{}
// @Router /questionnaires/{questionnaireID} [put]
// @Security ApiKeyAuth
func DeleteQuestionnaire(w http.ResponseWriter, r *http.Request) {

}

// NewQuestionnaire godoc
// @Summary Create a new questionnaire object for a booking
// @Description Create a new questionnaire object for a booking
// @Tags User
// @ID questionnaire-new
// @Accept  json
// @Produce  json
// @Param bookingID path int true "Booking ID"
// @Param body body models.QuestionnaireData true "Questionnaire Object"
// @Success 200 {object} models.QuestionnaireData
// @Failure 400 {object} interface{}
// @Router /{bookingID}/questionnaire [post]
// @Security ApiKeyAuth
func NewQuestionnaire(w http.ResponseWriter, r *http.Request) {

}

func GetQuestion(w http.ResponseWriter, r *http.Request) {

}

// UpdateQuestion godoc
// @Summary Update question object
// @Description Update question object
// @Tags User
// @ID question-update
// @Accept  json
// @Produce  json
// @Param questionnaireID path int true "Questionnaire ID"
// @Param body body models.QuestionData true "Question Object"
// @Success 200 {object} models.QuestionData
// @Failure 400 {object} interface{}
// @Router /questionnaires/{questionnaireID} [post]
// @Security ApiKeyAuth
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

// DeleteQuestion godoc
// @Summary Delete a question object
// @Description Delete a question object
// @Tags User
// @ID question-delete
// @Accept  json
// @Produce  json
// @Param questionnaireID path int true "Questionnaire ID"
// @Param questionID path int true "Question ID"
// @Success 200 {object} models.JsonResultMessage
// @Failure 400 {object} interface{}
// @Router /questionnaires/{questionnaireID}/{questionID} [delete]
// @Security ApiKeyAuth
func DeleteQuestion(w http.ResponseWriter, r *http.Request) {

}

// NewQuestion godoc
// @Summary Create new question object for a questionnaire
// @Description Create new question object for a questionnaire
// @Tags User
// @ID question-new
// @Accept  json
// @Produce  json
// @Param questionnaireID path int true "Questionnaire ID"
// @Param body body models.QuestionData true "Question Object"
// @Success 200 {object} models.QuestionData
// @Failure 400 {object} interface{}
// @Router /questionnaires/{questionnaireID}/ [post]
// @Security ApiKeyAuth
func NewQuestion(w http.ResponseWriter, r *http.Request) {

}
