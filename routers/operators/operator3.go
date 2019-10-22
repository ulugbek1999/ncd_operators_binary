package operators

import (
	"html/template"
	"log"
	m "ncd_operators/models"
	"ncd_operators/pkg/raven"
	s "ncd_operators/pkg/settings"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func Operator3Page(w http.ResponseWriter, r *http.Request) {

	//experience []m.Experience
	var (
		user        = context.Get(r, "user").(*m.User)
		vars        = mux.Vars(r)
		e           m.Employee
		err         error
		educations  []m.Education
		languages   []m.Language
		army        []m.Army
		rewards     []m.Reward
		family      []m.Family
		relatives   []m.Relative
		experiences []m.Experience
	)
	// get employee
	err = s.UDB.Get(&e, `SELECT e.id, e.register_number FROM employee e WHERE e.id = $1`, vars["id"])
	if err != nil {
		raven.ReportIfError(err)
		w.WriteHeader(404)
		return
	}
	// get educations list
	err = s.UDB.Select(&educations, `SELECT e.id, e.name_ru, e.address_ru, e.specialization_ru, e.additional_ru,
	e.date_started, e.date_finished, et.name_ru AS "education_type.name_ru" FROM employee__education e 
	INNER JOIN directory_education_type et ON e.type_id = et.id WHERE e.employee_id = $1`, vars["id"])
	raven.ReportIfError(err)
	for i := range educations {
		var eduFiles []m.EducationFile
		err = s.UDB.Select(&eduFiles, `select file from employee__education__file where education_id = $1`, educations[i].Id)
		if err != nil {
			raven.ReportIfError(err)
			continue
		}
		educations[i].Files = eduFiles
	}
	// get languages list
	err = s.UDB.Select(&languages, `SELECT l.id, l.level, l.is_required_to_check, dl.name_ru AS "dlanguage.name_ru" 
	FROM employee__language l INNER JOIN directory_language dl on l.language_id = dl.id WHERE l.employee_id = $1`, vars["id"])
	if err != nil {
		log.Println(err.Error())
	}
	for i := range languages {
		var langFiles []m.LanguageFile
		err = s.UDB.Select(&langFiles, `select file from employee__language__file where language_id = $1`, languages[i].Id)
		if err != nil {
			raven.ReportIfError(err)
			continue
		}
		languages[i].Files = langFiles
	}
	// get army list
	err = s.UDB.Select(&army, `SELECT a.id, a.name_ru, a.region_ru, a.specialization_ru, a.date_started, a.date_finished, a.rank_ru
		FROM employee__army a WHERE a.employee_id = $1`, vars["id"])
	raven.ReportIfError(err)
	for i := range army {
		var armyFiles []m.ArmyFile
		err = s.UDB.Select(&armyFiles, `select file from employee__army__file where army_id = $1`, army[i].Id)
		if err != nil {
			raven.ReportIfError(err)
			continue
		}
		army[i].Files = armyFiles
	}
	// get reward list
	err = s.UDB.Select(&rewards, `SELECT r.id, r.name_ru, r.description_ru FROM employee__reward r WHERE r.employee_id = $1`, vars["id"])
	raven.ReportIfError(err)
	for i := range rewards {
		var rewardFiles []m.RewardFile
		err = s.UDB.Select(&rewardFiles, `select file from employee__reward__file where reward_id = $1`, rewards[i].Id)
		if err != nil {
			raven.ReportIfError(err)
			continue
		}
		rewards[i].Files = rewardFiles
	}
	// get family list
	err = s.UDB.Select(&family, `SELECT f.id, f.status, f.children_amount FROM employee__family f WHERE f.employee_id = $1`, vars["id"])
	raven.ReportIfError(err)
	for i := range family {
		var famFiles []m.FamilyFile
		err = s.UDB.Select(&famFiles, `select file from employee__family__file where family_id = $1`, family[i].Id)
		if err != nil {
			raven.ReportIfError(err)
			continue
		}
		family[i].Files = famFiles
	}
	// get relative list
	err = s.UDB.Select(&relatives, `SELECT r.id, r.level, r.fullname_ru, r.birth_date, r.study_work_place_ru, r.position_ru, r.living_address_ru 
		FROM employee__relative r WHERE r.employee_id = $1`, vars["id"])
	raven.ReportIfError(err)
	for i := range relatives {
		var relFiles []m.RelativeFile
		err = s.UDB.Select(&relFiles, `select file from employee__relative__file where relative_id = $1`, relatives[i].Id)
		if err != nil {
			raven.ReportIfError(err)
			continue
		}
		relatives[i].Files = relFiles
	}
	// get experiences list
	err = s.UDB.Select(&experiences, `SELECT e.id, e.organization_ru, e.date_started, e.date_finished, e.position_ru, e.address_ru,
       e.sub_division_ru FROM employee__experience e where e.employee_id = $1`, vars["id"])
	raven.ReportIfError(err)
	for i := range experiences {
		var expFiles []m.ExperienceFile
		err = s.UDB.Select(&expFiles, `select file from employee__experience__file where experience_id = $1`, experiences[i].Id)
		if err != nil {
			raven.ReportIfError(err)
			continue
		}
		experiences[i].Files = expFiles
	}
	data := map[string]interface{}{
		"user":        user,
		"employee":    e,
		"educations":  educations,
		"languages":   languages,
		"army":        army,
		"rewards":     rewards,
		"family":      family,
		"relatives":   relatives,
		"experiences": experiences,
		"mediaUrl":    s.Conf["MEDIA_URL"],
	}
	tpl := template.Must(template.ParseFiles("templates/operator3/operator.html"))
	err = tpl.ExecuteTemplate(w, "operator_3", data)
	if err != nil {
		raven.ReportIfError(err)
		_, err = w.Write([]byte(err.Error()))
		raven.ReportIfError(err)
	}
	return
}

func Operator3_2Employees(w http.ResponseWriter, r *http.Request) {
	var (
		user      = context.Get(r, "user").(*m.User)
		employees []m.Employee
	)

	q := `SELECT e.id, e.full_name_ru, e.full_name_en, e.register_number, e.passport_serial, e.phone, e.created
	FROM employee e WHERE e.group_id = $1 AND e.step_finished = 2 ORDER BY -id`
	err := s.UDB.Select(&employees, q, user.Operator.OperatorGroup.Id)
	raven.ReportIfError(err)
	data := map[string]interface{}{
		"user":      user,
		"employees": employees,
	}
	tpl := template.Must(template.ParseFiles("templates/operator3/operator_2.html"))
	err = tpl.ExecuteTemplate(w, "operator_3_2", data)
	raven.ReportIfError(err)
	return
}

/////////////////////////////////////////////////////////////////

func EducationAddPage(w http.ResponseWriter, r *http.Request) {
	var (
		err     error
		e       *m.Employee
		eduType []m.EducationType
		vars    = mux.Vars(r)
		user    = context.Get(r, "user").(*m.User)
	)
	// get employee
	e, err = m.EmployeeById(vars["id"])
	if err != nil {
		raven.ReportIfError(err)
		w.WriteHeader(404)
		return
	}
	// get education type list
	err = s.UDB.Select(&eduType, `SELECT e.id, e.name_ru FROM directory_education_type e ORDER BY -id`)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	data := map[string]interface{}{
		"user":          user,
		"employee":      e,
		"educationType": eduType,
	}
	tpl := template.Must(template.ParseFiles("templates/operator3/education_add.html"))
	err = tpl.ExecuteTemplate(w, "operator_3_education", data)
	raven.ReportIfError(err)
	return
}

func ArmyAddPage(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		e    *m.Employee
		vars = mux.Vars(r)
		user = context.Get(r, "user").(*m.User)
	)

	// get employee
	e, err = m.EmployeeById(vars["id"])
	if err != nil {
		raven.ReportIfError(err)
		w.WriteHeader(404)
		return
	}
	data := map[string]interface{}{
		"user":     user,
		"employee": e,
	}
	tpl := template.Must(template.ParseFiles("templates/operator3/army_add.html"))
	err = tpl.ExecuteTemplate(w, "operator_3_army", data)
	raven.ReportIfError(err)
	return
}

