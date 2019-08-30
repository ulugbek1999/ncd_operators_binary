package models

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"ncd_operators/pkg/raven"
	s "ncd_operators/pkg/settings"
)

func InitDB() {
	var (
		err error
	)
	s.DB, err = sqlx.Connect("postgres",
		fmt.Sprintf("dbname=%s user=%s password=%s sslmode=disable",
			s.Conf["DBNAME"], s.Conf["DBUSER"], s.Conf["DBPSSWD"]))
	if err != nil {
		raven.ReportIfError(err)
		log.Fatalf("%+v", err)
	}
	if err = s.DB.Ping(); err != nil {
		raven.ReportIfError(err)
		log.Fatalf("%+v", err)
	}
	s.UDB = s.DB.Unsafe()
	log.Println("connected db...")
}
