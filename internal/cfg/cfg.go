package cfg

import (
	"flag"
	"github.com/sakirsensoy/genv"
	"github.com/sakirsensoy/genv/dotenv"
)

var (
	Server         serverCfg
	Database       databaseCfg
	databaseDSN    string
	serverAddress  string
	accuralAddress string
)

func Initialize() {
	flag.StringVar(&databaseDSN, "d", "", "")
	flag.StringVar(&serverAddress, "a", "localhost:1080", "")
	flag.StringVar(&accuralAddress, "b", "http://localhost:8080", "")
	flag.Parse()

	err := dotenv.Load("./.env")
	if err != nil {
		return
	}

	Server = serverCfg{
		Address: genv.Key(
			"RUN_ADDRESS").Default(serverAddress).String(),
		CookieLifetime: 30 * 24 * 60 * 60,
		AccrualSystemAddress: genv.Key(
			"ACCRUAL_SYSTEM_ADDRESS").Default(accuralAddress).String(),
	}
	Database = databaseCfg{
		DatabaseUri: genv.Key(
			"DATABASE_URI").Default(databaseDSN).String(),
	}
}