func LanguageAddPage(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		e    *m.Employee
		dL   []m.DLanguage
		vars = mux.Vars(r)
		user = context.Get(r, "user").(*m.User)
	)

	// get employee
	e, err = m.EmployeeById(vars["id"])
	if err != nil {
		raven.ReportIfError(err)
		w.WriteHeader(404)
		return
	}
	// get dlanguages
	err = s.UDB.Select(&dL, `SELECT l.id, l.name_ru FROM directory_language l ORDER BY -id`)
	if err != nil {
		raven.ReportIfError(err)
		w.WriteHeader(404)
		return
	}
	data := map[string]interface{}{
		"user":       user,
		"employee":   e,
		"dlanguages": dL,
	}

	tpl := template.Must(template.ParseFiles("templates/operator3/language_add.html"))
	err = tpl.ExecuteTemplate(w, "operator_3_language", data)
	raven.ReportIfError(err)
	return
}

func FamilyAddPage(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		e    *m.Employee
		vars = mux.Vars(r)
		user = context.Get(r, "user").(*m.User)
	)

	// get employee
	e, err = m.EmployeeById(vars["id"])
	if err != nil {
		raven.ReportIfError(err)
		w.WriteHeader(404)
		return
	}
	data := map[string]interface{}{
		"user":     user,
		"employee": e,
	}

	tpl := template.Must(template.ParseFiles("templates/operator3/family_add.html"))
	err = tpl.ExecuteTemplate(w, "operator_3_family", data)
	raven.ReportIfError(err)
	return
}

