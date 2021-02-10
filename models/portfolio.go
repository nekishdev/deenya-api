package models

import "time"

type PortfolioData struct {
	ID           *int64     `db:"id" json:"id,omitempty" readonly:"true"`
	ConsultantID *int64     `db:"consultant_id" json:"consulant_id,omitempty"`
	ClientID     *int64     `db:"client_id" json:"client_id"`
	BeforeID     *int64     `db:"before_media_id" json:"before_id,omitempty"`
	AfterID      *int64     `db:"after_media_id" json:"after_id,omitempty"`
	CreatedAt    *time.Time `db:"created_at" json:"created_at,omitempty" readonly:"true"`
	Caption      *string    `db:"caption" json:"caption,omitempty"`
	IsPublished  *bool      `db:"is_published" json:"is_published,omitempty"`
	Views        *int64     `db:"views" json:"views,omitempty"`
	Likes        *int64     `db:"likes" json:"likes,omitempty"`
}

type Portfolio struct {
	PortfolioData

	Client     *User  `json:"client,omitempty"`
	Consultant *User  `json:"consultant,omitempty"`
	Before     *Media `json:"before_media,omitempty"` //`emb:"id"` //emb tag to scan into struct.ID
	After      *Media `json:"after_media,omitempty"`
}
