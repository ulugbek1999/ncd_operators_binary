package api

import (
	m "ncd_operators/models"
	"ncd_operators/pkg/raven"
	s "ncd_operators/pkg/settings"
	u "ncd_operators/pkg/utils"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func EducationCreateAPI(w http.ResponseWriter, r *http.Request) {
	var (
		vars = mux.Vars(r)
	)
	eduType := r.FormValue("type")
	nameRu := r.FormValue("name_ru")
	addressRu := r.FormValue("address_ru")
	specializationRu := r.FormValue("specialization_ru")
	dateStarted := u.NewNullTime(r.FormValue("date_started"))
	dateFinished := u.NewNullTime(r.FormValue("date_finished"))
	additionalRu := r.FormValue("additional_ru")

	q := `INSERT INTO employee__education (name_ru, address_ru, specialization_ru, date_started,
		  date_finished, additional_ru, is_new, employee_id, type_id)
	 	  VALUES (:name_ru, :address_ru, :specialization_ru, :date_started, :date_finished,
		  :additional_ru, :is_new, :employee_id, :type_id) RETURNING id`
	v := map[string]interface{}{
		"name_ru":           nameRu,
		"address_ru":        addressRu,
		"specialization_ru": specializationRu,
		"date_started":      dateStarted,
		"date_finished":     dateFinished,
		"additional_ru":     additionalRu,
		"is_new":            false,
		"employee_id":       vars["id"],
		"type_id":           eduType,
	}
	rows, err := s.UDB.NamedQuery(q, v)
	if err != nil {
		raven.ReportIfError(err)
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	var id int
	for rows.Next() {
		err = rows.Scan(&id)
		raven.ReportIfError(err)
	}
	e, err := m.EmployeeById(vars["id"])
	raven.ReportIfError(err)
	fNames := u.FileSave(r, "education", e.PassportSerial)
	for i := range fNames {
		if fNames[i] != "" {
			_, err = s.UDB.NamedExec(`INSERT INTO employee__education__file (file, education_id) VALUES (:file, :edu_id)`,
				map[string]interface{}{
					"file":   fNames[i],
					"edu_id": id,
				})
			raven.ReportIfError(err)
		}
	}
	w.WriteHeader(200)
	return
}

func LanguageCreateAPI(w http.ResponseWriter, r *http.Request) {
	var (
		vars = mux.Vars(r)
		err  error
		rows *sqlx.Rows
	)
	q := `INSERT INTO employee__language (level, is_required_to_check, is_new, employee_id, language_id) VALUES 
	  	  (:level, :check, :is_new, :employee_id, :lang_id) RETURNING id`
	rows, err = s.UDB.NamedQuery(q, map[string]interface{}{
		"lang_id":     r.FormValue("language"),
		"level":       r.FormValue("level"),
		"check":       false,
		"is_new":      false,
		"employee_id": vars["id"],
	})
	if err != nil {
		raven.ReportIfError(err)
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	var id int
	for rows.Next() {
		err = rows.Scan(&id)
		raven.ReportIfError(err)
	}
	e, err := m.EmployeeById(vars["id"])
	raven.ReportIfError(err)
	fNames := u.FileSave(r, "language", e.PassportSerial)
	for i := range fNames {
		if fNames[i] != "" {
			_, err = s.UDB.NamedExec(`insert into employee__language__file (file, language_id) values (:file, :lang_id)`,
				map[string]interface{}{
					"file":    fNames[i],
					"lang_id": id,
				})
			raven.ReportIfError(err)
		}
	}
	w.WriteHeader(200)
	return
}

func ArmyCreateAPI(w http.ResponseWriter, r *http.Request) {
	var (
		vars = mux.Vars(r)
		err  error
		rows *sqlx.Rows
	)
	q := `INSERT INTO employee__army (name_ru, region_ru, specialization_ru, date_started, date_finished, rank_ru, is_new, employee_id) 
	VALUES (:name_ru, :region_ru, :specialization_ru, :date_started, :date_finished, :rank_ru, :is_new, :employee_id) RETURNING id`
	rows, err = s.UDB.NamedQuery(q, map[string]interface{}{
		"name_ru":           r.FormValue("name_ru"),
		"region_ru":         r.FormValue("region_ru"),
		"specialization_ru": r.FormValue("specialization_ru"),
		"date_started":      u.NewNullTime(r.FormValue("date_started")),
		"date_finished":     u.NewNullTime(r.FormValue("date_finished")),
		"rank_ru":           r.FormValue("rank_ru"),
		"is_new":            false,
		"employee_id":       vars["id"],
	})
	if err != nil {
		raven.ReportIfError(err)
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	var id int
	for rows.Next() {
		err = rows.Scan(&id)
		raven.ReportIfError(err)
	}
	e, err := m.EmployeeById(vars["id"])
	raven.ReportIfError(err)
	fNames := u.FileSave(r, "army", e.PassportSerial)
	for i := range fNames {
		if fNames[i] != "" {
			_, err = s.UDB.NamedExec(`INSERT INTO employee__army__file (file, army_id) VALUES (:file, :army_id)`,
				map[string]interface{}{
					"file":    fNames[i],
					"army_id": id,
				})
			raven.ReportIfError(err)
		}
	}
	return
}

func RewardCreateAPI(w http.ResponseWriter, r *http.Request) {
	var (
		vars = mux.Vars(r)
		err  error
		rows *sqlx.Rows
	)
	q := `INSERT INTO employee__reward (name_ru, description_ru, is_new, employee_id) VALUES 
		  (:name_ru, :description_ru, :is_new, :employee_id) RETURNING id`
	rows, err = s.UDB.NamedQuery(q, map[string]interface{}{
		"name_ru":        r.FormValue("name_ru"),
		"description_ru": r.FormValue("description_ru"),
		"is_new":         false,
		"employee_id":    vars["id"],
	})
	if err != nil {
		raven.ReportIfError(err)
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	var id int
	for rows.Next() {
		err = rows.Scan(&id)
		raven.ReportIfError(err)
	}
	e, err := m.EmployeeById(vars["id"])
	raven.ReportIfError(err)
	fNames := u.FileSave(r, "reward", e.PassportSerial)
	for i := range fNames {
		if fNames[i] != "" {
			_, err = s.UDB.NamedExec(`insert into employee__reward__file (file, reward_id) values (:file, :reward_id)`,
				map[string]interface{}{
					"file":      fNames[i],
					"reward_id": id,
				})
			raven.ReportIfError(err)
		}
	}
	return
}

func FamilyCreateAPI(w http.ResponseWriter, r *http.Request) {
	var (
		vars = mux.Vars(r)
		err  error
		rows *sqlx.Rows
	)
	q := `INSERT INTO employee__family (status, children_amount, is_new, employee_id) VALUES
		 (:status, :children_amount, :is_new, :employee_id) RETURNING id`
	rows, err = s.UDB.NamedQuery(q, map[string]interface{}{
		"status":          r.FormValue("status"),
		"children_amount": r.FormValue("children_amount"),
		"is_new":          false,
		"employee_id":     vars["id"],
	})
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	var id int
	for rows.Next() {
		err = rows.Scan(&id)
		raven.ReportIfError(err)
	}
	e, err := m.EmployeeById(vars["id"])
	raven.ReportIfError(err)
	fNames := u.FileSave(r, "family", e.PassportSerial)
	for i := range fNames {
		if fNames[i] != "" {
			_, err = s.UDB.NamedExec(`insert into employee__family__file (file, family_id) values (:file, :family_id)`,
				map[string]interface{}{
					"file":      fNames[i],
					"family_id": id,
				})
			raven.ReportIfError(err)
		}
	}
	return
}

func RelativeCreateAPI(w http.ResponseWriter, r *http.Request) {
	var (
		vars = mux.Vars(r)
		err  error
		rows *sqlx.Rows
	)
	q := `INSERT INTO employee__relative (level, fullname_ru, birth_date, study_work_place_ru, position_ru, living_address_ru, is_new, employee_id) VALUES
		  (:level, :fullname_ru, :birth_date, :work_place_ru, :position_ru, :living_address_ru, :is_new, :employee_id) RETURNING id`
	rows, err = s.UDB.NamedQuery(q, map[string]interface{}{
		"level":             r.FormValue("level"),
		"fullname_ru":       r.FormValue("fullname_ru"),
		"birth_date":        r.FormValue("birth_date"),
		"work_place_ru":     r.FormValue("study_work_place_ru"),
		"position_ru":       r.FormValue("position_ru"),
		"living_address_ru": r.FormValue("living_address_ru"),
		"is_new":            false,
		"employee_id":       vars["id"],
	})
	if err != nil {
		raven.ReportIfError(err)
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	var id int
	for rows.Next() {
		err = rows.Scan(&id)
		raven.ReportIfError(err)
	}
	e, err := m.EmployeeById(vars["id"])
	raven.ReportIfError(err)
	fNames := u.FileSave(r, "relative", e.PassportSerial)
	for i := range fNames {
		if fNames[i] != "" {
			_, err = s.UDB.NamedExec(`insert into employee__relative__file (file, relative_id) values (:file, :relative_id)`,
				map[string]interface{}{
					"file":        fNames[i],
					"relative_id": id,
				})
			raven.ReportIfError(err)
		}
	}
	return
}

func ExperienceCreateAPI(w http.ResponseWriter, r *http.Request) {
	var (
		vars = mux.Vars(r)
		err  error
		rows *sqlx.Rows
	)
	q := `INSERT INTO employee__experience (organization_ru, date_started, date_finished, position_ru, sub_division_ru, address_ru, is_new, employee_id) VALUES 
          (:organization_ru, :date_started, :date_finished, :position_ru, :sub_division_ru, :address_ru, :is_new, :employee_id) RETURNING id`
	rows, err = s.UDB.NamedQuery(q, map[string]interface{}{
		"organization_ru": r.FormValue("organization_ru"),
		"date_started":    u.NewNullTime(r.FormValue("date_started")),
		"date_finished":   u.NewNullTime(r.FormValue("date_finished")),
		"position_ru":     r.FormValue("position_ru"),
		"sub_division_ru": r.FormValue("sub_division_ru"),
		"address_ru":      r.FormValue("address_ru"),
		"is_new":          false,
		"employee_id":     vars["id"],
	})
	if err != nil {
		raven.ReportIfError(err)
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	var id int
	for rows.Next() {
		err = rows.Scan(&id)
		raven.ReportIfError(err)
	}
	e, err := m.EmployeeById(vars["id"])
	raven.ReportIfError(err)
	fNames := u.FileSave(r, "experience", e.PassportSerial)
	for i := range fNames {
		if fNames[i] != "" {
			_, err = s.UDB.NamedExec(`insert into employee__experience__file (file, experience_id) values (:file, :experience_id)`,
				map[string]interface{}{
					"file":          fNames[i],
					"experience_id": id,
				})
			raven.ReportIfError(err)
		}
	}
	return
}

func Operator3Close(w http.ResponseWriter, r *http.Request) {
	var (
		vars = mux.Vars(r)
		err  error
	)
	q := `UPDATE employee SET step_finished = 3 where id = $1`
	_, err = s.UDB.Exec(q, vars["id"])
	if err != nil {
		raven.ReportIfError(err)
		w.Write([]byte(err.Error()))
		return
	}
	return
}

// EDIT

func EducationEditAPI(w http.ResponseWriter, r *http.Request) {
	var (
		vars = mux.Vars(r)
		err  error
	)
	q := `update employee__education  set (name_ru, address_ru, specialization_ru, date_started,
		  date_finished, additional_ru, type_id) = (:name_ru, :address_ru,
		  :specialization_ru, :date_started, :date_finished, :additional_ru, :type_id) where id = :id`
	_, err = s.UDB.NamedExec(q, map[string]interface{}{
		"name_ru":           r.FormValue("name_ru"),
		"address_ru":        r.FormValue("address_ru"),
		"specialization_ru": r.FormValue("specialization_ru"),
		"date_started":      r.FormValue("date_started"),
		"date_finished":     r.FormValue("date_finished"),
		"additional_ru":     r.FormValue("additional_ru"),
		"type_id":           r.FormValue("type"),
		"id":                vars["id"],
	})
	raven.ReportIfError(err)
	e, _ := m.EmployeeById(vars["emp_id"])
	fNames := u.FileSave(r, "education", e.PassportSerial)
	for i := range fNames {
		if fNames[i] != "" {
			_, err = s.UDB.NamedExec(`INSERT INTO employee__education__file (file, education_id) VALUES (:file, :edu_id)`,
				map[string]interface{}{
					"file":   fNames[i],
					"edu_id": vars["id"],
				})
			raven.ReportIfError(err)
		}
	}
}

func ArmyEditAPI(w http.ResponseWriter, r *http.Request) {
	var (
		vars = mux.Vars(r)
		err  error
	)
	q := `update employee__army set (name_ru, region_ru, specialization_ru, date_started, date_finished, rank_ru) 
	= (:name_ru, :region_ru, :specialization_ru, :date_started, :date_finished, :rank_ru) where id = :id`
	_, err = s.UDB.NamedExec(q, map[string]interface{}{
		"name_ru":           r.FormValue("name_ru"),
		"region_ru":         r.FormValue("region_ru"),
		"specialization_ru": r.FormValue("specialization_ru"),
		"date_started":      r.FormValue("date_started"),
		"date_finished":     r.FormValue("date_finished"),
		"rank_ru":           r.FormValue("rank_ru"),
		"id":                vars["id"],
	})
	raven.ReportIfError(err)
	e, _ := m.EmployeeById(vars["emp_id"])
	fNames := u.FileSave(r, "army", e.PassportSerial)
	for i := range fNames {
		if fNames[i] != "" {
			_, err = s.UDB.NamedExec(`INSERT INTO employee__army__file (file, army_id) VALUES (:file, :army_id)`,
				map[string]interface{}{
					"file":    fNames[i],
					"army_id": vars["id"],
				})
			raven.ReportIfError(err)
		}
	}
}

func LanguageEditAPI(w http.ResponseWriter, r *http.Request) {
	var (
		vars = mux.Vars(r)
		err  error
	)
	q := `update employee__language set (level, language_id) =
	  	  (:level, :lang_id) where id = :id`
	_, err = s.UDB.NamedExec(q, map[string]interface{}{
		"lang_id": r.FormValue("language"),
		"level":   r.FormValue("level"),
		//"check":   r.FormValue("is_required_to_check"),
		"id": vars["id"],
	})
	raven.ReportIfError(err)
	e, _ := m.EmployeeById(vars["emp_id"])
	fNames := u.FileSave(r, "language", e.PassportSerial)
	for i := range fNames {
		if fNames[i] != "" {
			_, err = s.UDB.NamedExec(`insert into employee__language__file (file, language_id) values (:file, :lang_id)`,
				map[string]interface{}{
					"file":    fNames[i],
					"lang_id": vars["id"],
				})
			raven.ReportIfError(err)
		}
	}
}

func FamilyEditAPI(w http.ResponseWriter, r *http.Request) {
	var (
		vars = mux.Vars(r)
		err  error
	)
	q := `update employee__family set (status, children_amount) =
		 (:status, :children_amount) where id = :id`
	_, err = s.UDB.NamedExec(q, map[string]interface{}{
		"status":          r.FormValue("status"),
		"children_amount": r.FormValue("children_amount"),
		"id":              vars["id"],
	})
	raven.ReportIfError(err)
	e, _ := m.EmployeeById(vars["emp_id"])
	fNames := u.FileSave(r, "family", e.PassportSerial)
	for i := range fNames {
		if fNames[i] != "" {
			_, err = s.UDB.NamedExec(`insert into employee__family__file (file, family_id) values (:file, :family_id)`,
				map[string]interface{}{
					"file":      fNames[i],
					"family_id": vars["id"],
				})
			raven.ReportIfError(err)
		}
	}
}

func RelativeEditAPI(w http.ResponseWriter, r *http.Request) {
	var (
		vars = mux.Vars(r)
		err  error
	)
	q := `update employee__relative set (level, fullname_ru, birth_date, study_work_place_ru, position_ru, living_address_ru) =
		  (:level, :fullname_ru, :birth_date, :work_place_ru, :position_ru, :living_address_ru) where id = :id`
	_, err = s.UDB.NamedExec(q, map[string]interface{}{
		"level":             r.FormValue("level"),
		"fullname_ru":       r.FormValue("fullname_ru"),
		"birth_date":        r.FormValue("birth_date"),
		"work_place_ru":     r.FormValue("study_work_place_ru"),
		"position_ru":       r.FormValue("position_ru"),
		"living_address_ru": r.FormValue("living_address_ru"),
		"id":                vars["id"],
	})
	raven.ReportIfError(err)
	e, _ := m.EmployeeById(vars["emp_id"])
	fNames := u.FileSave(r, "relative", e.PassportSerial)
	for i := range fNames {
		if fNames[i] != "" {
			_, err = s.UDB.NamedExec(`insert into employee__relative__file (file, relative_id) values (:file, :relative_id)`,
				map[string]interface{}{
					"file":        fNames[i],
					"relative_id": vars["id"],
				})
			raven.ReportIfError(err)
		}
	}
}

func ExperienceEditAPI(w http.ResponseWriter, r *http.Request) {
	var (
		vars = mux.Vars(r)
		err  error
	)
	q := `update employee__experience  set (organization_ru, date_started, date_finished, position_ru, sub_division_ru, address_ru) = 
          (:organization_ru, :date_started, :date_finished, :position_ru, :sub_division_ru, :address_ru) where id = :id`
	_, err = s.UDB.NamedExec(q, map[string]interface{}{
		"organization_ru": r.FormValue("organization_ru"),
		"date_started":    r.FormValue("date_started"),
		"date_finished":   r.FormValue("date_finished"),
		"position_ru":     r.FormValue("position_ru"),
		"sub_division_ru": r.FormValue("sub_division_ru"),
		"address_ru":      r.FormValue("address_ru"),
		"id":              vars["id"],
	})
	raven.ReportIfError(err)
	e, _ := m.EmployeeById(vars["emp_id"])
	fNames := u.FileSave(r, "experience", e.PassportSerial)
	for i := range fNames {
		if fNames[i] != "" {
			_, err = s.UDB.NamedExec(`insert into employee__experience__file (file, experience_id) values (:file, :experience_id)`,
				map[string]interface{}{
					"file":          fNames[i],
					"experience_id": vars["em_id"],
				})
			raven.ReportIfError(err)
		}
	}
}

func RewardEditAPI(w http.ResponseWriter, r *http.Request) {
	var (
		vars = mux.Vars(r)
		err  error
	)
	q := `update employee__reward  set (name_ru, description_ru) = 
		  (:name_ru, :description_ru) where id = :id`
	_, err = s.UDB.NamedExec(q, map[string]interface{}{
		"name_ru":        r.FormValue("name_ru"),
		"description_ru": r.FormValue("description_ru"),
		"id":             vars["id"],
	})
	raven.ReportIfError(err)
	e, _ := m.EmployeeById(vars["emp_id"])
	fNames := u.FileSave(r, "reward", e.PassportSerial)
	for i := range fNames {
		if fNames[i] != "" {
			_, err = s.UDB.NamedExec(`insert into employee__reward__file (file, reward_id) values (:file, :reward_id)`,
				map[string]interface{}{
					"file":      fNames[i],
					"reward_id": vars["id"],
				})
			raven.ReportIfError(err)
		}
	}
}
