package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"ncd_operators/pkg/raven"
	s "ncd_operators/pkg/settings"
	"net/http"
	"strings"
)

func Operator4SaveAPI(w http.ResponseWriter, r *http.Request) {
	var (
		vars = mux.Vars(r)
		err  error
	)
	q := `UPDATE employee SET (step_finished, wages, is_ready_for_university, criminal_number, criminal_issue, criminal_comment_ru,
		  add_spec_signs_ru, hobby_ru, other_ru, psycholog, busy) = (:step_finished, :wages, :is_ready_for_university,
		  :criminal_number, :criminal_issue, :criminal_comment_ru, :add_spec_signs_ru, :hobby_ru,
		  :other_ru, :psycholog, :busy) WHERE id = :employee_id`
	_, err = s.UDB.NamedExec(q, map[string]interface{}{
		"step_finished":           4,
		"wages":                   r.FormValue("wages"),
		"is_ready_for_university": r.FormValue("is_ready_for_university"),
		"criminal_number":         r.FormValue("criminal_number"),
		"criminal_issue":          r.FormValue("criminal_issue"),
		"criminal_comment_ru":     r.FormValue("criminal_comment_ru"),
		"add_spec_signs_ru":       r.FormValue("add_spec_signs_ru"),
		"hobby_ru":                r.FormValue("hobby_ru"),
		"other_ru":                r.FormValue("other_ru"),
		"psycholog":               r.FormValue("psycholog"),
		"employee_id":             vars["id"],
		"busy":                    false,
	})
	if err != nil {
		raven.ReportIfError(err)
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	q = fmt.Sprintf("UPDATE employee SET %s = true", r.FormValue("level"))
	_, err = s.UDB.Exec(q)
	if err != nil {
		raven.ReportIfError(err)
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	countries := r.FormValue("country")
	c := strings.Split(countries, ",")
	for i := range c {
		_, err = s.UDB.Exec(`insert into employee__countries (country_id, employee_id) VALUES ($1, $2)`, c[i], vars["id"])
		raven.ReportIfError(err)
	}
	w.WriteHeader(200)
	return
}
