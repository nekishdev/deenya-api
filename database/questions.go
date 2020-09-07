package database

import (
	"deenya-api/models"
	"fmt"

	"github.com/lib/pq"
)

func GetQuestions(id int64) ([]models.Question, error) {
	var list []models.Question
	q := `SELECT id, questionnaire_id, question, answer, type, choices, asked_at, answered_at, owner_id FROM public.questionnaire_question WHERE questionnaire_id = $1`
	rows, err := db.Query(q, id)
	if err != nil {
		fmt.Println(err)
		return list, err
	}
	for rows.Next() {
		var data models.Question
		err := rows.Scan(&data.ID, &data.QuestionnaireID, &data.Question, &data.Answer, &data.Type, pq.Array(&data.Choices), &data.AskedAt, &data.AnsweredAt, &data.OwnerID)
		if err != nil {
			fmt.Println(err)
			return list, err
		}
		list = append(list, data)
	}
	return list, err
}

func UpdateQuestion(data models.Question) error {
	var err error
	umap := StructMap(data)
	uquery, err := UpString(umap)
	if err != nil {
		return err
	}
	q := "UPDATE public.questionnaire_question " + uquery + " WHERE id = :id AND owner_id = :owner_id" //

	_, err = db.NamedQuery(q, data)
	if err != nil {
		return err
	}
	return nil
}

func DeleteQuestion(id int64, uid int64) error {
	q := `DELETE FROM public.questionnaire_question WHERE id = $1 AND owner_id = $2` //
	// 	q := `DELETE FROM questionnaire_question f
	// WHERE id = $1 AND $2 IN (SELECT consultant_id FROM questionnaire WHERE id = f.questionnaire_id)`
	_, err := db.Exec(q, id, uid)
	if err != nil {
		return err
	}

	return nil
}

func NewQuestion(data *models.Question) error {
	_, csv, csvc := PrepareInsert(*data)
	sql := "INSERT INTO public.questionnaire_question" + " (" + csv + ") VALUES (" + csvc + ") RETURNING id, questionnaire_id, question, answer, type, choices, asked_at, answered_at, owner_id" //

	row, err := db.NamedQuery(sql, data)
	if err != nil {
		fmt.Println(err)
	}
	if row.Next() {
		row.Scan(&data.ID, &data.QuestionnaireID, &data.Question, &data.Answer, &data.Type, pq.Array(&data.Choices), &data.AskedAt, &data.AnsweredAt, &data.OwnerID)
	}

	return err
}

//---

func NewQuestionnaire(new *models.Questionnaire) error {
	var id int64

	_, csv, csvc := PrepareInsert(*new)
	sql := "INSERT INTO public.questionnaire" + " (" + csv + ") VALUES (" + csvc + ") RETURNING id" //

	row, err := db.NamedQuery(sql, new)
	if err != nil {
		fmt.Println(err)
	}
	if row.Next() {
		row.Scan(&id)
	}

	new.ID = &id
	return err
}

func DeleteQuestionnaire(id int64, uid int64) error {
	var err error
	q := `DELETE FROM public.questionnaire WHERE id = $1 AND consultant_id = $2` //

	_, err = db.Exec(q, id, uid)
	if err != nil {
		return err
	}

	return nil
}

func UpdateQuestionnaire(data models.Questionnaire) error {
	var err error
	umap := StructMap(data)
	uquery, err := UpString(umap)
	if err != nil {
		return err
	}
	q := "UPDATE public.questionnaire " + uquery + " WHERE id = :id AND consultant_id = :consultant_id" //

	_, err = db.NamedQuery(q, data)
	if err != nil {
		return err
	}
	return nil
}

func GetQuestionnaire(id int64, uid int64, utype string) (models.Questionnaire, error) {
	var err error
	var data models.Questionnaire

	q := fmt.Sprintf(`SELECT * FROM public.questionnaire WHERE id = $1 AND %s_id = $2`, utype)

	err = db.Get(&data, q, id, uid)
	if err != nil {
		return data, err
	}
	var questions []models.Question
	questions, err = GetQuestions(id)

	if err != nil {
		fmt.Println(err)
		return data, err
	}

	for _, question := range questions {
		data.Questions = append(data.Questions, &question)
	}

	return data, nil
}

func MyQuestionnaires(uid int64, utype string) ([]models.Questionnaire, error) {
	var data []models.Questionnaire
	q := fmt.Sprintf(`SELECT * FROM public.questionnaire WHERE %s_id = $1`, utype)
	err := db.Get(&data, q, uid)
	if err != nil {
		fmt.Println(err)
	}
	return data, err
}

func UserQuestionnaires(mid int64, mtype string, tid int64) ([]models.Questionnaire, error) {
	var ttype string
	if mtype == "client" {
		ttype = "consultant"
	}
	if mtype == "consultant" {
		ttype = "client"
	}
	var data []models.Questionnaire
	q := fmt.Sprintf(`SELECT * FROM public.questionnaire WHERE %s_id = $1 AND %s_id = $2 ORDER BY created_at DESC`, mtype, ttype)
	err := db.Get(&data, q, mid, tid)
	if err != nil {
		fmt.Println(err)
	}
	return data, err

}
