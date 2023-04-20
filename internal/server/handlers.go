package server

import (
	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.Status(200)
}
