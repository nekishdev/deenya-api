package models

import "time"

type ForumThreadData struct {
	ID              *int64     `db:"id" json:"id,omitempty" readonly:"true"`
	OwnerID         *int64     `db:"owner_id" json:"owner_id,omitempty"`
	MainForumPostID *int64     `db:"main_post_id" json:"main_post_id,omitempty"`
	Title           *string    `db:"title" json:"title,omitempty"`
	CreatedAt       *time.Time `db:"created_at" json:"created_at,omitempty" readonly:"true"`
	Tags            []*string  `db:"tags" json:"tags.omitempty"`
}

type ForumThread struct {
	ForumThreadData

	MainForumPost *ForumPost   `json:"main_post,omitempty"`
	ForumPosts    []*ForumPost `json:"forum_posts,omitempty"`
	Owner         *User        `json:"owner,omitempty"`
}

type ForumPostData struct {
	ID        *int64     `db:"id" json:"id,omitempty" readonly:"true"`
	OwnerID   *int64     `db:"owner_id" json:"owner_id,omitempty"`
	ThreadID  *int64     `db:"thread_id" json:"thread_id,omitempty"`
	MediaIDs  []*int64   `db:"media_ids" json:"media_ids,omitempty"`
	Content   *string    `db:"content" json:"content,omitempty"`
	Points    *int64     `db:"points" json:"points,omitempty"`
	CreatedAt *time.Time `db:"created_at" json:"created_at,omitempty" readonly:"true"`
}

type ForumPost struct {
	ForumPostData
	Medias []*Media `json:"medias,omitempty"`
	Owner  *User    `json:"owner,omitempty"`
}

// type ForumTag struct {
// 	ID     *int64
// 	Name   *string
// 	Colour *string
// }
