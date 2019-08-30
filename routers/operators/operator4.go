package operators

import (
	"html/template"
	m "ncd_operators/models"
	"ncd_operators/pkg/raven"
	s "ncd_operators/pkg/settings"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func Operator4_3Employees(w http.ResponseWriter, r *http.Request) {
	c := context.Get(r, "user")
	user := c.(*m.User)
	var employees []m.Employee

	q := `SELECT e.id, e.full_name_ru, e.full_name_en, e.register_number, e.passport_serial, e.phone, e.created
	FROM employee e WHERE e.group_id = $1 AND e.step_finished = 3 ORDER BY -id`
	err := s.UDB.Select(&employees, q, user.Operator.OperatorGroup.Id)
	tpl := template.Must(template.ParseFiles("templates/operator4/operator_3.html"))
	data := map[string]interface{}{
		"user":      user,
		"employees": employees,
	}
	err = tpl.ExecuteTemplate(w, "operator_4_3", data)
	raven.ReportIfError(err)
	return
}

func Operator4Page(w http.ResponseWriter, r *http.Request) {
	var (
		user = context.Get(r, "user").(*m.User)
		vars = mux.Vars(r)
		e    *m.Employee
		err  error
	)
	e, err = m.EmployeeById(vars["id"])
	if err != nil {
		raven.ReportIfError(err)
		w.WriteHeader(404)
		return
	}
	data := map[string]interface{}{"user": user, "employee": e, "countries": m.CountryList(), "professions": m.ProfessionAvailableList()}
	tpl := template.Must(template.ParseFiles("templates/operator4/operator.html"))
	err = tpl.ExecuteTemplate(w, "operator_4", data)
	raven.ReportIfError(err)
	return
}
