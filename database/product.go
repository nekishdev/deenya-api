package database

import (
	"deenya-api/models"
	"fmt"

	"github.com/lib/pq"
)

func GetProduct(id int64) (models.Product, error) {
	var data models.Product
	q := `SELECT * FROM public.product WHERE id = $1`
	err := db.Get(&data, q, id)
	if err != nil {
		fmt.Println(err)
	}

	return data, err
}

func UpdateProduct(data models.Product) error {
	var err error
	umap := StructMap(data)
	uquery, err := UpString(umap)
	if err != nil {
		return err
	}
	q := "UPDATE public.product " + uquery + " WHERE id = :id AND owner_id = :owner_id"

	_, err = db.NamedQuery(q, data)
	if err != nil {
		return err
	}
	return nil
}

func DeleteProduct(id int64, uid int64) error {
	var media []*string
	q := `DELETE FROM public.product WHERE id = $1 AND owner_id = $2 RETURNING media`

	row := db.QueryRow(q, id, uid)

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

func NewProduct(data *models.Product) error {

	_, csv, csvc := PrepareInsert(data.ProductData)

	sql := "INSERT INTO public.product" + " (" + csv + ") VALUES (" + csvc + ") RETURNING *"

	row, err := db.NamedQuery(sql, data.ProductData)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if row.Next() {
		err = row.StructScan(&data)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return err
}

func UserProducts(uid int64) ([]models.Product, error) {
	var data []models.Product
	var err error
	q := `SELECT id, owner_id, media_ids, tags, name, price, stock, created_at FROM public.product WHERE owner_id = $1`
	rows, err := db.Queryx(q, uid)
	if err != nil {
		fmt.Println(err)
		return data, err
	}

	for rows.Next() {
		var product models.Product
		err := rows.StructScan(&product)
		if err != nil {
			fmt.Println(err)
			return data, err
		}
		data = append(data, product)
	}

	for _, product := range data {
		var medias []models.Media
		product.Model, err = GetProductModel(*product.ProductModelID)
		if err != nil {
			fmt.Println(err)
			return data, err
		}
		medias, err = ArrayMedia(product.Model.MediaIDs)
		if err != nil {
			fmt.Println(err)
			return data, err
		}

		for _, media := range medias {
			product.Model.Medias = append(product.Model.Medias, &media)
		}
	}

	return data, err
}

func MyProducts(uid int64) ([]models.Product, error) {
	var data []models.Product
	var err error
	q := `SELECT * FROM public.product WHERE owner_id = $1`
	rows, err := db.Queryx(q, uid)
	if err != nil {
		fmt.Println(err)
		return data, err
	}

	for rows.Next() {
		var product models.Product
		err := rows.StructScan(&product)
		if err != nil {
			fmt.Println(err)
			return data, err
		}

		product.Model, err = GetProductModel(*product.ProductModelID)

		data = append(data, product)

	}

	// for _, product := range data {
	// 	product.Medias, err = ListProductMedia(product.MediaIDs)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return data, err
	// 	}
	// }

	return data, err
}

func GetProductReviews(product int64) ([]models.ProductReview, error) {
	var data []models.ProductReview
	q := `SELECT * FROM public.product_review WHERE product_id = $1`

	err := db.Select(&data, q, product)
	if err != nil {
		fmt.Println(err)
		return data, err
	}

	return data, err
}

func GetProductModel(id int64) (models.ProductModel, error) {
	var data models.ProductModel
	q := `SELECT id, media_ids, tags, name FROM public.product_model WHERE id = $1`

	row := db.QueryRow(q, id)

	err := row.Scan(&data.ID, &data.MediaIDs, pq.Array(&data.Tags), &data.Name)
	if err != nil {
		fmt.Println(err)
		return data, err
	}

	return data, nil
}

func SearchProductModels(query string, tags []string) ([]models.ProductModel, error) {
	var list []models.ProductModel
	var arrString string
	if len(tags) > 0 {
		for i, cid := range tags {
			if i == 0 {
				arrString = fmt.Sprintf("[%s", cid)
				continue
			}
			arrString = fmt.Sprintf(arrString+", %d", cid)
		}
		arrString += "]"
	}
	//media_ids or medias
	q := `SELECT id, media_ids, tags, name, price FROM public.product_model`
	if query != "" {
		q = q + fmt.Sprintf(" WHERE (description ILIKE '%s' OR name ILIKE '%s')", query, query)
	}
	if arrString != "" {
		var join string
		if query == "" {

			join = "WHERE"
		}
		if query != "" {
			join = "AND"
		}
		q = q + " " + join + "category_ids IN " + arrString
	}

	rows, err := db.Query(q)
	if err != nil {
		fmt.Println(err)
		return list, err
	}

	for rows.Next() {
		var data models.ProductModel
		err := rows.Scan(&data.ID, &data.MediaIDs, pq.Array(&data.Tags), &data.Name)
		if err != nil {
			fmt.Println(err)
			return list, err
		}
		list = append(list, data)
	}

	return list, nil
}

func SuggestProductModels(categories []string) ([]models.ProductModel, error) {
	var list []models.ProductModel
	var arrString string

	for i, cid := range categories {
		if i == 0 {
			arrString = fmt.Sprintf("[%d", cid)
			continue
		}
		arrString = fmt.Sprintf(arrString+", %d", cid)
	}
	arrString += "]"

	q := `SELECT id, media_ids, category_ids, name, price FROM public.product_model WHERE category_ids IN ` + arrString

	rows, err := db.Query(q)
	if err != nil {
		fmt.Println(err)
		return list, err
	}

	for rows.Next() {
		var data models.ProductModel
		err := rows.Scan(&data.ID, &data.MediaIDs, &data.Tags, &data.Name)
		if err != nil {
			fmt.Println(err)
			return list, err
		}
		list = append(list, data)
	}

	return list, nil
}
