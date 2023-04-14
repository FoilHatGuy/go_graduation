package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"go_graduation/internal/cfg"
)

func databaseInitialize() *pgxpool.Pool {
	cfg.Initialize()
	fmt.Println(cfg.Database.DatabaseUri)
	db, err := pgxpool.New(context.Background(), cfg.Database.DatabaseUri)
	if err != nil && !errors.Is(err, new(pgconn.PgError)) {
		return nil
	}

	exec, err := db.Exec(context.Background(), `
	CREATE TABLE IF 	NOT EXISTS users (
	    login 			text 	UNIQUE NOT NULL,
	    password		text	NOT NULL
	                                )
`)
	fmt.Println(exec)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	exec, err = db.Exec(context.Background(), `
	DO $$
	BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'ord_stat') THEN
		CREATE TYPE ord_stat
			AS ENUM ('REGISTERED', 'INVALID', 'PROCESSING', 'PROCESSED');
		END IF;
	END
	$$;
`)
	fmt.Println(exec)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	exec, err = db.Exec(context.Background(), `
	CREATE TABLE IF NOT EXISTS orders (
	    order_id 		text 		UNIQUE NOT NULL, 
	    creation_date	timestamp	NOT NULL DEFAULT NOW(), 
		status		 	ord_stat 	NOT NULL DEFAULT 'REGISTERED', 
	    user_id			text 		NOT NULL,
	    FOREIGN KEY (user_id)
	    	REFERENCES users (login)
	    	ON DELETE CASCADE
	                                )
`)
	fmt.Println(exec)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	exec, err = db.Exec(context.Background(), `
	CREATE TABLE IF NOT EXISTS withdrawals (
	    order_id 		text 		UNIQUE NOT NULL, 
	    accural			integer		NOT NULL,			
	    processed_at	timestamp	NOT NULL DEFAULT NOW(), 
	    FOREIGN KEY (order_id)
	    	REFERENCES orders (order_id)
	    	ON DELETE CASCADE
	                                )
`)
	fmt.Println(exec)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	exec, err = db.Exec(context.Background(), `
	CREATE TABLE IF 	NOT EXISTS accural (
	    user_id			text 	UNIQUE NOT NULL PRIMARY KEY, 
	    accural			text 	UNIQUE NOT NULL,
	    withdrawn		text	NOT NULL,			
	    FOREIGN KEY (user_id)
	    	REFERENCES users (login)
	    	ON DELETE CASCADE
	                                )
`)
	fmt.Println(exec)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	exec, err = db.Exec(context.Background(), `
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
		return nil
	}
	return db
}
