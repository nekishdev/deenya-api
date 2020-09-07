package database

import (
	"deenya-api/models"
	"fmt"

	"github.com/lib/pq"
)

func UserByUsername(username string) (models.User, error) {
	var user models.User
	q := `SELECT * FROM public.user WHERE username = $1`

	err := db.Get(&user, q, username)

	if err != nil {
		fmt.Println(err)
	}

	return user, err
}

func UserByEmail(email string) (models.User, error) {
	var user models.User
	q := `SELECT * FROM public.user WHERE email = $1`

	err := db.Get(&user, q, email)

	if err != nil {
		fmt.Println(err)
	}

	return user, err
}

func NewUser(new *models.User) error {
	// var id int64
	_, bcsv, bcsvc := PrepareInsert(new.UserBase)

	sql := "INSERT INTO public.user" + " (" + bcsv + ") VALUES (" + bcsvc + ") RETURNING *"
	//fmt.Println(sql)
	row, err := db.NamedQuery(sql, new.UserBase)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if row.Next() {
		err := row.StructScan(&new.UserBase)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	//fmt.Println(id)
	// new.ID = &id
	new.UserID = new.ID
	_, dcsv, dcsvc := PrepareInsert(new.UserDetails)
	//fmt.Println(dcsv, dcsvc)
	sql = fmt.Sprintf("INSERT INTO public.user_detail" + " (" + dcsv + ") VALUES (" + dcsvc + ")")
	fmt.Println(sql)
	_, err = db.NamedExec(sql, new.UserDetails)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if *new.Type == "consultant" {
		new.ConsultantID = new.ID
		_, ccsv, ccsvc := PrepareInsert(new.UserConsultant)
		sql = fmt.Sprintf("INSERT INTO public.user_consultant" + " (" + ccsv + ") VALUES (" + ccsvc + ")")
		fmt.Println(sql)
		_, err = db.NamedExec(sql, new.UserConsultant)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}

func GetUser(id int64) (models.User, error) {
	var test models.User
	q := `SELECT * FROM public.user WHERE id = $1`
	err := db.Get(&test, q, id)
	if err != nil {
		fmt.Println(err)
	}
	return test, err
}

func UpdateUser(data models.UserData) error {
	//test this upsert
	var err error
	umap := StructMap(data.UserBase)
	uquery, err := UpString(umap)
	if err != nil {
		return err
	}
	q := "UPDATE public.user " + uquery + " WHERE id = :id"
	//fmt.Println(q)
	_, err = db.NamedQuery(q, data.UserBase)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(id int64) error {
	q := `DELETE FROM public.user WHERE id = $1`

	_, err := db.Exec(q, id)
	if err != nil {
		return err
	}

	//log.Save("Deleted user:" + id)

	return nil
}

func GetUserDetails(uid int64) (models.UserDetails, error) {
	var data models.UserDetails
	q := `SELECT profile_picture_id, first_name, last_name, title, gender, country, city, timezone, tags FROM public.user_detail WHERE user_id = $1`
	row := db.QueryRow(q, uid)
	err := row.Scan(&data.ProfilePictureID, &data.FirstName, &data.LastName, &data.Title, &data.Gender, &data.Country, &data.City, &data.Timezone, pq.Array(&data.Tags))
	if err != nil {
		fmt.Println(err)
		return data, err
	}
	return data, nil
}

func UpdateUserDetails(test models.UserDetails) error {
	var err error
	umap := StructMap(test)
	uquery, err := UpString(umap)
	if err != nil {
		return err
	}
	q := "UPDATE public.user_detail " + uquery + " WHERE user_id = :user_id"
	//fmt.Println(q)
	_, err = db.NamedQuery(q, test)
	if err != nil {
		return err
	}
	return nil
}

func GetUserConsultant(uid int64) (models.UserConsultant, error) {
	var data models.UserConsultant
	q := `SELECT * FROM public.user_consultant`
	err := db.Get(&data, q)
	if err != nil {
		fmt.Println(err)
		return data, err
	}
	return data, nil
}

func UpdateUserConsultant(data models.UserConsultant) (models.UserConsultant, error) {
	umap := StructMap(data)
	uquery, err := UpString(umap)
	if err != nil {
		fmt.Println(err)
		return data, err
	}
	q := "UPDATE public.user_consultant " + uquery + " WHERE user_id = :user_id"
	_, err = db.NamedQuery(q, data)
	if err != nil {
		fmt.Println(err)
		return data, err
	}
	return data, nil
}
