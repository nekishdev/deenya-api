package database

import (
	"deenya-api/models"
	"fmt"
)

// func NewTest(new models.Test) (models.Test, error) {
// 	//panics when using nil reference in VALUES, use direct pointer (var not *var)
// 	q := `INSERT INTO test (tester, opt) VALUES ($1, $2) RETURNING id`
// 	// err := db.QueryRow(context.Background(), q, new.Tester, new.Opt).Scan(&new.ID)
// 	err := db.QueryRow(q, new.Tester, new.Opt).Scan(&new.ID)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	row := db.NamedQuery()

// 	return new, err
// }

func NewTest(new *models.Test) error {
	var id int64
	_, csv, csvc := PrepareInsert(*new)
	sql := "INSERT INTO test" + " (" + csv + ") VALUES (" + csvc + ") RETURNING id"
	//fmt.Println(sql)
	row, err := db.NamedQuery(sql, new)
	if err != nil {
		fmt.Println(err)
	}
	if row.Next() {
		row.Scan(&id)
	}
	//fmt.Println(id)
	new.ID = &id
	return err
}

func GetTest(id int64) (models.Test, error) {
	var test models.Test
	q := `SELECT * FROM test WHERE id = $1`
	err := db.Get(&test, q, id)
	if err != nil {
		fmt.Println(err)
	}
	return test, err
}

func UpdateTest(test models.Test) error {
	//test this upsert
	var err error
	umap := StructMap(test)
	uquery, err := UpString(umap)
	if err != nil {
		return err
	}
	q := "UPDATE test " + uquery + " WHERE id = :id"
	//fmt.Println(q)
	_, err = db.NamedQuery(q, test)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTest(id int64) error {
	q := `DELETE FROM test WHERE id = $1`

	res, err := db.Exec(q, id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()

	if err != nil {
		return err
	}

	return nil
}

// func UpsertTest(test models.Test) error {
// 	//test this upsert
// 	var err error
// 	umap := StructMap(test)
// 	fields := DBFields(umap)
// 	csv, csvc := prepFields(fields)
// 	uquery := UpString(umap)
// 	q := "INSERT INTO test" + " (" + csv + ") VALUES (" + csvc + ") ON CONFLICT (id) DO UPDATE " + uquery
// 	fmt.Println(q)
// 	_, err = db.NamedQuery(q, test)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return err
// }
