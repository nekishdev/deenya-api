package database

import (
	"deenya-api/models"
	"fmt"
)

func GetStripeConnect(id int64) (models.StripeConnect, error) {
	var data models.StripeConnect
	q := `SELECT * FROM public.stripe_connect WHERE consultant_id = $1`
	err := db.Get(&data, q, id)
	if err != nil {
		fmt.Println(err)
	}
	return data, err
}

func NewStripeConnect(new models.StripeConnect) error {
	_, csv, csvc := PrepareInsert(new)
	sql := "INSERT INTO public.stripe_connect" + " (" + csv + ") VALUES (" + csvc + ")"

	_, err := db.NamedQuery(sql, new)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func GetCustomer(clientID int64) (models.StripeCustomer, error) {
	var data models.StripeCustomer
	q := `SELECT * FROM public.stripe_customer WHERE client_id = $1`
	err := db.Get(&data, q, clientID)
	if err != nil {
		fmt.Println(err)
	}
	return data, err
}

func GetCustomerByID(customerID string) (models.StripeCustomer, error) {
	var data models.StripeCustomer
	q := `SELECT * FROM public.stripe_customer WHERE stripe_customer_id = $1`
	err := db.Get(&data, q, customerID)
	if err != nil {
		fmt.Println(err)
	}
	return data, err
}

func UpdateCustomer(data models.StripeCustomer) error {
	var err error
	umap := StructMap(data)
	uquery, err := UpString(umap)
	if err != nil {
		return err
	}
	q := "UPDATE public.stripe_customer " + uquery + " WHERE id = :id"

	_, err = db.NamedQuery(q, data)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCustomer(customerID string) error {
	q := `DELETE FROM public.stripe_customer WHERE stripe_customer_id = $1`

	_, err := db.Exec(q, customerID)
	if err != nil {
		return err
	}

	return nil
}

func NewCustomer(new *models.StripeCustomer) error {
	var id int64

	_, csv, csvc := PrepareInsert(*new)
	sql := "INSERT INTO public.stripe_customer" + " (" + csv + ") VALUES (" + csvc + ") RETURNING id"

	row, err := db.NamedQuery(sql, new)
	if err != nil {
		return err
	}
	if row.Next() {
		row.Scan(&id)
	}

	new.ID = &id
	return err
}

func ListCustomers(uid int64) ([]models.Customer, error) {
	var data []models.Customer
	q := `SELECT * FROM public.customer WHERE user = $1`
	err := db.Get(&data, q, uid)
	if err != nil {
		fmt.Println(err)
	}
	return data, err
}
