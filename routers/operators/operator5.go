package operators

import (
	"fmt"
	"html/template"
	m "ncd_operators/models"
	"ncd_operators/pkg/raven"
	s "ncd_operators/pkg/settings"
	"net/http"

	"github.com/gorilla/context"
)

//Operator5Page is responsible for handling operator 5 page
func Operator5Page(w http.ResponseWriter, r *http.Request) {
	var (
		err     error
		eduType []m.EducationType
		dL      []m.DLanguage
	)
	c := context.Get(r, "user")
	user := c.(*m.User)

	reg, err := m.GenerateRegisterNumber(user.Operator.OperatorGroup.District.City.Code)
	if err != nil {
		raven.ReportIfError(err)
		w.Write([]byte("Couldn't generate `register number`"))
		return
	}
	err = s.UDB.Select(&eduType, `SELECT id, name_ru FROM directory_education_type ORDER BY -id`)
	if err != nil {
		raven.ReportIfError(err)
		w.Write([]byte("Something went wrong while getting data from directory_education_type"))
	}
	err = s.UDB.Select(&dL, `SELECT id, name_ru FROM directory_language ORDER BY -id`)
	if err != nil {
		raven.ReportIfError(err)
		w.Write([]byte("Cannot get data from directory_language"))
	}
	data := map[string]interface{}{
		"user":            user,
		"register_number": reg,
		"educationType":   eduType,
		"dlanguages":      dL,
		"countries":       m.CountryList(),
	}
	fmt.Println(dL)
	tpl := template.Must(template.ParseFiles("templates/operator5/operator.html"))
	err = tpl.ExecuteTemplate(w, "operator_5", data)
	raven.ReportIfError(err)
	return
}
