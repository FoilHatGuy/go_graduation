package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go_graduation/internal/cfg"
	"net/http"
)

func Login(c *gin.Context) {
	//POST  — аутентификация пользователя;
	var body struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.Status(http.StatusBadRequest)
	}
	cookie, err := helperLogin(c, body.Login, body.Password)
	if err != nil {
		if errors.Is(err, errors.New("PW INCORRECT")) {
			c.Status(http.StatusUnauthorized)
		} else {
			c.Status(http.StatusBadRequest)
		}
		return
	}
	c.SetCookie("sid",
		cookie,
		cfg.Server.CookieLifetime,
		"/api/user",
		cfg.Server.Address,
		false,
		false)
	c.Status(200)
}
