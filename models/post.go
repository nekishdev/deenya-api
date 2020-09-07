package models

import "time"

type PostThreadTest struct {
	ID      *int64  `db:"id" json:"id,omitempty"`
	OwnerID *int64  `db:"owner_id" json:"owner_id,omitempty"`
	Title   *string `db:"title" json:"title,omitempty"`
	//Excerpt *string  `db:"excerpt" json:"excerpt,omitempty"`
	Content *string   `db:"content" json:"content,omitempty"` //render in html - media src embedded
	Tags    []*string `db:"tags" json:"tags,omitempty"`       //instead of tags
	Views   *int64    `db:"views" json:"view,omitempty"`
	Points  *int64    `db:"points" json:"points,omitempty"` //instead of likes

	IsPublished *bool      `db:"is_published" json:"is_published,omitempty"`
	IsPrivate   *bool      `db:"is_private" json:"is_private,omitempty"` //not publicly available in forum, only user blog
	CreatedAt   *int64     `db:"created_at" json:"created_at,omitempty"`
	ScheduledAt *time.Time `db:"scheduled_at" json:"scheduled_at,omitempty"`

	Owner    *User         `json:"owner,omitempty"`
	Comments []PostComment `json:"comments,omitempty"`
}

type PostData struct {
	ID          *int64     `db:"id" json:"id,omitempty"`
	OwnerID     *int64     `db:"owner_id" json:"owner_id,omitempty"`
	Title       *string    `db:"title" json:"title,omitempty"`
	Excerpt     *string    `db:"excerpt" json:"excerpt,omitempty"`
	Content     *string    `db:"content" json:"content,omitempty"` //render in html - media src embedded
	Tags        []*string  `db:"tags" json:"tags,omitempty"`       //instead of tags
	Views       *int64     `db:"views" json:"views,omitempty"`
	Likes       *int64     `db:"likes" json:"likes,omitempty"` //instead of likes
	IsPublished *bool      `db:"is_published" json:"is_published,omitempty"`
	CreatedAt   *time.Time `db:"created_at" json:"created_at,omitempty"`
	ScheduledAt *time.Time `db:"scheduled_at" json:"scheduled_at,omitempty"`
	Medias      []*Media   `json:"medias,omitempty"` //store array of media src or array of media ids
	MediaIDs    []*int64   `db:"media_ids" json:"-"`
}

type Post struct {
	PostData

	Owner    *User          `json:"owner,omitempty"`
	Comments []*PostComment `json:"comments,omitempty"`
	//Medias   []Media
}

type PostCommentData struct {
	ID        *int64  `db:"id" json:"id,omitempty"`
	OwnerID   *int64  `db:"owner_id" json:"owner_id,omitempty"`
	PostID    *int64  `db:"thread_id" json:"thread_id,omitempty"`
	Content   *string `db:"content" json:"content,omitempty"`
	Points    *int64
	CreatedAt *time.Time `db:"created_at" json:"created_at,omitempty"`
}

type PostComment struct {
	PostCommentData

	Owner *User `json:"owner,omitempty"`
}
