package database

import (
	"deenya-api/models"
	"fmt"

	"github.com/lib/pq"
)

func CheckConversation(id int64, uid int64) (bool, error) {
	var check bool

	q := `SELECT EXISTS(SELECT 1 FROM public.conversation WHERE $1 IN participant AND id = $2 LIMIT 1)`

	row := db.QueryRow(q, uid, id)

	err := row.Scan(&check)

	if err != nil {
		return false, err
	}

	//research optimum way or alternative

	return check, nil
}

func GetConversation(id int64, uid int64) (models.Conversation, error) {
	var data models.Conversation
	q := `SELECT id, participant_id, created_at FROM public.conversation WHERE id = $1 AND $2 IN participant_id` //AND $2 IN participant` //uid
	row := db.QueryRow(q, id, uid)
	err := row.Scan(&data.ID, pq.Array(&data.ParticipantIDs), &data.CreatedAt)
	if err != nil {
		return data, err
	}
	for _, pid := range data.ParticipantIDs {
		var user models.User
		q2 := `SELECT * FROM public.user_detail WHERE user_id = $1`
		err := db.Get(&user, q2, *pid)

		if err != nil {
			return data, err
		}
		data.Participants = append(data.Participants, &user)
	}

	q3 := `SELECT id, conversation_id, media_id, content, created_at, updated_at, owner_id, read_at FROM public.conversation_message WHERE conversation_id = $1`
	rows, err := db.Query(q3, id)

	for rows.Next() {
		var message models.Message
		err := row.Scan(&message.ID, &message.ConversationID, pq.Array(&message.MediaIDs), &message.Content, &message.CreatedAt, &message.OwnerID, &message.ReadAt)
		if err != nil {
			return data, err
		}
		data.Messages = append(data.Messages, &message)
	}

	return data, err
}

func UpdateConversation(data models.Conversation, uid int64) error {
	var err error
	umap := StructMap(data)
	uquery, err := UpString(umap)
	if err != nil {
		return err
	}
	q := fmt.Sprintf("UPDATE public.conversation "+uquery+" WHERE id = :id AND %d IN participant_id", uid)

	_, err = db.NamedQuery(q, data)
	if err != nil {
		return err
	}
	return nil
}

func DeleteConversation(id int64, uid int64) error {
	q := `DELETE FROM public.conversation WHERE id = $1 AND $2 IN participant_id`

	_, err := db.Exec(q, id, uid)

	if err != nil {
		return err
	}

	return nil
}

func NewConversation(new *models.Conversation) error {
	_, csv, csvc := PrepareInsert(new)
	sql := "INSERT INTO public.conversation" + " (" + csv + ") VALUES (" + csvc + ") RETURNING id, participant_id, created_at"

	row, err := db.NamedQuery(sql, new)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if row.Next() {
		err := row.Scan(&new.ID, &new.ParticipantIDs, &new.CreatedAt)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return err
}

func MyConversations(uid int64) ([]models.Conversation, error) {
	var list []models.Conversation

	q := `SELECT id, participant_id, created_at FROM public.conversation WHERE $1 in participant_id`
	rows, err := db.Query(q, uid)
	if err != nil {
		fmt.Println(err)
		return list, err
	}
	for rows.Next() {
		var data models.Conversation
		rows.Scan(&data.ID, &data.ParticipantIDs, &data.CreatedAt)
		list = append(list, data)
	}

	for _, conversation := range list {
		q2 := `SELECT id, conversation_id, media_id, content, created_at, updated_at, owner_id, read_at FROM public.conversation_message WHERE conversation_id = $1 ORDER BY created_at DESC LIMIT 1`
		row := db.QueryRow(q2, *conversation.ID)
		err := row.Scan(&conversation.LatestMessage.ID, &conversation.LatestMessage.ConversationID, pq.Array(&conversation.LatestMessage.MediaIDs), &conversation.LatestMessage.Content, &conversation.LatestMessage.CreatedAt, &conversation.LatestMessage.OwnerID, &conversation.LatestMessage.ReadAt)

		if err != nil {
			return list, err
		}
	}

	return list, nil
}

func GetMessage(id int64, uid int64) (models.Message, error) {
	var message models.Message
	q := `SELECT id, conversation_id, media_id, content, created_at, updated_at, owner_id, read_at FROM public.conversation_message WHERE id = $1 and owner_id = $2`
	row := db.QueryRow(q, id, uid)
	err := row.Scan(&message.ID, &message.ConversationID, pq.Array(&message.MediaIDs), &message.Content, &message.CreatedAt, &message.OwnerID, &message.ReadAt)

	if err != nil {
		return message, err
	}

	//data.MediaID = nil

	return message, err
}

func UpdateMessage(data models.Message) error {
	var err error
	umap := StructMap(data)
	uquery, err := UpString(umap)
	if err != nil {
		return err
	}
	q := "UPDATE public.conversation_message " + uquery + " WHERE id = :id and owner_id = :owner_id"

	_, err = db.NamedQuery(q, data)
	if err != nil {
		return err
	}
	return nil
}

func DeleteMessage(id int64, mid int64) error {
	var media []*int64
	q := `DELETE FROM public.conversation_message WHERE id = $1 AND owner_id = $2 RETURNING media_id`

	row := db.QueryRow(q, id, mid)

	err := row.Scan(pq.Array(&media))

	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, media := range media {
		err := DeleteMedia(*media)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}

func NewMessage(data *models.Message) error {
	_, csv, csvc := PrepareInsert(data)
	sql := "INSERT INTO public.conversation_message" + " (" + csv + ") VALUES (" + csvc + ") RETURNING id, conversation_id, media_id, content, created_at, updated_at, owner_id, read_at"

	row, err := db.NamedQuery(sql, data)
	if err != nil {
		fmt.Println(err)
	}
	if row.Next() {
		err := row.Scan(&data.ID, &data.ConversationID, pq.Array(&data.MediaIDs), &data.Content, &data.CreatedAt, &data.OwnerID, &data.ReadAt)
		if err != nil {
			fmt.Println(err)
		}
	}

	return err
}

func ListMessages(id int64) ([]models.Message, error) {
	var list []models.Message
	q := `SELECT id, conversation_id, media_id, content, created_at, updated_at, owner_id, read_at FROM public.conversation_message WHERE conversation_id = $1`
	rows, err := db.Query(q, id)
	if err != nil {
		fmt.Println(err)
		return list, err
	}
	for rows.Next() {
		var message models.Message
		err := rows.Scan(&message.ID, &message.ConversationID, pq.Array(&message.MediaIDs), &message.Content, &message.CreatedAt, &message.OwnerID, &message.ReadAt)
		list = append(list, message)
		if err != nil {
			fmt.Println(err)
			return list, err
		}
	}
	return list, err
}

// func ListConversationsClient() {
// var data []models.Conversation

// if *user.Type == "client" {

// }

// if *user.Type == "consultant" {

// }
// }
