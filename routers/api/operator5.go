package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"ncd_operators/pkg/raven"
	s "ncd_operators/pkg/settings"
	u "ncd_operators/pkg/utils"
	"net/http"
	"net/url"
	"strconv"
)



// UniversalCreateAPI handles requests from the template of operator 5
func UniversalCreateAPI(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		id   int
		eduF int
	)
	fullNameRu := r.FormValue("full_name_ru")
	fullNameEn := r.FormValue("full_name_en")
	birthDate := u.NewNullTime(r.FormValue("birth_date"))
	phone := r.FormValue("phone")
	email := r.FormValue("email")
	gender := r.FormValue("gender")
	passportSerial := r.FormValue("passport_serial")
	passportGivenDate := u.NewNullTime(r.FormValue("passport_given_date"))
	passportExpires := u.NewNullTime(r.FormValue("passport_expires"))
	inn := r.FormValue("inn")
	birthPlaceRu := r.FormValue("birth_place_ru")
	livingAddressRu := r.FormValue("living_address_ru")
	registerNumber := r.FormValue("register_number")
	sendSMS := r.FormValue("send_sms")
	sendEmail := r.FormValue("send_email")
	username := r.FormValue("username")
	_, _, err = r.FormFile("passport_image")

	appearance := u.NewNullInt(r.FormValue("appearance"))
	neatness := u.NewNullInt(r.FormValue("neatness"))
	skin := u.NewNullInt(r.FormValue("skin"))
	height := u.NewNullFloat(r.FormValue("height"))
	weight := u.NewNullFloat(r.FormValue("weight"))
	fatness := u.NewNullFloat(r.FormValue("fatness"))
	bloodGroup := u.NewNullInt(r.FormValue("blood_group"))
	bloodResus := r.FormValue("blood_resus")
	visionL := u.NewNullFloat(r.FormValue("vision_l"))
	visionR := u.NewNullFloat(r.FormValue("vision_r"))

	isReadForUn := r.FormValue("is_ready_for_university")
	// criminalNumber := r.FormValue("criminal_number")
	// criminalIssue := r.FormValue("criminal_issue")
	// criminalCommentRu := r.FormValue("criminal_comment_ru")
	hobbyRu := r.FormValue("hobby_ru")
	otherRu := r.FormValue("other_ru")

	q := `INSERT INTO employee (
		step_finished, full_name_ru, full_name_en, birth_date, phone, email, gender, 
		passport_serial, passport_given_date, passport_expires, inn, birth_place_ru,
		living_address_ru, register_number, send_sms, send_email, username,
		appearance, neatness, skin, height, weight, fatness, blood_group, blood_resus, 
		vision_l, vision_r, is_ready_for_university, hobby_ru, other_ru) 
		VALUES (:step_finished, :full_name_ru, :full_name_en, :birth_date, :phone,
		:email, :gender, :passport_serial, :passport_given_date, :passport_expires,
		:inn, :birth_place_ru, :living_address_ru, :register_number, 
		:send_sms, :send_email, :username, :appearance, :neatness, :skin, :height, 
		:weight, :fatness, :blood_group, :blood_resus, :vision_l, :vision_r,
		:is_ready_for_university,
		:hobby_ru, :other_ru) RETURNING id
	`
	v := map[string]interface{}{
		"step_finished":           4,
		"full_name_ru":            fullNameRu,
		"full_name_en":            fullNameEn,
		"birth_date":              birthDate,
		"phone":                   phone,
		"email":                   email,
		"gender":                  gender,
		"passport_serial":         passportSerial,
		"passport_given_date":     passportGivenDate,
		"passport_expires":        passportExpires,
		"inn":                     inn,
		"birth_place_ru":          birthPlaceRu,
		"living_address_ru":       livingAddressRu,
		"register_number":         registerNumber,
		"send_sms":                sendSMS,
		"send_email":              sendEmail,
		"username":                username,
		"appearance":              appearance,
		"neatness":                neatness,
		"skin":                    skin,
		"height":                  height,
		"weight":                  weight,
		"fatness":                 fatness,
		"blood_group":             bloodGroup,
		"blood_resus":             bloodResus,
		"vision_l":                visionL,
		"vision_r":                visionR,
		"is_ready_for_university": isReadForUn,
		"hobby_ru":                hobbyRu,
		"other_ru":                otherRu,
	}

	rows, err := s.DB.NamedQuery(q, v)
	if err != nil {
		raven.ReportIfError(err)
		w.Write([]byte(err.Error()))
		panic(err)
	}
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			raven.ReportIfError(err)
			w.Write([]byte(err.Error()))
			return
		}
	}

	fNames := u.FileSave(r, "passport_image", passportSerial)
	for i := range fNames {
		if fNames[i] != "" {
			_, err = s.UDB.Exec(`UPDATE employee SET passport_image = $1 WHERE 
			passport_serial = $2`, fNames[i], passportSerial)
		}
	}

	fNames = u.FileSave(r, "photo_1", passportSerial)
	for i := range fNames {
		if fNames[i] != "" {
			_, err = s.UDB.Exec(`UPDATE employee SET photo_1 = $1 WHERE passport_serial = $2`, fNames[i], passportSerial)
			raven.ReportIfError(err)
			panic(err)
		}
	}

	fNames = u.FileSave(r, "photo_2", passportSerial)
	for i := range fNames {
		if fNames[i] != "" {
			_, err = s.UDB.Exec(`UPDATE employee SET photo_2 = $1 WHERE passport_serial = $2`, fNames[i], passportSerial)
			raven.ReportIfError(err)
			panic(err)
		}
	}

	fNames = u.FileSave(r, "photo_3", passportSerial)
	for i := range fNames {
		if fNames[i] != "" {
			_, err = s.UDB.Exec(`UPDATE employee SET photo_3 = $1 WHERE passport_serial = $2`, fNames[i], passportSerial)
			raven.ReportIfError(err)
			panic(err)
		}
	}

	fNames = u.FileSave(r, "photo_4", passportSerial)
	for i := range fNames {
		if fNames[i] != "" {
			_, err = s.UDB.Exec(`UPDATE employee SET photo_4 = $1 WHERE passport_serial = $2`, fNames[i], passportSerial)
			raven.ReportIfError(err)
			panic(err)
		}
	}

	// Operator 3
	// Education
	fNames = u.FileSave(r, "edu_files", passportSerial)
	var eduList []map[string]interface{}
	if err := json.Unmarshal([]byte(r.FormValue("edu_data")), &eduList); err != nil {
		panic(err)
	}

	fmt.Println(eduList)
	q = `INSERT INTO employee__education (name_ru, address_ru, specialization_ru, date_started,
		date_finished, additional_ru, is_new, employee_id, type_id)
		 VALUES (:name_ru, :address_ru, :specialization_ru, :date_started, :date_finished,
		:additional_ru, :is_new, :employee_id, :type_id) RETURNING id`
	for i, edu := range eduList {
		fmt.Println(i)
		v = map[string]interface{}{
			"name_ru":           edu["eduName"],
			"address_ru":        edu["eduAddress"],
			"specialization_ru": edu["eduSpec"],
			"date_started":      edu["eduStartDate"],
			"date_finished":     edu["eduGradDate"],
			"additional_ru":     edu["eduAddit"],
			"is_new":            false,
			"employee_id":       id,
			"type_id":           edu["eduType"],
		}
		rows, err = s.UDB.NamedQuery(q, v)
		for rows.Next() {
			err = rows.Scan(&eduF)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		fmt.Println(eduF)
		return
		_, err = s.UDB.NamedExec(`INSERT INTO employee__education__file (file, education_id) VALUES (:file, :edu_id)`,
			map[string]interface{}{
				"file":   fNames[i],
				"edu_id": eduF,
			})
		if err != nil {
			panic(err)
		}
	}

	// for _, fh := range fhs {
	// 	_, err = s.UDB.NamedExec(`INSERT INTO employee__education__file (file, education_id)`)
	// }

	// Language
	fmt.Println(r.FormValue("lang_data"))
	lhs := r.MultipartForm.File["lang_files"]
	for _, lh := range lhs {
		fmt.Println(lh.Filename)
	}

	// Operator 4

	data := url.Values{}
	data.Add("emp_id", strconv.Itoa(id))
	req, err := http.NewRequest("POST", s.Conf["BASE_DOMAIN"]+"api/v2/ncd/user/create", bytes.NewBufferString(data.Encode()))
	fmt.Println(err)
	raven.ReportIfError(err)
	var client = &http.Client{}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	_, err = client.Do(req)
	raven.ReportIfError(err)
	w.WriteHeader(201)
	return
}
