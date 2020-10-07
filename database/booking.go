package database

import (
	"deenya-api/models"
	"fmt"
	"time"

	"github.com/lib/pq"
)

//get, update, delete, new, my list (aid), user's list(my id, my type, target id, target type) [check if other user is a client or consultant]
//#####

func GetBooking(id int64, uid int64, utype string) (models.Booking, error) {
	var data models.Booking
	q := fmt.Sprintf(`SELECT id, consultant_id, client_id, invoice_id, questionnaire_id, treatment_id, conversation_id, inquiry, tags, created_at, started_at, ended_at, scheduled_at, is_accepted, is_remote FROM public.booking WHERE id = $1 AND %s_id = $2 LIMIT 1`, utype)
	row := db.QueryRow(q, id, uid)
	err := row.Scan(&data.ID, &data.ConsultantID, &data.ClientID, &data.InvoiceID, &data.QuestionnaireID, &data.TreatmentID, &data.ConversationID, &data.Inquiry, pq.Array(&data.Tags), &data.CreatedAt, &data.StartedAt, &data.EndedAt, &data.ScheduledAt, &data.IsAccepted, &data.IsRemote)
	if err != nil {
		fmt.Println(err)
		return data, err
	}

	return data, nil
}

func UpdateBooking(data models.Booking, utype string) error {
	var err error
	umap := StructMap(data.BookingData)
	uquery, err := UpString(umap)
	if err != nil {
		return err
	}
	q := fmt.Sprintf("UPDATE public.booking "+uquery+" WHERE id = :id AND %s_id = :%s_id", utype, utype)

	_, err = db.NamedQuery(q, data)
	if err != nil {
		return err
	}
	return nil
}

func DeleteBooking(id int64, uid int64, utype string) error {
	q := fmt.Sprintf(`DELETE FROM public.booking WHERE id = $1 AND %s_id = $2`, utype)

	_, err := db.Exec(q, id, uid)
	if err != nil {
		return err
	}

	return nil
}

func NewBooking(new *models.Booking) error {
	var id int64
	_, csv, csvc := PrepareInsert(new.BookingData)
	sql := "INSERT INTO public.booking" + " (" + csv + ") VALUES (" + csvc + ") RETURNING id"

	row, err := db.NamedQuery(sql, new.BookingData)
	if err != nil {
		fmt.Println(err)
	}
	if row.Next() {
		row.Scan(&id)
	}

	new.ID = &id
	return err
}

func MyBookings(uid int64, utype string) ([]models.Booking, error) {
	var list []models.Booking
	q := fmt.Sprintf(`SELECT id, consultant_id, client_id, invoice_id, questionnaire_id, treatment_id, conversation_id, inquiry, tags, created_at, started_at, ended_at, scheduled_at, is_accepted, is_remote FROM public.booking WHERE %s_id = $1 ORDER BY scheduled_at DESC`, utype)
	rows, err := db.Query(q, uid)
	if err != nil {
		fmt.Println(err)
		return list, err
	}
	for rows.Next() {
		var data models.Booking
		err = rows.Scan(&data.ID, &data.ConsultantID, &data.ClientID, &data.InvoiceID, &data.QuestionnaireID, &data.TreatmentID, &data.ConversationID, &data.Inquiry, pq.Array(&data.Tags), &data.CreatedAt, &data.StartedAt, &data.EndedAt, &data.ScheduledAt, &data.IsAccepted, &data.IsRemote)
		if err != nil {
			fmt.Println(err)
			return list, err
		}
		list = append(list, data)
	}
	return list, nil
}

func UserBookings(mid int64, mtype string, tid int64) ([]models.Booking, error) {
	var ttype string
	if mtype == "client" {
		ttype = "consultant"
	}
	if mtype == "consultant" {
		ttype = "client"
	}
	var list []models.Booking
	q := fmt.Sprintf(`SELECT id, consultant_id, client_id, invoice_id, questionnaire_id, treatment_id, conversation_id, inquiry, tags, created_at, started_at, ended_at, scheduled_at, is_accepted, is_remote FROM public.booking WHERE %s_id = $1 AND %s_id = $2 ORDER BY scheduled_at DESC`, mtype, ttype)
	rows, err := db.Query(q, mid, tid)
	if err != nil {
		fmt.Println(err)
		return list, err
	}
	for rows.Next() {
		var data models.Booking
		err = rows.Scan(&data.ID, &data.ConsultantID, &data.ClientID, &data.InvoiceID, &data.QuestionnaireID, &data.TreatmentID, &data.ConversationID, &data.Inquiry, pq.Array(&data.Tags), &data.CreatedAt, &data.StartedAt, &data.EndedAt, &data.ScheduledAt, &data.IsAccepted, &data.IsRemote)
		if err != nil {
			fmt.Println(err)
			return list, err
		}
		list = append(list, data)
	}
	return list, nil

}

func VerifyBooking(id int64, uid int64, utype string) bool {
	var verified bool

	q := fmt.Sprintf(`SELECT EXISTS(SELECT 1 FROM public.booking WHERE id = $1 AND %s_id = $2 LIMIT 1)`, utype)

	err := db.Get(&verified, q, id, uid)

	if err != nil {
		return false
	}

	return verified
}

// func AcceptBooking(id int64, uid int64) error {
// 	q := `UPDATE public.booking SET is_accepted = 'true' WHERE id = $1`
// }

func AvailableBookings(uid int64, date time.Time, tz string) ([]time.Time, error) {
	//get date from frontend with timezone
	//get slots within available range of consultant as requested timezone (time at time zone 'America/Los_Angeles')
	//(in frontend, check if chosen date = today, if so then send hours, if not then hours = start of day) from request time to end of day in requested timezone
	//if time.Before start of day/current time or time.After end of day (of consultant time zone slots converted to requested timezone), remove slots from response

	var list []time.Time
	user, err := GetUserDetails(uid)
	if err != nil {
		fmt.Println(err)
		return list, err
	}
	consultant, err := GetUserConsultant(uid)
	if err != nil {
		fmt.Println(err)
		return list, err
	}
	reqLoc, err := time.LoadLocation(tz)
	if err != nil {
		fmt.Println(err)
		return list, err
	}
	userLoc, err := time.LoadLocation(*user.Timezone)
	if err != nil {
		fmt.Println(err)
		return list, err
	}

	reqStart := time.Date(date.Year(), date.Month(), date.Day(), date.Hour(), 0, 0, 0, reqLoc)
	reqEndOfDay := time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 0, reqLoc)
	reqEnd := reqStart.Add(24 * time.Hour)

	q := `SELECT dt FROM generate_series($1::timestamp without timezone, $2::timestamp without timezone, '30 m') dt LEFT JOIN public.booking b ON b.scheduled_at >= dt AND b.scheduled_at < dt + interval '30 m'`
	fmt.Println(q)
	rows, err := db.Query(q, reqStart, reqEnd)
	if err != nil {
		fmt.Println(err)
		return list, err
	}

	for rows.Next() {
		var data time.Time

		rows.Scan(&data)
		if data.In(userLoc).Hour() > *consultant.AvailableTo {
			//if date is after consultant hours
			break
		}
		if data.In(userLoc).Hour() < *consultant.AvailableFrom {
			//if date is before consultant hours
			break
		}
		if data.In(reqLoc).After(reqEndOfDay) {
			//if date goes into next day for end user
			break
		}
		list = append(list, data)
	}

	return list, nil
}
