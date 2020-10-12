package database

import (
	"deenya-api/models"
	"errors"
	"fmt"

	"github.com/lib/pq"
)

func MyClinic(mid int64) (models.Clinic, error) {
	var data models.Clinic
	var member models.ClinicMember
	q := `SELECT * FROM public.clinic_member WHERE is_accepted = true AND consultant_id = $1 LIMIT 1`
	err := db.Get(&member, q, mid)
	if err != nil {
		fmt.Println(err)
		return data, err
	}
	q2 := `SELECT id, name, lat, lng, city, country, street, postcode, admins FROM public.clinic WHERE id = $1`
	row := db.QueryRow(q2, *member.ClinicID)
	err = row.Scan(&data.ID, &data.Name, &data.Lat, &data.Lng, &data.City, &data.Country, &data.Street, &data.Postcode, pq.Array(&data.Admins))
	if err != nil {
		fmt.Println(err)
		return data, err
	}
	return data, nil
}

func NewClinic(data *models.Clinic) error {
	_, csv, csvc := PrepareInsert(data.ClinicData)
	sql := "INSERT INTO public.clinic" + " (" + csv + ") VALUES (" + csvc + ") RETURNING id"

	row, err := db.NamedQuery(sql, *data)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if row.Next() {
		row.Scan(&data.ID)
	}
	return err
}

func GetClinic(id int64) (models.Clinic, error) {
	var data models.Clinic
	q := `SELECT id, name, lat, lng, city, country, street, postcode, admins FROM public.clinic WHERE id = $1`
	row := db.QueryRow(q, id)
	err := row.Scan(&data.ID, &data.Name, &data.Lat, &data.Lng, &data.City, &data.Country, &data.Street, &data.Postcode, pq.Array(&data.Admins))
	if err != nil {
		fmt.Println(err)
		return data, err
	}
	return data, nil
}

func UpdateClinic(data models.Clinic, mid int64) error {
	var err error
	umap := StructMap(data)
	uquery, err := UpString(umap)
	if err != nil {
		return err
	}
	q := fmt.Sprintf("UPDATE public.clinic "+uquery+" WHERE id = :id AND %d IN admins", mid)
	_, err = db.NamedQuery(q, data)
	if err != nil {
		return err
	}
	return nil
}

func IsClinicAdmin(id int64, mid int64) bool {
	verified := false
	clinic, err := GetClinic(id)
	if err != nil {
		fmt.Println(err)
		return verified
	}
	for _, admin := range clinic.AdminIds {
		if mid == *admin {
			verified = true
		}
	}
	return verified
}

func AcceptClinicRequest(id int64, mid int64, tid int64) error {
	clinic, err := GetClinic(id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	verified := false
	for _, admin := range clinic.AdminIds {
		if mid == *admin {
			verified = true
		}
	}
	if !verified {
		return errors.New("Unauthorized")
	}

	q := `UPDATE public.clinic_member SET is_accepted = true WHERE clinic_id = $1 AND consultant_id = $2  `

	_, err = db.Exec(q, id, tid)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func NewClinicRequest(data *models.ClinicMember) error {
	_, csv, csvc := PrepareInsert(*data)
	sql := "INSERT INTO public.clinic_member" + " (" + csv + ") VALUES (" + csvc + ") RETURNING *"

	row, err := db.NamedQuery(sql, *data)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if row.Next() {
		row.Scan(&data)
	}
	return nil
}

func RemoveClinicMember(id int64, mid int64, tid int64) error {
	if !IsClinicAdmin(id, mid) {
		return errors.New("Unauthorized")
	}

	q := `DELETE FROM public.clinic_member WHERE id = $1 AND consultant_id = $2`

	_, err := db.Exec(q, id, tid)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil

}

func LeaveClinic(id int64, mid int64) error {
	q := `DELETE FROM public.clinic_member WHERE id = $1 AND consultant_id = $2`

	_, err := db.Exec(q, id, mid)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func ListClinicRequests(id int64, mid int64) ([]models.ClinicMember, error) {
	var list []models.ClinicMember
	if !IsClinicAdmin(id, mid) {
		return list, errors.New("Unauthorized")
	}
	q := `SELECT * FROM public.clinic_member WHERE clinic_id = $1 AND is_accepted = false` //change to innerjoin
	err := db.Select(list, q, id)
	if err != nil {
		fmt.Println(err)
		return list, err
	}
	//or build array string statement
	for _, member := range list {
		var data models.User
		data, err := GetUser(*member.ConsultantID)
		if err != nil {
			fmt.Println(err)
			return list, err
		}
		member.Consultant = &data
	}

	return list, err
}

func ListClinicConsultants(id int64) ([]models.User, error) {
	var list []models.User
	var members []models.ClinicMember
	q := `SELECT * FROM public.clinic_member WHERE clinic_id = $1 AND is_accepted = true` //change to innerjoin
	err := db.Select(members, q, id)
	if err != nil {
		fmt.Println(err)
		return list, err
	}
	//or build array string statement
	for _, member := range members {
		var data models.User
		data, err := GetUser(*member.ConsultantID)
		if err != nil {
			fmt.Println(err)
			return list, err
		}
		list = append(list, data)
	}

	return list, err
}
