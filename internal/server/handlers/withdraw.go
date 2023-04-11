package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"go_graduation/internal/database"
	"go_graduation/internal/security"
	"io"
	"net/http"
)

func Withdraw(c *gin.Context) {
	//POST — запрос на списание баллов с накопительного счёта в счёт оплаты нового заказа;

	buf, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	input := string(buf)
	valid, err := security.Engine.ValidateOrder(input)
	if err != nil {
		c.String(http.StatusInternalServerError,
			"[server:utils:order_post] while validating\n%w", err)
		return
	}
	if !valid {
		c.Status(http.StatusUnprocessableEntity)
		return
	}
	order, err := database.Controller.GetOrder(c, input)
	uid, _ := c.Get("uid")
	if err != nil {
		c.String(http.StatusInternalServerError,
			"[server:utils:order_post] while checking existing record\n%w", err)
		return
	}
	if errors.Is(err, pgx.ErrNoRows) {
		if order.User != uid {
			c.String(http.StatusConflict,
				"already exists, other user")
			return
		} else {
			c.String(http.StatusOK,
				"already exists")
			return
		}
	}
	err = database.Controller.AddOrder(c, input)

	if err != nil {
		c.String(http.StatusInternalServerError,
			"[server:utils:order_post] while checking existing record\n%w", err)
		return
	}
	c.Status(http.StatusAccepted)
}
