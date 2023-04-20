package middleware

import (
	"github.com/gin-gonic/gin"
	"go_graduation/internal/cfg"
	"go_graduation/internal/database"
	"go_graduation/internal/security"
	"net/http"
)

func Cooker() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("sid")
		var key string
		if err != nil {
			c.Status(http.StatusUnauthorized)
		}

		key, err = security.Engine.DecypherCookie(cookie)
		if err != nil {
			c.Status(http.StatusUnauthorized)
		}
		login, err := database.Controller.GetSession(c, key)
		if err != nil {
			return
		}
		c.SetCookie("sid", cookie, cfg.Server.CookieLifetime, "/", cfg.Server.Address, false, true)
		c.Set("uid", login)
		//fmt.Println("UID KEY:\n", key)
		c.Next()
		return
	}
}
