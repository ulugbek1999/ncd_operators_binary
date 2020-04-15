package main

import (
	"log"
	m "ncd_operators/models"
	"ncd_operators/pkg/raven"
	"ncd_operators/pkg/server"
	s "ncd_operators/pkg/settings"

	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"
)

func init() {
	sentry.Init(sentry.ClientOptions{
		Dsn: "https://c535a7563f14461592d08e77b16b05c6@sentry.io/1831691",
	})
	err := godotenv.Load()
	if err != nil {
		raven.ReportIfError(err)
		log.Fatalln("Error loading .env file")
	}
	s.Conf, err = godotenv.Read()
	raven.ReportIfError(err)
	m.InitDB()
}

func main() {
	server.StartWebServer()
}
