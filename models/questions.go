package models

import "time"

type QuestionnaireData struct {
	ID           *int64     `db:"id" json:"id,omitempty" readonly:"true"`
	ConsultantID *int64     `db:"consultant_id" json:"consultant_id,omitempty"`
	ClientID     *int64     `db:"client_id" json:"client_id,omitempty"`
	BookingID    *int64     `db:"booking_id" json:"booking_id"`
	Name         *string    `db:"name" json:"name,omitempty"`
	CreatedAt    *time.Time `db:"created_at" json:"created_at,omitempty" readonly:"true"`
	UpdatedAt    *time.Time `db:"updated_at" json:"updated_at,omitempty" readonly:"true"`
}

type Questionnaire struct {
	QuestionnaireData

	Client     *User       `json:"client,omitempty"`
	Consultant *User       `json:"consultant,omitempty"`
	Questions  []*Question `json:"questions,omitempty"`
}

type QuestionData struct {
	ID              *int64     `db:"id" json:"id,omitempty" readonly:"true"`
	QuestionnaireID *int64     `db:"questionnaire_id" json:"questionnaire_id,omitempty"` //ID of questionnaire it belongs to
	Question        *string    `db:"question" json:"question,omitempty"`                 //array or delimiter
	Answer          *string    `db:"answer" json:"answer,omitempty"`                     //array in case of checklist, multiple answers. Or just use a delimiter such as comma and process string depending on type?
	Type            *string    `db:"type" json:"type,omitempty"`                         //multi, single, small_text, large_text
	Choices         []*string  `db:"choices" json:"choices,omitempty"`                   //nullable
	AskedAt         *time.Time `db:"asked_at" json:"asked_at,omitempty"`
	AnsweredAt      *time.Time `db:"answered_at" json:"answered_at,omitempty"`
	OwnerID         *int64     `db:"owner_id" json:"owner_id,omitempty"`
}

type Question struct {
	QuestionData
}
