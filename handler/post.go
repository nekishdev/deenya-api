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

// GetPost godoc
// @Summary GetPost
// @Description GetPost
// @Tags Post
// @ID GetPost
// @Accept  json
// @Produce  json
// @Success 200 {object} models.PostData
// @Failure 400 {object} interface{}
// @Router /posts [get]
func GetPost(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var q string

	q = chi.URLParam(r, "postID")

	id, err = strconv.ParseInt(q, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	post, err := database.GetPostWithOwner(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// owner, err := database.GetUser(*post.OwnerID)

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// post.Owner = &owner

	js, err := json.Marshal(post)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// UpdatePost godoc
// @Summary UpdatePost
// @Description UpdatePost
// @Tags Post
// @ID UpdatePost
// @Accept  json
// @Produce  json
// @Param postID path int true "Post ID"
// @Param body body models.PostData true "Post data"
// @Success 200 {object} models.PostData
// @Failure 400 {object} interface{}
// @Router /posts/{postID} [put]
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	var data models.Post

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	post, err := database.GetPost(*data.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if *post.OwnerID != GetAuthID(r) {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	database.NewPost(&data)

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// DeletePost godoc
// @Summary DeletePost
// @Description DeletePost
// @Tags Post
// @ID DeletePost
// @Accept  json
// @Produce  json
// @Param postID path int true "Post ID"
// @Success 200 {object} models.JsonResultMessage
// @Failure 400 {object} interface{}
// @Router /posts/{postID} [delete]
func DeletePost(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	var q string

	q = chi.URLParam(r, "postID")

	id, err = strconv.ParseInt(q, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	post, err := database.GetPost(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if *post.OwnerID != GetAuthID(r) {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	err = database.DeletePost(id)

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

// NewPost godoc
// @Summary NewPost
// @Description NewPost
// @Tags Post
// @ID NewPost
// @Accept  json
// @Produce  json
// @Param body body models.PostData true "Post data"
// @Success 200 {object} models.PostData
// @Failure 400 {object} interface{}
// @Router /posts [post]
func NewPost(w http.ResponseWriter, r *http.Request) {
	var data models.Post

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//check for every handler, set owner_id = auth id
	fmt.Println(data)
	mid := GetAuthID(r)

	data.OwnerID = &mid

	fmt.Println(mid)

	database.NewPost(&data)

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// UserPosts godoc
// @Summary Get user posts by user ID
// @Description Get user posts by user ID
// @Tags User
// @ID get-user-posts-by-id
// @Accept  json
// @Produce  json
// @Param userID path int true "User ID"
// @Success 200 {array} models.PostData
// @Failure 400 {object} interface{}
// @Router /users/{userID}/posts [get]
func UserPosts(w http.ResponseWriter, r *http.Request) {
	var err error

	q := chi.URLParam(r, "userID")

	uid, err := strconv.ParseInt(q, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	posts, err := database.UserPosts(uid)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	js, err := json.Marshal(posts)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// MyPosts godoc
// @Summary MyPosts
// @Description MyPosts
// @Tags Post
// @ID MyPosts
// @Accept  json
// @Produce  json
// @Success 200 {array} models.PostData
// @Failure 400 {object} interface{}
// @Router /posts [get]
func MyPosts(w http.ResponseWriter, r *http.Request) {
	var err error

	mid := GetAuthID(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	posts, err := database.MyPosts(mid)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	js, err := json.Marshal(posts)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func SlugPost(w http.ResponseWriter, r *http.Request) {
}

// func PostCtx(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		q := chi.URLParam(r, "postID")

// 		id, err := strconv.ParseInt(q, 10, 64)
// 		if err != nil {
// 			http.Error(w, http.StatusText(404), 404)
// 			return
// 		}

// 		post, err := database.GetPost(id)
// 		if err != nil {
// 			http.Error(w, http.StatusText(404), 404)
// 			return
// 		}

// 		ctx := context.WithValue(r.Context(), "post", post)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }
