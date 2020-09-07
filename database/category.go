package database

import (
	"deenya-api/models"
	"fmt"
)

func ListCategories() ([]models.Category, error) {
	var list []models.Category
	q := `SELECT * FROM public.categories`
	err := db.Select(&list, q)
	if err != nil {
		fmt.Println(err)
		return list, err
	}

	return list, nil
}

func GetCategory(id int64) (models.Category, error) {
	var data models.Category
	q := `SELECT * FROM public.categories WHERE id = $1`

	err := db.Get(&data, q, id)
	if err != nil {
		fmt.Println(err)
		return data, err
	}

	return data, nil

}
