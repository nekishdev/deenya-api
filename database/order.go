package database

import (
	"deenya-api/models"
	"fmt"
)

func GetOrder(id int64, uid int64, utype string) (models.Order, error) {
	var data models.Order
	//SELECT orders as (), * FROM public.order
	q := fmt.Sprintf(`SELECT * FROM public.order WHERE id = $1 AND %s_id = $2`, utype)
	err := db.Get(&data, q, id, uid)
	if err != nil {
		fmt.Println(err)
	}

	return data, err
}

func UpdateOrder(data models.Order, uid int64, utype string) error {
	var err error
	umap := StructMap(data)
	uquery, err := UpString(umap)
	if err != nil {
		return err
	}
	q := fmt.Sprintf("UPDATE public.order "+uquery+" WHERE id = :id AND %s_id = '%d'", utype, uid)

	_, err = db.NamedQuery(q, data)
	if err != nil {
		return err
	}
	return nil
}

func DeleteOrder(id int64) error {
	q := `DELETE FROM public.order WHERE id = $1`

	_, err := db.Exec(q, id)
	if err != nil {
		return err
	}

	return nil
}

func NewOrder(new *models.Order) error {

	_, csv, csvc := PrepareInsert(*new)
	sql := "INSERT INTO public.order" + " (" + csv + ") VALUES (" + csvc + ") RETURNING *"

	row, err := db.NamedQuery(sql, new)
	if err != nil {
		fmt.Println(err)
	}
	if row.Next() {
		row.Scan(&new)
	}

	for _, item := range new.OrderItems {
		err := NewOrderItem(&item)
		if err != nil {
			fmt.Println(err)
			status := "error"
			new.Status = &status
			UpdateOrder(*new, *new.ConsultantID, "consultant")
			return err
		}
	}
	return err
}

func MyOrders(uid int64, utype string) ([]models.Order, error) {
	var data []models.Order
	q := fmt.Sprintf(`SELECT * FROM public.order WHERE %s_id = $1 ORDER BY created_at DESC`, utype)
	err := db.Select(&data, q, uid)
	if err != nil {
		fmt.Println(err)
	}
	return data, err
}

func UserOrders(mid int64, mtype string, tid int64) ([]models.Order, error) {
	var ttype string
	if mtype == "client" {
		ttype = "consultant"
	}
	if mtype == "consultant" {
		ttype = "client"
	}
	var data []models.Order
	q := fmt.Sprintf(`SELECT * FROM public.order WHERE %s_id = $1 AND %s_id = $2 ORDER BY created_at DESC`, mtype, ttype)
	err := db.Select(&data, q, mid, tid)
	if err != nil {
		fmt.Println(err)
	}
	return data, err

}

func NewOrderItem(new *models.OrderItem) error {

	_, csv, csvc := PrepareInsert(*new)
	sql := "INSERT INTO public.order_item" + " (" + csv + ") VALUES (" + csvc + ") RETURNING *"

	row, err := db.NamedQuery(sql, new)
	if err != nil {
		fmt.Println(err)
	}
	if row.Next() {
		row.Scan(&new)
	}

	return err
}

func ListOrderItems(order int64) ([]models.OrderItem, error) {
	var data []models.OrderItem

	q := `SELECT * FROM order_item WHERE order_id = $1`

	err := db.Select(&data, q, order)

	if err != nil {
		fmt.Println(err)
		return data, err
	}

	return data, nil
}
