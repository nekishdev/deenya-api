package models

import "time"

type ConversationData struct {
	ID             *int64   `db:"id" json:"id,omitempty" readonly:"true"`
	ParticipantIDs []*int64 `db:"participant_ids" json:"participant_ids,omitempty"`
	//LatestMessageID *int64   `db:"latest_message_id" json:"latest_message_id,omitempty"`
	// ConsultantID *int64 `db:"consultant_id" json:"consultant_id,omitempty"`
	// ClientID     *int64 `db:"client_id" json:"client_id,omitempty"`
	//UpdatedAt
	CreatedAt *time.Time `db:"created_at" json:"created_at,omitempty" readonly:"true"`
}

type Conversation struct {
	ConversationData
	Messages      []*Message `json:"messages,omitempty"`
	Participants  []*User    `json:"participants,omitempty"`
	LatestMessage *Message   `json:"latest_message,omitempty"` //or inner join?
	// Consultant *User      `json:"consultant,omitempty"`
	// Client     *User      `json:"client,omitempty"`
}

type MessageData struct {
	ID             *int64     `db:"id" json:"id,omitempty" readonly:"true"`
	ConversationID *int64     `db:"conversation_id" json:"conversation_id,omitempty"`
	MediaIDs       []*int64   `db:"media_ids" json:"media_ids,omitempty"`
	Content        *string    `db:"content" json:"content,omitempty"`
	CreatedAt      *time.Time `db:"created_at" json:"created_at" readonly:"true"`
	UpdatedAt      *time.Time `db:"updated_at" json:"updated_at" readonly:"true"`
	OwnerID        *int64     `db:"owner_id" json:"owner_id"`
	ReadAt         *time.Time `db:"read_at" json:"read_at"`
}

type Message struct {
	MessageData
	Medias []*Media `json:"medias,omitempty"`
	Owner  *User    `json:"owner,omitempty"`
}
