package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"go_graduation/internal/cfg"
)

type databaseT struct {
	database *pgxpool.Pool
}

func databaseInitialize() *databaseT {
	db, err := pgxpool.New(context.Background(), cfg.Database.DatabaseUri)
	if err != nil && !errors.Is(err, new(pgconn.PgError)) {
		return nil
	}
	res := new(databaseT)
	res.database = db
	res.initialize()
	return res
}
func (d databaseT) initialize() {
	exec, err := d.database.Exec(context.Background(), `
	CREATE TABLE IF 	NOT EXISTS users (
	    login 			text 	UNIQUE NOT NULL,
	    password		text	NOT NULL
	                                )

	CREATE TYPE ord_stat AS ENUM ('REGISTERED', 'INVALID', 'PROCESSING', 'PROCESSED');
	CREATE TABLE IF NOT EXISTS orders (
	    order_id 		text 		UNIQUE NOT NULL, 
	    creation_date	text 		NOT NULL, 
		status		 	ord_stat 	NOT NULL, 
	    user_id			text 		NOT NULL,
	    accural			integer,			
	    FOREIGN KEY (user_id)
	    	REFERENCES users (login)
	    	ON DELETE CASCADE
	                                )

	CREATE TABLE IF 	NOT EXISTS accural (
	    user_id			text 	UNIQUE NOT NULL PRIMARY KEY, 
	    accural			text 	UNIQUE NOT NULL,
	    withdrawn		text	NOT NULL,			
	    FOREIGN KEY (user_id)
	    	REFERENCES users (login)
	    	ON DELETE CASCADE
	                                )

	CREATE TABLE IF 	NOT EXISTS sessions (
	    user_id			text 	UNIQUE NOT NULL, 
	    cookie			text 	UNIQUE NOT NULL,
	    FOREIGN KEY (user_id)
	    	REFERENCES users (login)
	    	ON DELETE CASCADE
	                                )
`)
	fmt.Println(exec)
	if err != nil {
		fmt.Println(err)
		return
	}
}
