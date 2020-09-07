package database

import (
	"deenya-api/models"
	"fmt"
)

func GetOrderDelivery(id int64) (models.OrderDelivery, error) {
	var data models.OrderDelivery
	q := `SELECT * FROM public.order_delivery WHERE id = $1`
	err := db.Get(&data, q, id)
	if err != nil {
		fmt.Println(err)
	}
	return data, err
}

func UpdateOrderDelivery(data models.OrderDelivery) error {
	var err error
	umap := StructMap(data)
	uquery, err := UpString(umap)
	if err != nil {
		return err
	}
	q := "UPDATE public.order_delivery " + uquery + " WHERE id = :id"

	_, err = db.NamedQuery(q, data)
	if err != nil {
		return err
	}
	return nil
}

func DeleteOrderDelivery(id int64) error {
	q := `DELETE FROM public.order_delivery WHERE id = $1`

	_, err := db.Exec(q, id)
	if err != nil {
		return err
	}

	return nil
}

func NewOrderDelivery(new *models.OrderDelivery) error {
	_, csv, csvc := PrepareInsert(*new)
	sql := "INSERT INTO public.order_delivery" + " (" + csv + ") VALUES (" + csvc + ") RETURNING *"

	row, err := db.NamedQuery(sql, new)
	if err != nil {
		fmt.Println(err)
	}
	if row.Next() {
		row.Scan(&new)
	}

	return err
}

func ListDeliveries(uid int64) ([]models.OrderDelivery, error) {
	var data []models.OrderDelivery
	q := `SELECT * FROM public.order_delivery WHERE order_id = $1`
	err := db.Select(&data, q, uid)
	if err != nil {
		fmt.Println(err)
	}
	return data, err
}
