package settings

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

var (
	Mux       *mux.Router
	IpAddress = "127.0.0.1"
	Port      = "8000"
	DB        *sqlx.DB
	UDB       *sqlx.DB
	Conf      map[string]string
)

const (
	APP_VERSION = "0.1.0"
)
