package database

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type controllerT struct {
	db *pgxpool.Pool
}

var Controller controllerT

func init() {
	Controller = controllerT{
		db: databaseInitialize(),
	}
}

func (c *controllerT) AddUser(ctx context.Context, login string, password string) error {
	exec, err := c.db.Exec(ctx, `
			INSERT INTO users VALUES ($1, $2)
		`, login, password)
	fmt.Println(exec)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (c *controllerT) GetUser(ctx context.Context, login string) (password string, err error) {
	return "", nil
}

func (c *controllerT) AddSession(ctx context.Context, uid string, cookie string) (err error) {
	exec, err := c.db.Exec(ctx, `
			INSERT INTO sessions VALUES ($1, $2)
		`, uid, cookie)
	fmt.Println(exec)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (c *controllerT) GetSession(ctx context.Context, sid string) (login string, err error) {
	err = c.db.QueryRow(ctx, `SELECT user_id FROM sessions WHERE cookie=$1`, sid).Scan(&login)
	if err != nil {
		return "", err
	}
	return login, nil
}

func (c *controllerT) AddOrder(ctx context.Context, uid string, oid string) (err error) {
	exec, err := c.db.Exec(ctx, `
			INSERT INTO orders (order_id, user_id)
			VALUES ($1, $2)
		`, oid, uid)
	fmt.Println(exec)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}

type order struct {
	User       string
	Status     string
	UploadDate string
}

func (c *controllerT) GetOrder(ctx context.Context, ordNum string) (ord order, err error) {
	var user, status, date string
	err = c.db.QueryRow(ctx, `SELECT * FROM orders WHERE user_id=$1`, ordNum).Scan(&user, &status, &date)
	if err != nil {
		return order{}, err
	}
	return order{user, status, date}, nil
}

func (c *controllerT) GetOrders(ctx context.Context, uid string) (resp []order, err error) {
	rows, err := c.db.Query(ctx, `SELECT * FROM accural WHERE user_id=$1`, uid)
	if err != nil {
		return []order{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var user, status, upload string
		err := rows.Scan(&user, &status, &upload)
		if err != nil {
			return nil, err
		}
		resp = append(resp, order{user, status, upload})

	}
	return []order{}, nil
}

type accural struct {
	Current   int
	Withdrawn int
}

func (c *controllerT) GetAccural(ctx *gin.Context, uid string) (acc accural, err error) {
	var current, withdrawn int
	err = c.db.QueryRow(ctx, `SELECT * FROM accural WHERE user_id=$1`, uid).Scan(&current, &withdrawn)
	if err != nil {
		return accural{}, err
	}
	return accural{0, 0}, nil
}

type withdrawal struct {
	Order        int
	Sum          int
	ProcesssedAt string `json:"processsed_at"`
}

func (c *controllerT) GetWithdrawals(ctx *gin.Context, uid string) (resp []withdrawal, err error) {

	rows, err := c.db.Query(ctx, `SELECT withdrawals.* FROM withdrawals, orders WHERE user_id=$1`, uid)
	if err != nil {
		return []withdrawal{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var ordern, sum int
		var processed string
		err := rows.Scan(&ordern, &sum, &processed)
		if err != nil {
			return nil, err
		}
		resp = append(resp, withdrawal{ordern, sum, processed})

	}
	return resp, nil
}

func (c *controllerT) AddWithdrawal(ctx *gin.Context, uid string, oid string, sum int) (err error) {
	exec, err := c.db.Exec(ctx, `
		INSERT INTO withdrawals VALUES($1, $2)
		`, sum, uid)
	fmt.Println(exec)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
