package handler

import (
	"deenya-api/database"
	"deenya-api/models"
	"net/http"
	"strconv"

	"github.com/clarketm/json"
	"github.com/go-chi/chi"
)

// GetMedia godoc
// @Summary GetMedia
// @Description GetMedia
// @Tags Media
// @ID get-media-by-id
// @Accept  json
// @Produce  json
// @Param mediaID path int true "Media ID"
// @Success 200 {object} models.MediaData
// @Failure 400 {object} interface{}
// @Router /media/{mediaID} [get]
// @Security ApiKeyAuth
func GetMedia(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var q string

	q = chi.URLParam(r, "mediaID")

	id, err = strconv.ParseInt(q, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	media, err := database.GetMedia(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	js, err := json.Marshal(media)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// UpdateMedia godoc
// @Summary UpdateMedia
// @Description UpdateMedia
// @Tags Media
// @ID update-media-by-id
// @Accept  json
// @Produce  json
// @Param mediaID path int true "Media ID"
// @Param body body models.MediaData true "Media Object"
// @Success 200 {object} models.MediaData
// @Failure 400 {object} interface{}
// @Router /media/{mediaID} [put]
// @Security ApiKeyAuth
func UpdateMedia(w http.ResponseWriter, r *http.Request) {
	var data models.Media

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.NewMedia(&data)

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// DeleteMedia godoc
// @Summary DeleteMedia
// @Description DeleteMedia
// @Tags Media
// @ID delete-media-by-id
// @Accept  json
// @Produce  json
// @Param mediaID path int true "Media ID"
// @Success 200 {object} models.JsonResultMessage
// @Failure 400 {object} interface{}
// @Router /media/{mediaID} [delete]
// @Security ApiKeyAuth
func DeleteMedia(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var q string

	q = chi.URLParam(r, "mediaID")

	id, err = strconv.ParseInt(q, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = database.DeleteMedia(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := models.JsonResultMessage{
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

// NewMedia godoc
// @Summary NewMedia
// @Description NewMedia
// @Tags Media
// @ID new-media
// @Accept  json
// @Produce  json
// @Param body body models.MediaData true "Media Object"
// @Success 200 {object} models.MediaData
// @Failure 400 {object} interface{}
// @Router /media [post]
// @Security ApiKeyAuth
func NewMedia(w http.ResponseWriter, r *http.Request) {
	var data models.Media

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.NewMedia(&data)

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func UserMedia(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var q string

	q = chi.URLParam(r, "userID")

	id, err = strconv.ParseInt(q, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	medias, err := database.ListMedias(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	js, err := json.Marshal(medias)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// MyMedia godoc
// @Summary My media
// @Description My media list
// @Tags Media
// @ID my-media
// @Accept  json
// @Produce  json
// @Success 200 {array} models.MediaData
// @Failure 400 {object} interface{}
// @Router /media [get]
// @Security ApiKeyAuth
func MyMedia(w http.ResponseWriter, r *http.Request) {
	var err error

	mid := GetAuthID(r)

	medias, err := database.MyMedia(mid)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	js, err := json.Marshal(medias)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
