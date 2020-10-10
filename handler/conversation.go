package handler

import (
	"deenya-api/database"
	"deenya-api/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/clarketm/json"

	"github.com/go-chi/chi"
)

// GetConversation godoc
// @Summary Get conversation object by ID
// @Description Get conversation object by ID
// @Tags Conversation
// @ID conversation-get
// @Accept  json
// @Produce  json
// @Param id path int true "Conversation ID"
// @Success 200 {object} models.ConversationData
// @Failure 400 {object} interface{}
// @Router /conversations/{id} [get]
// @Security Basic

func GetConversation(w http.ResponseWriter, r *http.Request) {
	// var verified bool

	q := chi.URLParam(r, "conversationID")

	id, err := strconv.ParseInt(q, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mid := GetAuthID(r)

	data, err := database.GetConversation(id, mid) //check if uid is authed to get data in every request

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// for _, p := range data.ParticipantID {
	// 	if *p == mid {
	// 		verified = true
	// 		break
	// 	}
	// }

	// if !verified {
	// 	http.Error(w, errors.New("Unverified user").Error(), http.StatusBadRequest)
	// 	return
	// }

	js, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

// UpdateConversation godoc
// @Summary Update conversation object
// @Description Update conversation object
// @Tags Conversation
// @ID conversation-update
// @Accept  json
// @Produce  json
// @Param id path int true "Conversation ID"
// @Param conversation body models.ConversationData true "Conversation"
// @Success 200 {object} models.ConversationData
// @Failure 400 {object} interface{}
// @Router /conversations/{id} [put]
// @Security Basic

func UpdateConversation(w http.ResponseWriter, r *http.Request) {
	var data models.Conversation

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mid := GetAuthID(r)

	// for _, p := range data.ParticipantID {
	// 	if *p == uid {
	// 		verified = true
	// 		break
	// 	}
	// }

	// if !verified {
	// 	http.Error(w, errors.New("Unverified user").Error(), http.StatusBadRequest)
	// 	return
	// }

	err := database.UpdateConversation(data, mid)

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

// DeleteConversation godoc
// @Summary Delete conversation object
// @Description Delete conversation object
// @Tags Conversation
// @ID conversation-update
// @Accept  json
// @Produce  json
// @Param id path int true "Conversation ID"
// @Success 200 {object} models.JsonResultMessage
// @Failure 400 {object} interface{}
// @Router /conversations/{id} [delete]
// @Security Basic

func DeleteConversation(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error

	q := chi.URLParam(r, "conversationID")

	id, err = strconv.ParseInt(q, 10, 64)

	mid := GetAuthID(r)

	database.DeleteConversation(id, mid)

	msg := struct {
		Message string
	}{
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

// NewConversation godoc
// @Summary Create conversation object
// @Description Create conversation object
// @Tags Conversation
// @ID conversation-new
// @Accept  json
// @Produce  json
// @Param conversation body models.ConversationData true "Conversation"
// @Success 200 {object} models.ConversationData
// @Failure 400 {object} interface{}
// @Router /conversations [post]
// @Security Basic

func NewConversation(w http.ResponseWriter, r *http.Request) {
	var data models.Conversation
	var verified bool

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	aid := GetAuthID(r)

	for _, p := range data.ParticipantIDs {
		if *p == aid {
			verified = true
			break
		}
	}

	if !verified {
		http.Error(w, errors.New("Unverified user").Error(), http.StatusBadRequest)
		return
	}

	database.NewConversation(&data)

	js, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// MyConversations godoc
// @Summary List a user's own conversation objects
// @Description List a user's own conversation objects
// @Tags My Conversation
// @ID conversation-new
// @Accept  json
// @Produce  json
// @Success 200 {array} models.ConversationData
// @Failure 400 {object} interface{}
// @Router /conversations [post]
// @Security Basic
func MyConversations(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error

	id = GetAuthID(r)

	data, err := database.MyConversations(id)

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

// GetMessage godoc
// @Summary Get message object
// @Description Get message object
// @Tags Message
// @ID message-get
// @Accept  json
// @Produce  json
// @Param messageID path int true "Message ID"
// @Param conversationID path int true "Conversation ID"
// @Success 200 {object} models.MessageData
// @Failure 400 {object} interface{}
// @Router /conversations/{conversationID}/{messageID} [get]
// @Security Basic
func GetMessage(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	q := chi.URLParam(r, "messageID")

	id, err = strconv.ParseInt(q, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mid := GetAuthID(r)

	data, err := database.GetMessage(id, mid)

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

// UpdateMessage godoc
// @Summary Update message object
// @Description Update message object
// @Tags Message
// @ID message-update
// @Accept  json
// @Produce  json
// @Param messageID path int true "Message ID"
// @Param conversationID path int true "Conversation ID"
// @Param message body int true "Message"
// @Success 200 {object} models.MessageData
// @Failure 400 {object} interface{}
// @Router /conversations/{conversationID}/{messageID} [put]
// @Security Basic
func UpdateMessage(w http.ResponseWriter, r *http.Request) {
	var data models.Message

	q := chi.URLParam(r, "messageID")

	id, err := strconv.ParseInt(q, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mid := GetAuthID(r)

	data.ID = &id

	data.OwnerID = &mid

	err = database.UpdateMessage(data)

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

// DeleteMessage godoc
// @Summary Delete message object
// @Description Delete message object
// @Tags Message
// @ID message-delete
// @Accept  json
// @Produce  json
// @Param messageID path int true "Message ID"
// @Param conversationID path int true "Conversation ID"
// @Success 200 {object} models.JsonResultMessage
// @Failure 400 {object} interface{}
// @Router /conversations/{conversationID}/{messageID} [delete]
// @Security Basic
func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error

	q := chi.URLParam(r, "messageID")

	id, err = strconv.ParseInt(q, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mid := GetAuthID(r)

	err = database.DeleteMessage(id, mid)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	js, err := json.Marshal(struct {
		Message string
	}{
		Message: "Success",
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

// NewMessage godoc
// @Summary Delete message object
// @Description Delete message object
// @Tags Message
// @ID message-new
// @Accept  json
// @Produce  json
// @Param message body models.MessageData true "Message"
// @Success 200 {object} models.MessageData
// @Failure 400 {object} interface{}
// @Router /conversations [post]
// @Security Basic
func NewMessage(w http.ResponseWriter, r *http.Request) {
	var data models.Message

	q := chi.URLParam(r, "conversationID")

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

	data.OwnerID = &mid
	data.ConversationID = &id

	verified, err := database.CheckConversation(*data.ConversationID, mid)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if verified {
		database.NewMessage(&data)

		js, err := json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}

	//check if owner_id in participants of public.conversation where id = $1
}

func ListMessages(w http.ResponseWriter, r *http.Request) {
	var id int64
	var err error
	q := chi.URLParam(r, "conversationID")

	id, err = strconv.ParseInt(q, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	verified, err := database.CheckConversation(id, GetAuthID(r))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if verified {
		data, err := database.ListMessages(id)

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

}

func MessageCtx() {
	//check if user.id in conversation.participants
}