func RewardAddPage(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		e    *m.Employee
		vars = mux.Vars(r)
		user = context.Get(r, "user").(*m.User)
	)

	// get employee
	e, err = m.EmployeeById(vars["id"])
	if err != nil {
		raven.ReportIfError(err)
		w.WriteHeader(404)
		return
	}
	data := map[string]interface{}{
		"user":     user,
		"employee": e,
	}
	tpl := template.Must(template.ParseFiles("templates/operator3/reward_add.html"))
	err = tpl.ExecuteTemplate(w, "operator_3_reward", data)
	raven.ReportIfError(err)
	return
}

func RelativeAddPage(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		e    *m.Employee
		vars = mux.Vars(r)
		user = context.Get(r, "user").(*m.User)
	)

	// get employee
	e, err = m.EmployeeById(vars["id"])
	if err != nil {
		raven.ReportIfError(err)
		w.WriteHeader(404)
		return
	}
	data := map[string]interface{}{
		"user":     user,
		"employee": e,
	}
	tpl := template.Must(template.ParseFiles("templates/operator3/relative_add.html"))
	err = tpl.ExecuteTemplate(w, "operator_3_relative", data)
	raven.ReportIfError(err)
	return
}

func ExperienceAddPage(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		e    *m.Employee
		vars = mux.Vars(r)
		user = context.Get(r, "user").(*m.User)
	)

	// get employee
	e, err = m.EmployeeById(vars["id"])
	if err != nil {
		raven.ReportIfError(err)
		w.WriteHeader(404)
		return
	}
	data := map[string]interface{}{
		"user":     user,
		"employee": e,
	}
	tpl := template.Must(template.ParseFiles("templates/operator3/experience_add.html"))
	err = tpl.ExecuteTemplate(w, "operator_3_experience", data)
	raven.ReportIfError(err)
	return
}

// EDIT

