package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"go_graduation/internal/database"
	"net/http"
)

func OrdersGet(c *gin.Context) {
	//GET — получение списка загруженных пользователем номеров заказов,
	//		статусов их обработки и информации о начислениях;
	uid, _ := c.Get("uid")
	orders, err := database.Controller.GetOrders(c, uid.(string))
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		c.String(http.StatusInternalServerError,
			"[server:utils:order_post] while checking existing record\n%w", err)
		return
	}
	if errors.Is(err, pgx.ErrNoRows) {
		c.Status(http.StatusNoContent)
	}
	c.JSON(http.StatusOK, orders)
}
