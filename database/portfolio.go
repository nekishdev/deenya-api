package database

import (
	"deenya-api/models"
	"fmt"
)

func GetPortfolio(id int64) (models.Portfolio, error) {
	var data models.Portfolio
	q := `SELECT * FROM public.portfolio WHERE id = $1`
	err := db.Get(&data, q, id)
	if err != nil {
		fmt.Println(err)
	}
	return data, err
}

func UpdatePortfolio(data models.Portfolio, utype string) error {
	var err error
	umap := StructMap(data)
	uquery, err := UpString(umap)
	if err != nil {
		return err
	}
	q := fmt.Sprintf("UPDATE public.portfolio "+uquery+" WHERE id = :id AND %s_id = :%s_id", utype, utype)

	_, err = db.NamedQuery(q, data)
	if err != nil {
		return err
	}
	return nil
}

func DeletePortfolio(id int64, uid int64, utype string) error {
	q := fmt.Sprintf(`DELETE FROM public.portfolio WHERE id = $1 AND %s_id = $2`, utype)

	_, err := db.Exec(q, id, uid)
	if err != nil {
		return err
	}

	return nil
}

func NewPortfolio(new *models.Portfolio) error {
	_, csv, csvc := PrepareInsert(*new)
	sql := "INSERT INTO public.portfolio" + " (" + csv + ") VALUES (" + csvc + ") RETURNING *"

	row, err := db.NamedQuery(sql, new)
	if err != nil {
		fmt.Println(err)
	}
	if row.Next() {
		row.Scan(&new)
	}

	return err
}

func UserPortfolios(uid int64) ([]models.Portfolio, error) {
	var data []models.Portfolio
	q := `SELECT * FROM public.portfolio WHERE (consultant_id = $1 OR client_id = $1) AND is_published = true` // OR client_id = $1`
	err := db.Select(&data, q, uid)
	if err != nil {
		fmt.Println(err)
	}
	return data, err
}

func MyPortfolios(mid int64, mtype string) ([]models.Portfolio, error) {
	var data []models.Portfolio
	q := fmt.Sprintf(`SELECT * FROM public.portfolio WHERE %s_id = $1`, mtype) // OR client_id = $1`
	err := db.Select(&data, q, mid)
	if err != nil {
		fmt.Println(err)
	}
	return data, err
}

func VerifyPortfolio(id int64, uid int64) bool {
	var verified bool

	q := `SELECT EXISTS(SELECT 1 FROM public.portfolio WHERE id = $1 AND (consultant_id = $2 OR client_id = $2) LIMIT 1)`

	err := db.Get(&verified, q, id, uid)

	if err != nil {
		return false
	}

	return verified
}
