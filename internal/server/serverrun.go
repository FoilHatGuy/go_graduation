package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_graduation/internal/cfg"
	"go_graduation/internal/server/handlers"
	"go_graduation/internal/server/middleware"
	"log"
)

func Run() {
	r := gin.Default()
	baseRouter := r.Group("")
	{
		baseRouter.Use(middleware.Gzip())
		baseRouter.Use(middleware.Gunzip())
		baseRouter.GET("/ping", handlers.Ping)
		api := baseRouter.Group("/api/user")
		{
			api.POST("/register", handlers.Register)
			api.POST("/login", handlers.Login)
			authorized := baseRouter.Group("/").Use(middleware.Cooker())
			{
				authorized.POST("/orders", handlers.OrdersPost)
				authorized.GET("/orders", handlers.OrdersGet)
				authorized.GET("/balance", handlers.Balance)
				authorized.POST("/balance/withdraw", handlers.Withdraw)
				authorized.GET("/withdrawals", handlers.Withdrawals)
			}
		}
	}

	fmt.Println("SERVER LISTENING ON", cfg.Server.Address)
	log.Fatal(r.Run(cfg.Server.Address))
}