func EducationEditPage(w http.ResponseWriter, r *http.Request) {
	var (
		err      error
		vars     = mux.Vars(r)
		edu      m.Education
		user     = context.Get(r, "user").(*m.User)
		eduType  []m.EducationType
		eduFiles []m.EducationFile
	)
	e, _ := m.EmployeeById(vars["emp_id"])
	err = s.UDB.Get(&edu, `select e.id, e.name_ru, e.address_ru, e.specialization_ru, e.date_started, e.date_finished,
       e.additional_ru, et.id as "education_type.id" from employee__education e inner join directory_education_type et
           on e.type_id = et.id where e.id = $1`, vars["id"])
	raven.ReportIfError(err)
	err = s.UDB.Select(&eduFiles, `select file from employee__education__file where education_id = $1`, vars["id"])
	raven.ReportIfError(err)
	edu.Files = eduFiles
	err = s.UDB.Select(&eduType, `SELECT e.id, e.name_ru FROM directory_education_type e ORDER BY -id`)
	raven.ReportIfError(err)
	data := map[string]interface{}{
		"education":     edu,
		"user":          user,
		"educationType": eduType,
		"employee":      e,
	}
	tpl := template.Must(template.ParseFiles("templates/operator3/education_edit.html"))
	err = tpl.ExecuteTemplate(w, "education_edit", data)
	raven.ReportIfError(err)
}

func ArmyEditPage(w http.ResponseWriter, r *http.Request) {
	var (
		err       error
		vars      = mux.Vars(r)
		army      m.Army
		armyFiles []m.ArmyFile
		user      = context.Get(r, "user").(*m.User)
	)
	e, _ := m.EmployeeById(vars["emp_id"])
	err = s.UDB.Get(&army, `select id, name_ru, region_ru, specialization_ru, date_started,
       date_finished, rank_ru from employee__army where id = $1`, vars["id"])
	raven.ReportIfError(err)
	err = s.UDB.Select(&armyFiles, `select file from employee__army__file where army_id = $1`, vars["id"])
	raven.ReportIfError(err)
	army.Files = armyFiles
	data := map[string]interface{}{
		"army":     army,
		"user":     user,
		"employee": e,
	}
	tpl := template.Must(template.ParseFiles("templates/operator3/army_edit.html"))
	err = tpl.ExecuteTemplate(w, "army_edit", data)
	raven.ReportIfError(err)
}

func LanguageEditPage(w http.ResponseWriter, r *http.Request) {
	var (
		err       error
		vars      = mux.Vars(r)
		lang      m.Language
		langFiles []m.LanguageFile
		dLang     []m.DLanguage
		user      = context.Get(r, "user").(*m.User)
	)
	e, _ := m.EmployeeById(vars["emp_id"])
	err = s.UDB.Get(&lang, `select l.id, l.level, l.is_required_to_check, dl.id as "dlanguage.id", dl.name_ru as "dlanguage.name_ru"
		from employee__language l inner join directory_language dl on l.language_id = dl.id where l.id = $1`, vars["id"])
	raven.ReportIfError(err)
	err = s.UDB.Select(&langFiles, `select file from employee__language__file where language_id = $1`, vars["id"])
	raven.ReportIfError(err)
	lang.Files = langFiles
	err = s.UDB.Select(&dLang, `SELECT l.id, l.name_ru FROM directory_language l ORDER BY -id`)
	raven.ReportIfError(err)
	data := map[string]interface{}{
		"language":   lang,
		"dlanguages": dLang,
		"user":       user,
		"employee":   e,
	}
	tpl := template.Must(template.ParseFiles("templates/operator3/language_edit.html"))
	err = tpl.ExecuteTemplate(w, "language_edit", data)
	raven.ReportIfError(err)
}

