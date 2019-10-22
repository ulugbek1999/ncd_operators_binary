package api

import (
	"bytes"
	"fmt"
	"ncd_operators/pkg/raven"
	"strconv"

	"github.com/jmoiron/sqlx"

	//"encoding/json"

	m "ncd_operators/models"
	s "ncd_operators/pkg/settings"
	u "ncd_operators/pkg/utils"
	"net/http"
	"net/url"

	"github.com/gorilla/context"
)

// Operator1CreateAPI handles requests from operator 1 template
func Operator1CreateAPI(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		rows *sqlx.Rows
		id   int
		user = context.Get(r, "user").(*m.User)
	)

	sendSms := r.FormValue("send_sms")
	sendEmail := r.FormValue("send_email")
	username := r.FormValue("username")
	fullNameRu := r.FormValue("full_name_ru")
	fullNameEn := r.FormValue("full_name_en")
	birthDate := u.NewNullTime(r.FormValue("birth_date"))
	phone := r.FormValue("phone")
	email := r.FormValue("email")
	gender := r.FormValue("gender")
	passportSerial := r.FormValue("passport_serial")
	passportExpires := u.NewNullTime(r.FormValue("passport_expires"))
	passportGivenDate := u.NewNullTime(r.FormValue("passport_given_date"))
	inn := r.FormValue("inn")
	birthPlaceRu := r.FormValue("birth_place_ru")
	livingAddressRu := r.FormValue("living_address_ru")
	registerNumber := r.FormValue("register_number")
	q := `INSERT INTO employee (
		step_finished, full_name_ru, full_name_en, passport_serial, passport_given_date,
     	passport_expires, birth_date, birth_place_ru, living_address_ru, gender, inn,
     	phone, email, register_number, group_id, username, send_sms, send_email
	) VALUES (
		:step_finished, :full_name_ru, :full_name_en, :passport_serial, :passport_given_date,
		:passport_expires, :birth_date, :birth_place_ru, :living_address_ru, :gender, :inn,
     	:phone, :email, :register_number, :group_id, :username, :send_sms, :send_email
	) RETURNING id`
	v := map[string]interface{}{
		"step_finished":       1,
		"full_name_ru":        fullNameRu,
		"full_name_en":        fullNameEn,
		"passport_serial":     passportSerial,
		"passport_given_date": passportGivenDate,
		"passport_expires":    passportExpires,
		"birth_date":          birthDate,
		"birth_place_ru":      birthPlaceRu,
		"living_address_ru":   livingAddressRu,
		"gender":              gender,
		"inn":                 inn,
		"phone":               phone,
		"email":               email,
		"register_number":     registerNumber,
		"username":            username,
		"send_sms":            sendSms,
		"send_email":          sendEmail,
		"group_id":            user.Operator.OperatorGroup.Id,
	}

	rows, err = s.DB.NamedQuery(q, v)
	if err != nil {
		raven.ReportIfError(err)
		w.Write([]byte(err.Error()))
		return
	}
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			raven.ReportIfError(err)
			w.Write([]byte(err.Error()))
			return
		}
	}

	fNames := u.FileSave(r, "passport_image", r.FormValue("passport_serial"))
	for i := range fNames {

		if fNames[i] != "" {
			_, err = s.UDB.Exec(`UPDATE employee SET passport_image = $1 WHERE passport_serial = $2`,
				fNames[i], r.FormValue("passport_serial"))
			raven.ReportIfError(err)
		} else {
			_, err = s.UDB.Exec(`UPDATE employee SET passport_image='default/default.jpg' WHERE passport_serial = $1`, r.FormValue("passport_serial"))
			raven.ReportIfError(err)
		}
	}

	data := url.Values{}
	data.Add("emp_id", strconv.Itoa(id))
	req, err := http.NewRequest("POST", s.Conf["BASE_DOMAIN"]+"api/v2/ncd/user/create/", bytes.NewBufferString(data.Encode()))
	fmt.Println(err)
	raven.ReportIfError(err)
	var client = &http.Client{}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	_, err = client.Do(req)
	raven.ReportIfError(err)
	w.WriteHeader(201)
	return
}
