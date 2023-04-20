package cfg

import (
	"github.com/sakirsensoy/genv"
	_ "github.com/sakirsensoy/genv/dotenv/autoload"
)

var (
	Server   serverCfg
	Database databaseCfg
)

func Initialize() {
	Server = serverCfg{
		Address: genv.Key(
			"RUN_ADDRESS").Default("localhost").String(),
		CookieLifetime: 30 * 24 * 60 * 60,
		AccrualSystemAddress: genv.Key(
			"SERVER_PORT").Default("8080").String(),
	}
	Database = databaseCfg{
		DatabaseUri: genv.Key(
			"DATABASE_URI ").Default("").String(),
	}
}
