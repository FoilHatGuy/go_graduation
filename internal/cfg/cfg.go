package cfg

import (
	"github.com/sakirsensoy/genv"
	"github.com/sakirsensoy/genv/dotenv"
)

var (
	Server   serverCfg
	Database databaseCfg
)

func Initialize() {
	err := dotenv.Load("./.env")
	if err != nil {
		return
	}

	Server = serverCfg{
		Address: genv.Key(
			"RUN_ADDRESS").Default("localhost:1000").String(),
		CookieLifetime: 30 * 24 * 60 * 60,
		AccrualSystemAddress: genv.Key(
			"ACCRUAL_SYSTEM_ADDRESS").Default("localhost:1000").String(),
	}
	Database = databaseCfg{
		DatabaseUri: genv.Key(
			"DATABASE_URI").Default("").String(),
	}
}
