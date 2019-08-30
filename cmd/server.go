package cmd

import (
	"github.com/urfave/cli"
	"ncd_operators/pkg/server"
)

var RunServer = cli.Command{
	Name:        "runserver",
	Usage:       "ncd_operators runserver",
	Description: "Starts webserver",
	Action:      server.StartWebServer,
	Flags: []cli.Flag{
		stringFlag("ipaddress, addr", "127.0.0.1", "ip address to run web server"),
		stringFlag("port, p", "8000", "port number to run web server"),
	},
}
