package database

import (
	"deenya-api/models"
	"fmt"

	"github.com/lib/pq"
	"github.com/stripe/stripe-go"
)

func NewInvoice(data *models.Invoice) error {
	_, csv, csvc := PrepareInsert(*data)

	q := "INSERT INTO public.invoice" + " (" + csv + ") VALUES (" + csvc + ") RETURNING *"
	row, err := db.NamedQuery(q, *data)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = row.Scan(&data)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil

}

func GetInvoice(id int64, mid int64, mtype string) (models.Invoice, error) {
	var data models.Invoice
	q := fmt.Sprintf(`SELECT * FROM public.invoice WHERE id = $1 AND %s_id = $2`, mtype)
	err := db.Get(&data, q, id, mid)
	if err != nil {
		fmt.Println(err)
		return data, err
	}

	return data, nil
}

func UpdateInvoice(data models.Invoice) error {
	umap := StructMap(data)
	uquery, err := UpString(umap)
	if err != nil {
		fmt.Println(err)
		return err
	}
	q := "UPDATE public.product " + uquery + " WHERE id = :id"
	_, err = db.NamedQuery(q, data)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func MyInvoices(mid int64, mtype string) ([]models.Invoice, error) {
	var data []models.Invoice

	q := fmt.Sprintf(`SELECT * FROM public.invoice WHERE %s_id = $1`, mtype)

	err := db.Get(&data, q, mid)

	if err != nil {
		fmt.Println(err)
		return data, err
	}

	return data, nil
}

func PayInvoice(id int64, mid int64) (stripe.Charge, error) {
	var charge stripe.Charge
	//get stripe customer details from db if exists, if not - redirect to create new customer
	//save charge details to database
	return charge, nil
	//stripe.NewCharge
}

func GetInvoiceBooking(id int64) (models.Booking, error) {
	//input invoice_id
	var data models.Booking
	q := `SELECT id, consultant_id, client_id, invoice_id, questionnaire_id, treatment_id, conversation_id, inquiry, tags, created_at, started_at, ended_at, scheduled_at, elapsed, status, is_accepted, is_remote FROM public.booking WHERE invoice_id = $1`
	row := db.QueryRow(q, id)
	err := row.Scan(&data.ID, &data.ConsultantID, &data.ClientID, &data.InvoiceID, &data.QuestionnaireID, &data.TreatmentID, &data.ConversationID, &data.Inquiry, pq.Array(&data.Tags), &data.CreatedAt, &data.StartedAt, &data.EndedAt, &data.ScheduledAt, &data.IsAccepted, &data.IsRemote)
	if err != nil {
		fmt.Println(err)
		return data, err
	}
	return data, nil
}
