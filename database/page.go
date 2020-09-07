package database

import (
	"deenya-api/models"
	"fmt"
)

func GetPage(id int64) (models.Page, error) {
	var data models.Page
	q := `SELECT * FROM public.page WHERE id = $1`
	err := db.Get(&data, q, id)
	if err != nil {
		fmt.Println(err)
	}
	return data, err
}

func UpdatePage(data models.Page) error {
	var err error
	umap := StructMap(data)
	uquery, err := UpString(umap)
	if err != nil {
		return err
	}
	q := "UPDATE public.page " + uquery + " WHERE id = :id"

	_, err = db.NamedQuery(q, data)
	if err != nil {
		return err
	}
	return nil
}

func DeletePage(id int64) error {
	q := `DELETE FROM public.page WHERE id = $1`

	_, err := db.Exec(q, id)
	if err != nil {
		return err
	}

	return nil
}

func NewPage(new *models.Page) error {
	var id int64
	_, csv, csvc := PrepareInsert(*new)
	sql := "INSERT INTO public.page" + " (" + csv + ") VALUES (" + csvc + ") RETURNING id"

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
