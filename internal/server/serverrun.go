package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_graduation/internal/cfg"
	"log"
)

func Run() {
	r := gin.Default()
	baseRouter := r.Group("")
	{
		//baseRouter.Use(Gzip())
		//baseRouter.Use(Gunzip())
		api := baseRouter.Group("/api")
		{
			api.GET("/ping", ping)
		}
	}

	//POST /api/user/register — регистрация пользователя;
	//POST /api/user/login — аутентификация пользователя;
	//POST /api/user/orders — загрузка пользователем номера заказа для расчёта;
	//GET /api/user/orders — получение списка загруженных пользователем номеров заказов, статусов их обработки и информации о начислениях;
	//GET /api/user/balance — получение текущего баланса счёта баллов лояльности пользователя;
	//POST /api/user/balance/withdraw — запрос на списание баллов с накопительного счёта в счёт оплаты нового заказа;
	//GET /api/user/withdrawals — получение информации о выводе средств с накопительного счёта пользователем.

	fmt.Println("SERVER LISTENING ON", cfg.Server.Address)
	log.Fatal(r.Run(cfg.Server.Address))
}
