package database

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"

	"github.com/clarketm/json"

	"github.com/jmoiron/sqlx"
)

func Insert(db sqlx.Ext, table string, arg interface{}) (sql.Result, error) {
	fields := DBFields(arg) // e.g. []string{"id", "name", "description"}
	csv, csvc := prepFields(fields)
	sql := "INSERT INTO " + table + " (" + csv + ") VALUES (" + csvc + ")"
	return sqlx.NamedExec(db, sql, arg)
}

func prepFields(fields []string) (string, string) {
	var csv string
	var csvc string
	for i, field := range fields {
		if i == 0 {
			csv = field
			csvc = ":" + field
			continue
		}
		csv = csv + ", " + field
		csvc = csvc + ", " + ":" + field
	}
	return csv, csvc
}

func PrepareInsert(values interface{}) ([]string, string, string) {
	var field string
	var csv string
	var csvc string
	v := reflect.ValueOf(values)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	fields := []string{}
	if v.Kind() == reflect.Struct {
	fieldLoop:
		for i := 0; i < v.NumField(); i++ {
			if v.Field(i).IsZero() {
				continue fieldLoop
			}
			field = v.Type().Field(i).Tag.Get("db")
			if field != "" && field != "id" {
				if v.Field(i).Kind() == reflect.Slice { //if field is slice
					var sArr string

					for r := 0; r < v.Field(i).Len(); r++ {

						o := v.Field(i).Index(r)
						kind := o.Kind()
						if kind == reflect.Ptr {
							// fmt.Println("ccheck")
							o = o.Elem()
							kind = o.Kind()
						}

						if kind == reflect.Int64 {
							if r == 0 {
								sArr = fmt.Sprintf("ARRAY[%v", o.Interface())

								continue
							}
							sArr = sArr + fmt.Sprintf(", %v", o.Interface()) //get value of pointer in slice
						}

						if kind == reflect.String {
							if r == 0 {
								sArr = fmt.Sprintf("ARRAY['%v'", o.Interface())

								continue
							}
							sArr = sArr + fmt.Sprintf(", '%v'", o.Interface()) //get value of pointer in slice
						}

						if kind == reflect.Struct {
							//add nested struct slice functionality
						}

					}
					sArr = sArr + "]"
					//add array
					fields = append(fields, field)
					if csv == "" {
						csv = field
						csvc = sArr
						continue fieldLoop
					}
					csv = csv + ", " + field
					csvc = csvc + ", " + sArr
					continue fieldLoop
				}

				fields = append(fields, field)
				if csv == "" {
					csv = field
					csvc = ":" + field
					continue fieldLoop
				}
				csv = csv + ", " + field
				csvc = csvc + ", " + ":" + field
			}
		}
		return fields, csv, csvc
	}
	if v.Kind() == reflect.Map {
		for _, keyv := range v.MapKeys() {
			fields = append(fields, keyv.String())
		}
		return fields, csv, csvc
	}
	panic(fmt.Errorf("DBFields requires a struct or a map, found: %s", v.Kind().String()))
}

func DBFields(values interface{}) []string {
	var field string
	v := reflect.ValueOf(values)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	fields := []string{}
	if v.Kind() == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			if v.Type().Field(i).Type.Kind() == reflect.Slice { //if field is slice
				var sArr string
				for i := 0; i < v.NumField(); i++ {
					if i == 0 {
						sArr = fmt.Sprintf("ARRAY['%v'", v.Field(i).Elem())
						continue
					}
					sArr = sArr + fmt.Sprintf(", '%v'", v.Field(i).Elem()) //get value of pointer in slice
				}
				sArr = sArr + "]"
			}
			field = v.Type().Field(i).Tag.Get("db")
			if field != "" && field != "id" {
				fields = append(fields, field)
			}
			//v.FieldByName(v.Type().Field(i).Tag.Get("db")) get field
		}
		return fields
	}
	if v.Kind() == reflect.Map {
		for _, keyv := range v.MapKeys() {
			fields = append(fields, keyv.String())
		}
		return fields
	}
	panic(fmt.Errorf("DBFields requires a struct or a map, found: %s", v.Kind().String()))
}

func StructMap(values interface{}) map[string]interface{} {
	var dmap map[string]interface{}
	var err error
	js, err := json.Marshal(values)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(js, &dmap)
	if err != nil {
		fmt.Println(err)
	}
	return dmap
}

func UpString(cmap map[string]interface{}) (string, error) {
	var tstring string
	_, ok := cmap["id"]
	if !ok {
		return "", errors.New("Update failed: No ID in map")
	}
	delete(cmap, "id")
	delete(cmap, "user_id")
	delete(cmap, "owner_id")
	delete(cmap, "created_at")
	for k, v := range cmap {
		if tstring == "" {
			tstring = fmt.Sprintf("SET %v = '%v'", k, v)
			continue
		}
		tstring = fmt.Sprintf(tstring+", %v = '%v'", k, v)
	}
	return tstring, nil
}

func InsQuery(table string) {
	// q := "INSERT INTO " + table + "(" + "columns" + ")" + " (" + "values" + ")" + " RETURNING id"
}
