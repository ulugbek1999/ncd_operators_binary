package api

import (
	"encoding/json"
	m "ncd_operators/models"
	"ncd_operators/pkg/raven"
	s "ncd_operators/pkg/settings"
	"net/http"
)

func CitiesListAPI(w http.ResponseWriter, r *http.Request) {
	var (
		err    error
		cities []m.City
	)
	err = s.UDB.Select(&cities, `SELECT c.id, c.name_ru FROM directory_city c`)
	if err != nil {
		raven.ReportIfError(err)
		w.Write([]byte(err.Error()))
		return
	}
	for i := range cities {
		var d []m.District
		err = s.UDB.Select(&d, `SELECT d.id, d.name_ru FROM directory_district d WHERE city_id = $1`, cities[i].Id)
		if err != nil {
			raven.ReportIfError(err)
			return
		}
		cities[i].Districts = d
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&cities)
	raven.ReportIfError(err)
}
