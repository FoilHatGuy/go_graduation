package handlers

import (
	"github.com/gin-gonic/gin"
	"go_graduation/internal/cfg"
	"net/http"
)

func Register(c *gin.Context) {
	//POST — регистрация пользователя;
	// add to users + add to sessions
	var body struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	err := helperRegister(c, body.Login, body.Password)
	if err != nil {
		c.Status(http.StatusConflict)
		return
	}
	cookie, err := helperLogin(c, body.Login, body.Password)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.SetCookie("sessionid",
		cookie,
		cfg.Server.CookieLifetime,
		"/api/user",
		cfg.Server.Address,
		false,
		false)
	c.Status(200)
}
