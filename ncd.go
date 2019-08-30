package main

import (
	"log"
	"ncd_operators/cmd"
	m "ncd_operators/models"
	"ncd_operators/pkg/raven"
	s "ncd_operators/pkg/settings"
	"os"

	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

func init() {
	sentry.Init(sentry.ClientOptions{
		Dsn: "https://7b2a6b831f744f0786eafcf341c9e857@sentry.io/1521097",
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
	app := cli.NewApp()
	app.Name = "NCD Operators module"
	app.Description = ""
	app.Usage = ""
	app.Version = s.APP_VERSION
	app.Commands = []cli.Command{
		cmd.RunServer,
	}
	err := app.Run(os.Args)
	raven.ReportIfError(err)
}
