package models

import "time"

type Page struct {
	ID               *int64     `db:"id" json:"id,omitempty" readonly:"true"`
	Title            *string    `db:"title" json:"title,omitempty"`
	Description      *string    `db:"description" json:"description,omitempty"`
	FeaturedProducts []*int64   `db:"featured_product" json:"featured_product,omitempty"`
	FeaturedPosts    []*int64   `db:"featured_post" json:"featured_post,omitempty"`
	Tags             []*string  `db:"tags" json:"tags,omitempty"`
	UpdatedAt        *time.Time `db:"updated_at" json:"updated_at,omitempty" readonly:"true"`
}
