package models

import "time"

type MediaData struct {
	ID        *int64     `db:"id" json:"id,omitempty" readonly:"true"`
	Src       *string    `db:"src" json:"src,omitempty"`
	CreatedAt *time.Time `db:"created_at" json:"created_at,omitempty" readonly:"true"`
	OwnerID   *int64     `db:"owner_id" json:"owner_id,omitempty"`
}

type Media struct {
	MediaData
	Owner *User `json:"owner"`
}
