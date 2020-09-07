package database

import (
	"deenya-api/models"
	"fmt"

	"github.com/lib/pq"
)

func GetPost(id int64) (models.Post, error) {
	var data models.Post
	q := `SELECT id, owner_id, title, excerpt, content, tags, views, likes, is_published, created_at, scheduled_at FROM public.post WHERE id = $1`
	// q := `SELECT * FROM public.post WHERE id = $1`
	//ARRAYS don't work when directly scanned into, need to use pq.Array or a new solution
	// err := db.Get(&data, q, id)
	row := db.QueryRow(q, id)
	err := row.Scan(&data.ID, &data.OwnerID, &data.Title, &data.Excerpt, &data.Content, pq.Array(&data.Tags), &data.Views, &data.Likes, &data.IsPublished, &data.CreatedAt, &data.ScheduledAt)
	if err != nil {
		fmt.Println(err)
	}

	return data, err
}

func GetPostWithOwner(id int64) (models.Post, error) {
	var data models.Post
	q := `SELECT public.post.id, owner_id, title, excerpt, content, tags, views, likes, is_published, public.post.created_at, scheduled_at, public.user.username, public.user.id, public.user.created_at, public.user.type FROM public.post LEFT JOIN public.user ON public.post.owner_id = public.user.id WHERE public.post.id = $1`
	// q := `SELECT * FROM public.post WHERE id = $1`
	//ARRAYS don't work when directly scanned into, need to use pq.Array or a new solution
	// err := db.Get(&data, q, id)
	row := db.QueryRow(q, id)
	var owner models.User
	data.Owner = &owner
	err := row.Scan(&data.ID, &data.OwnerID, &data.Title, &data.Excerpt, &data.Content, pq.Array(&data.Tags), &data.Views, &data.Likes, &data.IsPublished, &data.CreatedAt, &data.ScheduledAt, &data.Owner.Username, &data.Owner.ID, &data.Owner.CreatedAt, &data.Owner.Type)
	if err != nil {
		fmt.Println(err)
	}

	return data, err
}

func UpdatePost(data models.Post) error {
	var err error
	umap := StructMap(data)
	uquery, err := UpString(umap)
	if err != nil {
		return err
	}
	q := "UPDATE public.post " + uquery + " WHERE id = :id"

	_, err = db.NamedQuery(q, data)
	if err != nil {
		return err
	}
	return nil
}

func DeletePost(id int64) error {
	q := `DELETE FROM public.post WHERE id = $1`

	_, err := db.Exec(q, id)
	if err != nil {
		return err
	}

	return nil
}

func NewPost(data *models.Post) error {
	//var id int64

	_, csv, csvc := PrepareInsert(data.PostData)
	sql := "INSERT INTO public.post" + " (" + csv + ") VALUES (" + csvc + ") RETURNING id, owner_id, title, excerpt, content, tags, views, likes, is_published, created_at, scheduled_at"
	fmt.Println(sql)
	row, err := db.NamedQuery(sql, data)
	if err != nil {
		fmt.Println(err)
	}
	if row.Next() {
		row.Scan(&data.ID, &data.OwnerID, &data.Title, &data.Excerpt, &data.Content, pq.Array(&data.Tags), &data.Views, &data.Likes, &data.IsPublished, &data.CreatedAt, &data.ScheduledAt)
	}

	//new.ID = &id
	return err
}

func UserPosts(uid int64) ([]models.Post, error) {
	var list []models.Post
	q := `SELECT id, owner_id, title, excerpt, content, tags, views, likes, is_published, created_at, scheduled_at FROM public.post WHERE owner_id = $1 AND is_published = true`
	rows, err := db.Query(q, uid)
	if err != nil {
		fmt.Println(err)
		return list, err
	}
	for rows.Next() {
		var data models.Post
		rows.Scan(&data.ID, &data.OwnerID, &data.Title, &data.Excerpt, &data.Content, pq.Array(&data.Tags), &data.Views, &data.Likes, &data.IsPublished, &data.CreatedAt, &data.ScheduledAt)
		list = append(list, data)
	}
	return list, err
}

func MyPosts(uid int64) ([]models.Post, error) {
	var list []models.Post
	q := `SELECT id, owner_id, title, excerpt, content, tags, views, likes, is_published, created_at, scheduled_at FROM public.post WHERE owner_id = $1`
	// err := db.Get(&data, q, uid)
	rows, err := db.Query(q, uid)
	if err != nil {
		fmt.Println(err)
		return list, err
	}
	for rows.Next() {
		var data models.Post
		rows.Scan(&data.ID, &data.OwnerID, &data.Title, &data.Excerpt, &data.Content, pq.Array(&data.Tags), &data.Views, &data.Likes, &data.IsPublished, &data.CreatedAt, &data.ScheduledAt)
		list = append(list, data)
	}

	return list, err
}
