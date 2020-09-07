package database

import (
	"deenya-api/models"
	"fmt"

	"github.com/lib/pq"
)

func GetTreatment(id int64) (models.Treatment, error) {
	var data models.Treatment
	q := `SELECT id, consultant_id, client_id, order_id, portfolio_id, notes, tags, created_at FROM public.treatment WHERE id = $1`
	row := db.QueryRow(q, id)
	err := row.Scan(&data.ID, &data.ConsultantID, &data.ClientID, &data.OrderID, &data.PortfolioID, &data.Notes, pq.Array(&data.Tags), &data.CreatedAt)
	if err != nil {
		fmt.Println(err)
		return data, nil
	}
	return data, err
}

func UpdateTreatment(data models.Treatment, utype string) error {
	var err error
	umap := StructMap(data)
	uquery, err := UpString(umap)
	if err != nil {
		return err
	}
	q := "UPDATE public.treatment " + uquery + " WHERE id = :id AND consultant_id = :consultant_id"
	_, err = db.NamedQuery(q, data)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTreatment(id int64, uid int64, utype string) error {
	q := fmt.Sprintf(`DELETE FROM public.treatment WHERE id = $1 AND %s_id = $2`, utype)

	_, err := db.Exec(q, id, uid)
	if err != nil {
		return err
	}

	return nil
}

func NewTreatment(data *models.Treatment) error {
	_, csv, csvc := PrepareInsert(*data)
	sql := "INSERT INTO public.treatment" + " (" + csv + ") VALUES (" + csvc + ") RETURNING id, consultant_id, client_id, order_id, portfolio_id, notes, tags, created_at"

	row, err := db.NamedQuery(sql, *data)
	if err != nil {
		fmt.Println(err)
	}
	if row.Next() {
		row.Scan(&data)
	}
	return err
}

func UserTreatments(uid int64, mid int64) ([]models.Treatment, error) {
	var list []models.Treatment
	q := `SELECT id, consultant_id, client_id, order_id, portfolio_id, notes, tags, created_at FROM public.treatment WHERE consultant_id = $1 AND client_id = $2` // OR client_id = $1`
	rows, err := db.Query(q, uid, mid)
	if err != nil {
		fmt.Println(err)
		return list, err
	}
	for rows.Next() {
		var data models.Treatment
		err = rows.Scan(&data.ID, &data.ConsultantID, &data.ClientID, &data.OrderID, &data.PortfolioID, &data.Notes, pq.Array(&data.Tags), &data.CreatedAt)
		if err != nil {
			fmt.Println(err)
			return list, nil
		}
		list = append(list, data)
	}
	return list, err
}

func MyTreatments(mid int64, mtype string) ([]models.Treatment, error) {
	var list []models.Treatment
	q := fmt.Sprintf(`SELECT id, consultant_id, client_id, order_id, portfolio_id, notes, tags, created_at FROM public.treatment WHERE %s_id = $1`, mtype) // OR client_id = $1`
	rows, err := db.Query(q, mid)
	if err != nil {
		fmt.Println(err)
		return list, err
	}
	for rows.Next() {
		var data models.Treatment
		err = rows.Scan(&data.ID, &data.ConsultantID, &data.ClientID, &data.OrderID, &data.PortfolioID, &data.Notes, pq.Array(&data.Tags), &data.CreatedAt)
		if err != nil {
			fmt.Println(err)
			return list, nil
		}
		list = append(list, data)
	}
	return list, err
}
