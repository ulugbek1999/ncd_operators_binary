package api

import (
	"github.com/gorilla/mux"
	m "ncd_operators/models"
	"ncd_operators/pkg/raven"
	s "ncd_operators/pkg/settings"
	u "ncd_operators/pkg/utils"
	"net/http"
)

func Operator2CreateAPI(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		e    *m.Employee
		vars = mux.Vars(r)
	)
	e, err = m.EmployeeById(vars["id"])
	if err != nil {
		w.WriteHeader(404)
		return
	}
	appearance := u.NewNullInt(r.FormValue("appearance"))
	neatness := u.NewNullInt(r.FormValue("neatness"))
	skin := u.NewNullInt(r.FormValue("skin"))
	height := u.NewNullFloat(r.FormValue("height"))
	weight := u.NewNullFloat(r.FormValue("weight"))
	fatness := u.NewNullFloat(r.FormValue("fatness"))
	bloodGroup := u.NewNullInt(r.FormValue("blood_group"))
	bloodRhesus := r.FormValue("blood_resus")
	visionL := u.NewNullFloat(r.FormValue("vision_l"))
	visionR := u.NewNullFloat(r.FormValue("vision_r"))

	q := `UPDATE employee SET (appearance, neatness, skin, height, weight, fatness, blood_group, blood_resus, vision_l, vision_r, step_finished)
		= (:appearance, :neatness, :skin, :height, :weight, :fatness, :blood_group, :blood_group, :vision_l, :vision_r, :step_finished)
		WHERE id = :id`
	v := map[string]interface{}{
		"step_finished": 2,
		"appearance":    appearance,
		"neatness":      neatness,
		"skin":          skin,
		"height":        height,
		"weight":        weight,
		"fatness":       fatness,
		"blood_group":   bloodGroup,
		"blood_resus":   bloodRhesus,
		"vision_l":      visionL,
		"vision_r":      visionR,
		"id":            vars["id"],
	}
	_, err = s.DB.NamedExec(q, v)
	if err != nil {
		raven.ReportIfError(err)
		w.Write([]byte(err.Error()))
	}
	var fNames []string
	fNames = u.FileSave(r, "photo_1", e.PassportSerial)
	for i := range fNames {
		if fNames[i] != "" {
			_, err = s.UDB.Exec(`UPDATE employee SET photo_1 = $1 WHERE id = $2`, fNames[i], vars["id"])
			raven.ReportIfError(err)
		}
	}

	fNames = u.FileSave(r, "photo_2", e.PassportSerial)
	for i := range fNames {
		if fNames[i] != "" {
			_, err = s.UDB.Exec(`UPDATE employee SET photo_2 = $1 WHERE id = $2`, fNames[i], vars["id"])
			raven.ReportIfError(err)
		}
	}

	fNames = u.FileSave(r, "photo_3", e.PassportSerial)
	for i := range fNames {
		if fNames[i] != "" {
			_, err = s.UDB.Exec(`UPDATE employee SET photo_3 = $1 WHERE id = $2`, fNames[i], vars["id"])
			raven.ReportIfError(err)
		}
	}

	fNames = u.FileSave(r, "photo_4", e.PassportSerial)
	for i := range fNames {
		if fNames[i] != "" {
			_, err = s.UDB.Exec(`UPDATE employee SET photo_4 = $1 WHERE id = $2`, fNames[i], vars["id"])
			raven.ReportIfError(err)
		}
	}
	w.WriteHeader(201)
	return
}
