package models

type Group struct {
	ID           *int64  `db:"id" json:"id,omitempty"`
	ConsultantID *int64  `db:"consultant_id" json:"consultant_id,omitempty"`
	Name         *string `db:"name" json:"name,omitempty"`

	Clients []*User `json:"clients,omitempty"`
}
