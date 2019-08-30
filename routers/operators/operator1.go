package operators

import (
	"github.com/gorilla/context"
	"html/template"
	m "ncd_operators/models"
	"ncd_operators/pkg/raven"
	"net/http"
)

func Operator1Page(w http.ResponseWriter, r *http.Request) {

	c := context.Get(r, "user")
	user := c.(*m.User)

	reg, err := m.GenerateRegisterNumber(user.Operator.OperatorGroup.District.City.Code)
	if err != nil {
		raven.ReportIfError(err)
		w.Write([]byte("Couldn't generate `register number`"))
		return
	}
	data := map[string]interface{}{
		"user":            user,
		"register_number": reg,
	}
	tpl := template.Must(template.ParseFiles("templates/operator1/operator.html"))
	err = tpl.ExecuteTemplate(w, "operator_1", data)
	raven.ReportIfError(err)
	return
}
