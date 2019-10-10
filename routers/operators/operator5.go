package operators

import (
	"html/template"
	m "ncd_operators/models"
	"ncd_operators/pkg/raven"
	"net/http"

	"github.com/gorilla/context"
)

//Operator5Page is responsible for handling operator 5 page
func Operator5Page(w http.ResponseWriter, r *http.Request) {
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
	tpl := template.Must(template.ParseFiles("templates/operator5/operator.html"))
	err = tpl.ExecuteTemplate(w, "operator_5", data)
	raven.ReportIfError(err)
	return
}
