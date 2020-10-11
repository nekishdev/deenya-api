package handler

import (
	"deenya-api/database"
	"deenya-api/models"
	"net/http"
	"strconv"

	"github.com/clarketm/json"
	"github.com/go-chi/chi"
)

// @Summary GetForumThread
// @Description GetForumThread
// @Tags Forum
// @ID GetForumThread
// @Accept  json
// @Produce  json
// @Param threadID path int true "Thread id"
// @Success 200 {object} models.ForumThreadData
// @Failure 400 {object} interface{}
// @Router /forum/{threadID} [get]
func GetForumThread(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var q string

	q = chi.URLParam(r, "threadID")

	id, err = strconv.ParseInt(q, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := database.GetForumThread(id)

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

// @Summary UpdateForumThread
// @Description UpdateForumThread
// @Tags Forum
// @ID UpdateForumThread
// @Accept  json
// @Produce  json
// @Param threadID path int true "Thread id"
// @Param body body models.ForumThreadData true "Thread data"
// @Success 200 {object} models.ForumThreadData
// @Failure 400 {object} interface{}
// @Router /forum/{threadID} [put]
func UpdateForumThread(w http.ResponseWriter, r *http.Request) {
	var data models.ForumThread

	q := chi.URLParam(r, "threadID")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mid := GetAuthID(r)

	data.ID = &id

	data.OwnerID = &mid

	database.UpdateForumThread(data)

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// @Summary DeleteForumThread
// @Description DeleteForumThread
// @Tags Forum
// @ID DeleteForumThread
// @Accept  json
// @Produce  json
// @Param threadID path int true "Thread id"
// @Success 200 {object} models.JsonResultMessage
// @Failure 400 {object} interface{}
// @Router /forum/{threadID} [delete]
func DeleteForumThread(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var q string

	q = chi.URLParam(r, "threadID")

	mid := GetAuthID(r)

	id, err = strconv.ParseInt(q, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = database.DeleteForumThread(id, mid)

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

// @Summary NewForumThread
// @Description NewForumThread
// @Tags Forum
// @ID NewForumThread
// @Accept  json
// @Produce  json
// @Param body body models.ForumThreadData true "Thread data"
// @Success 200 {object} models.ForumThreadData
// @Failure 400 {object} interface{}
// @Router /forum [post]
func NewForumThread(w http.ResponseWriter, r *http.Request) {
	var data models.ForumThread

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	aid := GetAuthID(r)

	data.OwnerID = &aid

	database.NewForumThread(&data)

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// UserForumThreads godoc
// @Summary Get user threads by user ID
// @Description Get user threads by user ID
// @Tags User
// @ID get-user-threads-by-id
// @Accept  json
// @Produce  json
// @Param userID path int true "User ID"
// @Success 200 {array} models.ForumThreadData
// @Failure 400 {object} interface{}
// @Router /users/{userID}/threads [get]
func UserForumThreads(w http.ResponseWriter, r *http.Request) {
	var uid int64
	var q string

	q = chi.URLParam(r, "userID")

	uid, err := strconv.ParseInt(q, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := database.UserForumThreads(uid)

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

// @Summary MyForumThreads
// @Description MyForumThreads
// @Tags Forum
// @ID MyForumThreads
// @Accept  json
// @Produce  json
// @Success 200 {array} models.ForumThreadData
// @Failure 400 {object} interface{}
// @Router /forum [get]
func MyForumThreads(w http.ResponseWriter, r *http.Request) {

	aid := GetAuthID(r)

	data, err := database.MyForumThreads(aid)

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

func SearchForumThreads(w http.ResponseWriter, r *http.Request) {
	//search forum threads for query term + params
}

// @Summary FeedForumThreads
// @Description FeedForumThreads
// @Tags Forum
// @ID FeedForumThreads
// @Accept  json
// @Produce  json
// @Success 200 {array} models.ForumThreadData
// @Failure 400 {object} interface{}
// @Router /forum/feed [get]
func FeedForumThreads(w http.ResponseWriter, r *http.Request) {

	data, err := database.ForumThreadFeed()

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

// @Summary GetForumPost
// @Description GetForumPost
// @Tags Forum
// @ID GetForumPost
// @Accept  json
// @Produce  json
// @Param threadID path int true "Thread id"
// @Param postID path int true "Post id"
// @Success 200 {object} models.ForumPostData
// @Failure 400 {object} interface{}
// @Router /forum/{threadID}/{postID} [get]
func GetForumPost(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	q := chi.URLParam(r, "postID")

	id, err = strconv.ParseInt(q, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := database.GetForumPost(id)

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

// @Summary UpdateForumPost
// @Description UpdateForumPost
// @Tags Forum
// @ID UpdateForumPost
// @Accept  json
// @Produce  json
// @Param threadID path int true "Thread id"
// @Param postID path int true "Post id"
// @Param body body models.ForumPostData true "Post data"
// @Success 200 {object} models.ForumPostData
// @Failure 400 {object} interface{}
// @Router /forum/{threadID}/{postID} [put]
func UpdateForumPost(w http.ResponseWriter, r *http.Request) {
	var data models.ForumPost

	q := chi.URLParam(r, "postID")

	id, err := strconv.ParseInt(q, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mid := GetAuthID(r)

	data.ID = &id

	data.OwnerID = &mid

	err = database.UpdateForumPost(data)

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

// @Summary DeleteForumPost
// @Description DeleteForumPost
// @Tags Forum
// @ID DeleteForumPost
// @Accept  json
// @Produce  json
// @Param threadID path int true "Thread id"
// @Param postID path int true "Post id"
// @Success 200 {object} models.JsonResultMessage
// @Failure 400 {object} interface{}
// @Router /forum/{threadID}/{postID} [delete]
func DeleteForumPost(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error

	q := chi.URLParam(r, "postID")

	id, err = strconv.ParseInt(q, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mid := GetAuthID(r)

	err = database.DeleteForumPost(id, mid)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	js, err := json.Marshal(models.JsonResultMessage{
		Message: "Success",
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

// @Summary NewForumPost
// @Description NewForumPost
// @Tags Forum
// @ID NewForumPost
// @Accept  json
// @Produce  json
// @Param threadID path int true "Thread id"
// @Param body body models.ForumPostData true "Forum Post data"
// @Success 200 {object} models.ForumPostData
// @Failure 400 {object} interface{}
// @Router /forum/{threadID} [post]
func NewForumPost(w http.ResponseWriter, r *http.Request) {
	var data models.ForumPost

	q := chi.URLParam(r, "threadID")

	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mid := GetAuthID(r)

	data.ThreadID = &id
	data.OwnerID = &mid

	err = database.NewForumPost(&data)
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

/*func MyForumPosts(w http.ResponseWriter, r *http.Request) {

	var err error

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mid := GetAuthID(r)

	data, err := database.MyForumPosts(mid)

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

}*/

// @Summary ForumThreadPosts
// @Description ForumThreadPosts
// @Tags Forum
// @ID ForumThreadPosts
// @Accept  json
// @Produce  json
// @Param threadID path int true "Thread id"
// @Success 200 {array} models.ForumPostData
// @Failure 400 {object} interface{}
// @Router /forum/{threadID}/posts [get]
func ForumThreadPosts(w http.ResponseWriter, r *http.Request) {
	var data []models.ForumPost

	q := chi.URLParam(r, "threadID")

	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err = database.ThreadForumPosts(id)
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
