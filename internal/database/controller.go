package database

import (
	"context"
	"github.com/gin-gonic/gin"
)

type controllerT struct {
	db *databaseT
}

var Controller = controllerT{
	db: databaseInitialize(),
}

func (c *controllerT) AddUser(ctx context.Context, login string, password string) error {
	return nil
}

func (c *controllerT) GetUser(ctx context.Context, login string) (password string, err error) {
	return "", nil
}

func (c *controllerT) AddSession(ctx context.Context, sid string, login string) (err error) {
	return nil
}

func (c *controllerT) GetSession(ctx context.Context, sid string) (login string, err error) {
	return "", nil
}

func (c *controllerT) AddOrder(ctx context.Context, input string) (err error) {
	return nil

}

type order struct {
	User       string
	Status     string
	UploadDate string
}

func (c *controllerT) GetOrder(ctx context.Context, ordNum string) (ord order, err error) {
	return order{"", "", ""}, nil
}

func (c *controllerT) GetOrders(ctx context.Context, uid string) (rows []order, err error) {
	return []order{}, nil
}

type accural struct {
	Current   int
	Withdrawn int
}

func (c *controllerT) GetAccural(ctx *gin.Context, uid string) (acc accural, err error) {
	return accural{0, 0}, nil
}

type withdrawal struct {
	current   int
	withdrawn int
}

func (c *controllerT) GetWithdrawals(ctx *gin.Context, s string) (acc withdrawal, err error) {
	return withdrawal{0, 0}, nil
}

func (c *controllerT) AddWithdrawal(ctx *gin.Context, s string, o string, sum int) (err error) {
	return nil
}
