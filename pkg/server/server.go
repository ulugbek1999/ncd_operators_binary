package server

import (
	"fmt"
	"log"
	"ncd_operators/middlewares"
	"ncd_operators/pkg/raven"
	s "ncd_operators/pkg/settings"
	"ncd_operators/routers/api"
	"ncd_operators/routers/operators"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/cli"
)

func StartWebServer(c *cli.Context) {
	if c.IsSet("ipaddress") {
		s.IpAddress = c.String("ipaddress")
	}
	if c.IsSet("port") {
		s.Port = c.String("port")
	}
	addr := fmt.Sprintf("%s:%s", s.IpAddress, s.Port)
	if s.Mux == nil {
		s.Mux = mux.NewRouter()
	}
	s.Mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	s.Mux.HandleFunc("/login", operators.Login).Methods("GET", "POST")
	s.Mux.HandleFunc("/logout", operators.Logout).Methods("GET")
	// regions list
	s.Mux.HandleFunc("/api/cities/list", api.CitiesListAPI).Methods("GET")

	s.Mux.HandleFunc("/operator-1", middlewares.AuthMiddleware(operators.Operator1Page)).Methods("GET")

	s.Mux.HandleFunc("/operator-2-1", middlewares.AuthMiddleware(operators.Operator2_1Employees)).Methods("GET")
	s.Mux.HandleFunc("/operator-2/{id}", middlewares.AuthMiddleware(operators.Operator2Page)).Methods("GET")

	s.Mux.HandleFunc("/operator-3-2", middlewares.AuthMiddleware(operators.Operator3_2Employees)).Methods("GET")
	s.Mux.HandleFunc("/operator-3/{id}/add/education", middlewares.AuthMiddleware(operators.EducationAddPage)).Methods("GET")
	s.Mux.HandleFunc("/operator-3/{id}/{emp_id}/edit/education", middlewares.AuthMiddleware(operators.EducationEditPage)).Methods("GET")
	s.Mux.HandleFunc("/operator-3/{id}/add/army", middlewares.AuthMiddleware(operators.ArmyAddPage)).Methods("GET")
	s.Mux.HandleFunc("/operator-3/{id}/{emp_id}/edit/army", middlewares.AuthMiddleware(operators.ArmyEditPage)).Methods("GET")
	s.Mux.HandleFunc("/operator-3/{id}/add/language", middlewares.AuthMiddleware(operators.LanguageAddPage)).Methods("GET")
	s.Mux.HandleFunc("/operator-3/{id}/{emp_id}/edit/language", middlewares.AuthMiddleware(operators.LanguageEditPage)).Methods("GET")
	s.Mux.HandleFunc("/operator-3/{id}/add/family", middlewares.AuthMiddleware(operators.FamilyAddPage)).Methods("GET")
	s.Mux.HandleFunc("/operator-3/{id}/{emp_id}/edit/family", middlewares.AuthMiddleware(operators.FamilyEditPage)).Methods("GET")
	s.Mux.HandleFunc("/operator-3/{id}/add/reward", middlewares.AuthMiddleware(operators.RewardAddPage)).Methods("GET")
	s.Mux.HandleFunc("/operator-3/{id}/{emp_id}/edit/reward", middlewares.AuthMiddleware(operators.RewardEditPage)).Methods("GET")
	s.Mux.HandleFunc("/operator-3/{id}/add/relative", middlewares.AuthMiddleware(operators.RelativeAddPage)).Methods("GET")
	s.Mux.HandleFunc("/operator-3/{id}/{emp_id}/edit/relative", middlewares.AuthMiddleware(operators.RelativeEditPage)).Methods("GET")
	s.Mux.HandleFunc("/operator-3/{id}/add/experience", middlewares.AuthMiddleware(operators.ExperienceAddPage)).Methods("GET")
	s.Mux.HandleFunc("/operator-3/{id}/{emp_id}/edit/experience", middlewares.AuthMiddleware(operators.ExperienceEditPage)).Methods("GET")
	s.Mux.HandleFunc("/operator-3/{id}", middlewares.AuthMiddleware(operators.Operator3Page)).Methods("GET")

	s.Mux.HandleFunc("/operator-4-3", middlewares.AuthMiddleware(operators.Operator4_3Employees)).Methods("GET")
	s.Mux.HandleFunc("/operator-4/{id}", middlewares.AuthMiddleware(operators.Operator4Page)).Methods("GET")

	s.Mux.HandleFunc("/api/operator-1/create", middlewares.AuthMiddleware(api.Operator1CreateAPI)).Methods("POST")
	s.Mux.HandleFunc("/api/operator-2/{id}/create", middlewares.AuthMiddleware(api.Operator2CreateAPI)).Methods("PUT")
	s.Mux.HandleFunc("/api/operator-3/{id}/close", api.Operator3Close).Methods("GET")
	s.Mux.HandleFunc("/api/operator-3/{id}/add/education", api.EducationCreateAPI).Methods("POST")
	s.Mux.HandleFunc("/api/operator-3/{id}/{emp_id}/edit/education", api.EducationEditAPI).Methods("POST")
	s.Mux.HandleFunc("/api/operator-3/{id}/add/language", api.LanguageCreateAPI).Methods("POST")
	s.Mux.HandleFunc("/api/operator-3/{id}/{emp_id}/edit/language", api.LanguageEditAPI).Methods("POST")
	s.Mux.HandleFunc("/api/operator-3/{id}/add/army", api.ArmyCreateAPI).Methods("POST")
	s.Mux.HandleFunc("/api/operator-3/{id}/{emp_id}/edit/army", api.ArmyEditAPI).Methods("POST")
	s.Mux.HandleFunc("/api/operator-3/{id}/add/reward", api.RewardCreateAPI).Methods("POST")
	s.Mux.HandleFunc("/api/operator-3/{id}/{emp_id}/edit/reward", api.RewardEditAPI).Methods("POST")
	s.Mux.HandleFunc("/api/operator-3/{id}/add/family", api.FamilyCreateAPI).Methods("POST")
	s.Mux.HandleFunc("/api/operator-3/{id}/{emp_id}/edit/family", api.FamilyEditAPI).Methods("POST")
	s.Mux.HandleFunc("/api/operator-3/{id}/add/relative", api.RelativeCreateAPI).Methods("POST")
	s.Mux.HandleFunc("/api/operator-3/{id}/{emp_id}/edit/relative", api.RelativeEditAPI).Methods("POST")
	s.Mux.HandleFunc("/api/operator-3/{id}/add/experience", api.ExperienceCreateAPI).Methods("POST")
	s.Mux.HandleFunc("/api/operator-3/{id}/{emp_id}/edit/experience", api.ExperienceEditAPI).Methods("POST")
	s.Mux.HandleFunc("/api/operator-4/{id}/save", api.Operator4SaveAPI).Methods("POST")

	s.Mux.HandleFunc("/", middlewares.AuthMiddleware(operators.IndexPage)).Methods("GET")
	s.Mux.HandleFunc("/operator-universal", middlewares.AuthMiddleware(operators.Operator5Page)).Methods("GET")
	s.Mux.HandleFunc("/api/operator-universal", api.UniversalCreateAPI).Methods("POST")

	log.Printf("WebServer is running on http://%s", addr)
	err := http.ListenAndServe(addr, s.Mux)
	if err != nil {
		raven.ReportIfError(err)
		log.Fatalf("%+v", err)
	}
}
