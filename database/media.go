package database

import (
	"deenya-api/models"
	"fmt"
	"strconv"
)

func GetMedia(id int64) (models.Media, error) {
	var data models.Media
	q := `SELECT * FROM public.media WHERE id = $1`
	err := db.Get(&data, q, id)
	if err != nil {
		fmt.Println(err)
	}

	return data, err
}

func UpdateMedia(data models.Media) error {
	var err error
	umap := StructMap(data)
	uquery, err := UpString(umap)
	if err != nil {
		return err
	}
	q := "UPDATE public.media " + uquery + " WHERE id = :id"

	_, err = db.NamedQuery(q, data)
	if err != nil {
		return err
	}
	return nil
}

func DeleteMedia(id int64) error {
	var media []*string
	q := `DELETE FROM public.media WHERE id = $1 RETURNING media`

	row := db.QueryRow(q, id)

	err := row.Scan(&media)

	var mediaArr string

	//delete corresponding media from array returned

	q2 := `DELETE FROM public.media WHERE id IN $1`

	_, err = db.Exec(q2, mediaArr)
	if err != nil {
		return err
	}

	return nil
}

func NewMedia(new *models.Media) error {

	_, csv, csvc := PrepareInsert(*new)
	sql := "INSERT INTO public.media" + " (" + csv + ") VALUES (" + csvc + ") RETURNING *"

	row, err := db.NamedQuery(sql, new)
	if err != nil {
		fmt.Println(err)
	}
	if row.Next() {
		row.Scan(&new)
	}

	return err
}

func ListMedias(uid int64) ([]models.Media, error) {
	var arr string
	var data []models.Media
	// for i, pid := range order.ItemID {
	// 	if i == 0 {
	// 		arr = "ARRAY[" + "'" + strconv.FormatInt(*pid, 10) + "'"
	// 		continue
	// 	}

	// 	arr = arr + ", '" + strconv.FormatInt(*pid, 10) + "'"
	// }
	// arr = arr + "]"

	q := `SELECT * FROM media WHERE owner_id = $1`

	err := db.Select(&data, q, arr)

	if err != nil {
		fmt.Println(err)
		return data, err
	}

	return data, nil
}

func MyMedia(uid int64) ([]models.Media, error) {
	var arr string
	var data []models.Media

	q := `SELECT * FROM media WHERE owner_id = $1`

	err := db.Select(&data, q, arr)

	if err != nil {
		fmt.Println(err)
		return data, err
	}

	return data, nil
}

func ArrayMedia(ids []*int64) ([]models.Media, error) {
	var arr string
	var medias []models.Media
	for i, pid := range ids {
		if i == 0 {
			arr = "ARRAY[" + "'" + strconv.FormatInt(*pid, 10) + "'"
			continue
		}

		arr = arr + ", '" + strconv.FormatInt(*pid, 10) + "'"
	}

	arr = arr + "]"
	//SELECT * FROM test_items where id in (select unnest(items) from test where id = '27')
	q := `SELECT * FROM public.media WHERE id IN $1`

	err := db.Get(&medias, q, arr)

	if err != nil {
		fmt.Println(err)
		return medias, err
	}

	//decide how to render on frontend

	return medias, nil
}
