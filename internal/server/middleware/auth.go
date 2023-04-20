package middleware

import (
	"github.com/gin-gonic/gin"
	"go_graduation/internal/cfg"
	"go_graduation/internal/security"
	"net/http"
)

func Cooker() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("user")
		var key string
		if err != nil {
			c.Status(http.StatusUnauthorized)
		}
		//fmt.Println("UID COOKIE PRESENT:\n", cookie)

		key, err = security.Engine.ValidateCookie(cookie)
		//fmt.Println("VALIDATION RESULT:\n", key, err)
		if err != nil {
			c.Status(http.StatusUnauthorized)
		}
		c.SetCookie("user", cookie, cfg.Server.CookieLifetime, "/", cfg.Server.Address, false, true)
		c.Set("owner", key)
		//fmt.Println("UID KEY:\n", key)
		c.Next()
		return
	}
}
