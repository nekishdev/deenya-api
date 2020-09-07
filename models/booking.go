package models

import "time"

type BookingData struct {
	ID              *int64 `db:"id" json:"id,omitempty"`
	ConsultantID    *int64 `db:"consultant_id" json:"consultant_id,omitempty"`
	ClientID        *int64 `db:"client_id" json:"client_id,omitempty"`
	InvoiceID       *int64 `db:"invoice_id" json:"invoice_id,omitempty"` //need to check if invoice is null on frontend?
	QuestionnaireID *int64 `db:"questionnaire_id" json:"questionnaire_id,omitempty"`
	TreatmentID     *int64 `db:"treatment_id" json:"treatment_id,omitempty"`
	ConversationID  *int64 `db:"conversation_id" json:"conversation_id,omitempty"`
	//FollowUpID   *int64     `db:"follow_up_id"`
	Inquiry     *string    `db:"inquiry" json:"inquiry,omitempty"`
	Tags        []*string  `db:"tags" json:"tags,omitempty"`
	CreatedAt   *time.Time `db:"created_at" json:"created_at,omitempty"`
	StartedAt   *time.Time `db:"started_at" json:"started_at,omitempty"`
	EndedAt     *time.Time `db:"ended_at" json:"ended_at,omitempty"`
	ScheduledAt *time.Time `db:"scheduled_at" json:"scheduled_at,omitempty"`
	//Elapsed     *int64     `db:"elapsed" json:"elapsed,omitempty"`
	//Status     *string `db:"status" json:"status,omitempty"`
	IsAccepted *bool `db:"is_accepted" json:"is_accepted,omitempty"`
	IsRemote   *bool `db:"is_remote" json:"is_remote,omitempty"`
}

type Booking struct {
	BookingData
	Questionnaire *Questionnaire `json:"questionnaire,omitempty"`
	Treatment     *Treatment     `json:"treatment,omitempty"`
	Client        *User          `json:"client,omitempty"`
	Consultant    *User          `json:"consultant,omitempty"`
	Invoice       *Invoice       `json:"invoice,omitempty"`
	Conversation  *Conversation  `json:"conversation,omitempty"`
	//tags?
	//FollowUp     *Booking
	//FollowedFrom *Booking
}

type BookingSlot struct {
	Date      int
	Time      time.Time
	Available bool
}

//use new invoice table to connect financing for multiple types e.g orders and bookings
