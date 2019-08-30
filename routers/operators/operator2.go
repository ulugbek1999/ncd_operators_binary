package operators

import (
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"html/template"
	m "ncd_operators/models"
	"ncd_operators/pkg/raven"
	s "ncd_operators/pkg/settings"
	"net/http"
)

func Operator2Page(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		user = context.Get(r, "user").(*m.User)
		vars = mux.Vars(r)
		e    *m.Employee
	)

	e, err = m.EmployeeById(vars["id"])
	if err != nil {
		w.WriteHeader(404)
		return
	}

	data := map[string]interface{}{
		"user":     user,
		"employee": e,
	}
	tpl := template.Must(template.ParseFiles("templates/operator2/operator.html"))
	err = tpl.ExecuteTemplate(w, "operator_2", data)
	raven.ReportIfError(err)
	return
}

func Operator2_1Employees(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		user = context.Get(r, "user").(*m.User)
		//vars      = mux.Vars(r)
		employees []m.Employee
	)

	q := `SELECT e.id, e.full_name_ru, e.full_name_en, e.register_number, e.passport_serial, e.phone, e.created
	FROM employee e WHERE e.group_id = $1 AND e.step_finished = 1 ORDER BY -id`
	err = s.UDB.Select(&employees, q, user.Operator.OperatorGroup.Id)
	if err != nil {
		raven.ReportIfError(err)
		w.Write([]byte(err.Error()))
	}

	data := map[string]interface{}{
		"user":      user,
		"employees": employees,
	}

	tpl := template.Must(template.ParseFiles("templates/operator2/operator_1.html"))
	err = tpl.ExecuteTemplate(w, "operator_2_1", data)
	if err != nil {
		raven.ReportIfError(err)
		_, err = w.Write([]byte(err.Error()))
		raven.ReportIfError(err)
	}
	return
}