func FamilyEditPage(w http.ResponseWriter, r *http.Request) {
	var (
		err        error
		vars       = mux.Vars(r)
		family     m.Family
		familyFile []m.FamilyFile
		user       = context.Get(r, "user").(*m.User)
	)
	e, _ := m.EmployeeById(vars["emp_id"])
	err = s.UDB.Get(&family, `select id, status, children_amount from employee__family where id = $1`, vars["id"])
	raven.ReportIfError(err)
	err = s.UDB.Select(&familyFile, `select file from employee__family__file where family_id = $1`, vars["id"])
	raven.ReportIfError(err)
	family.Files = familyFile
	data := map[string]interface{}{
		"family":   family,
		"user":     user,
		"employee": e,
	}
	tpl := template.Must(template.ParseFiles("templates/operator3/family_edit.html"))
	err = tpl.ExecuteTemplate(w, "family_edit", data)
	raven.ReportIfError(err)
}

func RewardEditPage(w http.ResponseWriter, r *http.Request) {
	var (
		err        error
		vars       = mux.Vars(r)
		reward     m.Reward
		rewardFile []m.RewardFile
		user       = context.Get(r, "user").(*m.User)
	)
	e, _ := m.EmployeeById(vars["emp_id"])
	err = s.UDB.Get(&reward, `select id, name_ru, description_ru
		from employee__reward where id = $1`, vars["id"])
	raven.ReportIfError(err)
	err = s.UDB.Select(&rewardFile, `select file from employee__reward__file where reward_id = $1`, vars["id"])
	raven.ReportIfError(err)
	reward.Files = rewardFile
	data := map[string]interface{}{
		"reward":   reward,
		"user":     user,
		"employee": e,
	}
	tpl := template.Must(template.ParseFiles("templates/operator3/reward_edit.html"))
	err = tpl.ExecuteTemplate(w, "reward_edit", data)
	raven.ReportIfError(err)
}

func RelativeEditPage(w http.ResponseWriter, r *http.Request) {
	var (
		err          error
		vars         = mux.Vars(r)
		relative     m.Relative
		relativeFile []m.RelativeFile
		user         = context.Get(r, "user").(*m.User)
	)
	e, _ := m.EmployeeById(vars["emp_id"])
	err = s.UDB.Get(&relative, `select id, level, fullname_ru, birth_date, study_work_place_ru, position_ru, living_address_ru
		from employee__relative where id = $1`, vars["id"])
	raven.ReportIfError(err)
	err = s.UDB.Select(&relativeFile, `select file from employee__relative__file where relative_id = $1`, vars["id"])
	raven.ReportIfError(err)
	relative.Files = relativeFile
	data := map[string]interface{}{
		"relative": relative,
		"user":     user,
		"employee": e,
	}
	tpl := template.Must(template.ParseFiles("templates/operator3/relative_edit.html"))
	err = tpl.ExecuteTemplate(w, "relative_edit", data)
	raven.ReportIfError(err)
}

func ExperienceEditPage(w http.ResponseWriter, r *http.Request) {
	var (
		err            error
		vars           = mux.Vars(r)
		experience     m.Experience
		experienceFile []m.ExperienceFile
		user           = context.Get(r, "user").(*m.User)
	)
	e, _ := m.EmployeeById(vars["emp_id"])
	err = s.UDB.Get(&experience, `select id, organization_ru, date_started, date_finished, position_ru,
       sub_division_ru, address_ru from employee__experience where id = $1`, vars["id"])
	raven.ReportIfError(err)
	err = s.UDB.Select(&experienceFile, `select file from employee__experience__file where experience_id = $1`, vars["id"])
	raven.ReportIfError(err)
	experience.Files = experienceFile
	data := map[string]interface{}{
		"experience": experience,
		"user":       user,
		"employee":   e,
	}
	tpl := template.Must(template.ParseFiles("templates/operator3/experience_edit.html"))
	err = tpl.ExecuteTemplate(w, "experience_edit", data)
	raven.ReportIfError(err)
}
