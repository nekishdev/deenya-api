package database

import (
	"deenya-api/models"
	"fmt"

	"github.com/lib/pq"
)

func GetForumThread(id int64) (models.ForumThread, error) {
	var data models.ForumThread
	q := `SELECT * FROM public.forum_thread WHERE id = $1`
	err := db.Get(&data, q, id)
	if err != nil {
		fmt.Println(err)
	}
	return data, err
}

func UpdateForumThread(data models.ForumThread) error {
	var err error
	umap := StructMap(data)
	uquery, err := UpString(umap)
	if err != nil {
		return err
	}
	q := "UPDATE public.forum_thread " + uquery + " WHERE id = :id AND owner_id = :owner_id"

	_, err = db.NamedQuery(q, data)
	if err != nil {
		return err
	}
	return nil
}

func DeleteForumThread(id int64, mid int64) error {
	q := `DELETE FROM public.forum_thread WHERE id = $1 AND owner_id = $2`

	_, err := db.Exec(q, id, mid)
	if err != nil {
		return err
	}

	return nil
}

func NewForumThread(new *models.ForumThread) error {
	_, csv, csvc := PrepareInsert(*new)
	sql := "INSERT INTO public.forum_thread" + " (" + csv + ") VALUES (" + csvc + ") RETURNING *"

	row, err := db.NamedQuery(sql, new)
	if err != nil {
		fmt.Println(err)
	}
	if row.Next() {
		row.Scan(&new)
	}
	return err
}

func UserForumThreads(uid int64) ([]models.ForumThread, error) {
	var data []models.ForumThread
	q := `SELECT * FROM public.forum_thread WHERE owner_id = $1`
	err := db.Select(&data, q, uid)
	if err != nil {
		fmt.Println(err)
	}
	return data, err
}

func MyForumThreads(uid int64) ([]models.ForumThread, error) {
	var data []models.ForumThread
	q := `SELECT * FROM public.forum_thread WHERE owner_id = $1`
	err := db.Select(&data, q, uid)
	if err != nil {
		fmt.Println(err)
	}
	return data, err
}

func ForumThreadFeed() ([]models.ForumThread, error) {
	var data []models.ForumThread
	q := `SELECT * FROM public.forum_thread ORDER BY created_at DESC LIMIT 100`
	err := db.Select(&data, q)
	if err != nil {
		fmt.Println(err)
	}
	return data, err
}

func GetForumPost(id int64) (models.ForumPost, error) {
	var data models.ForumPost
	q := `SELECT id, owner_id, thread_id, media_ids, content, points, created_at FROM public.forum_post WHERE id = $1`
	row := db.QueryRow(q, id)
	err := row.Scan(&data.ID, &data.OwnerID, &data.ThreadID, pq.Array(&data.MediaIDs), &data.Content, &data.Points, &data.CreatedAt)
	if err != nil {
		fmt.Println(err)
	}
	return data, err
}

func UpdateForumPost(data models.ForumPost) error {
	var err error
	umap := StructMap(data)
	uquery, err := UpString(umap)
	if err != nil {
		return err
	}
	q := "UPDATE public.forum_post " + uquery + " WHERE id = :id AND owner_id = :owner_id"

	_, err = db.NamedQuery(q, data)
	if err != nil {
		return err
	}
	return nil
}

func DeleteForumPost(id int64, mid int64) error {
	q := `DELETE FROM public.forum_post WHERE id = $1 AND owner_id = $2`

	_, err := db.Exec(q, id, mid)
	if err != nil {
		return err
	}

	return nil
}

func NewForumPost(new *models.ForumPost) error {
	var id int64
	_, csv, csvc := PrepareInsert(*new)
	sql := "INSERT INTO public.forum_post" + " (" + csv + ") VALUES (" + csvc + ") RETURNING id"

	row, err := db.NamedQuery(sql, new)
	if err != nil {
		fmt.Println(err)
	}
	if row.Next() {
		row.Scan(&id)
	}

	new.ID = &id
	return err
}

func ThreadForumPosts(id int64) ([]models.ForumPost, error) {
	var list []models.ForumPost
	q := `SELECT * FROM public.forum_post WHERE thread_id = $1`
	rows, err := db.Query(q, id)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var data models.ForumPost
		err := rows.Scan(&data.ID, &data.OwnerID, &data.ThreadID, pq.Array(&data.MediaIDs), &data.Content, &data.Points, &data.CreatedAt)
		if err != nil {
			fmt.Println(err)
		}
		list = append(list, data)
	}

	return list, err
}

func UserForumPosts(uid int64) ([]models.ForumPost, error) {
	var list []models.ForumPost
	q := `SELECT * FROM public.forum_post WHERE owner_id = $1`
	rows, err := db.Query(q, uid)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var data models.ForumPost
		err := rows.Scan(&data.ID, &data.OwnerID, &data.ThreadID, pq.Array(&data.MediaIDs), &data.Content, &data.Points, &data.CreatedAt)
		if err != nil {
			fmt.Println(err)
		}
		list = append(list, data)
	}

	return list, err
}

func MyForumPosts(uid int64) ([]models.ForumPost, error) {
	var list []models.ForumPost
	q := `SELECT * FROM public.forum_post WHERE owner_id = $1`
	rows, err := db.Query(q, uid)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var data models.ForumPost
		err := rows.Scan(&data.ID, &data.OwnerID, &data.ThreadID, pq.Array(&data.MediaIDs), &data.Content, &data.Points, &data.CreatedAt)
		if err != nil {
			fmt.Println(err)
		}
		list = append(list, data)
	}

	return list, err
}
