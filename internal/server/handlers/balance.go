package handlers

import (
	"github.com/gin-gonic/gin"
	"go_graduation/internal/database"
	"go_graduation/internal/security"
	"net/http"
)

func Balance(c *gin.Context) {
	//GET — получение текущего баланса счёта баллов лояльности пользователя;

	var body struct {
		Order string `json:"order"`
		Sum   int    `json:"sum"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	valid, err := security.Engine.ValidateOrder(body.Order)
	if err != nil {
		c.String(http.StatusInternalServerError,
			"[server:utils:order_post] while validating\n%w", err)
		return
	}
	if !valid {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	uid, _ := c.Get("uid")
	accural, err := database.Controller.GetAccural(c, uid.(string))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	if accural.Current < body.Sum {

	}

	err = database.Controller.AddWithdrawal(c, uid.(string), body.Order, body.Sum)
	if err != nil {
		c.String(http.StatusInternalServerError,
			"[server:utils:order_post] while checking existing record\n%w", err)
		return
	}
	//c.JSON(http.StatusOK, orders)
	c.Status(200)
}
