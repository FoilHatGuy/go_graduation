package handlers

import (
	"context"
	"fmt"
	"go_graduation/internal/database"
	"go_graduation/internal/security"
)

var dbController = database.Controller

func helperRegister(ctx context.Context, login string, password string) (err error) {
	//POST — регистрация пользователя;
	// add to users + add to sessions
	newPass, err := security.Engine.HashPassword(password)
	if err != nil {
		return fmt.Errorf("[server:utils:register] while hashing\n%w", err)
	}
	err = dbController.AddUser(ctx, login, newPass)
	if err != nil {
		return fmt.Errorf("[server:utils:register] while adding to DB\n%w", err)
	}

	return nil
}
func helperLogin(ctx context.Context, login string, password string) (cookie string, err error) {
	hash, err := database.Controller.GetUser(ctx, login)
	if err != nil {
		return "", fmt.Errorf("[server:utils:login] while getting from DB\n%w", err)
	}

	security.Engine.ValidatePassword(password, hash)
	if err != nil {
		return "", fmt.Errorf("PW INCORRECT")
	}

	cookie, sid, err := security.Engine.GenerateCookie()
	if err != nil {
		return "", fmt.Errorf("[server:utils:login] while creating cookie\n%w", err)
	}

	err = database.Controller.AddSession(ctx, sid, login)
	if err != nil {
		return "", fmt.Errorf("[server:utils:login] while adding session\n%w", err)
	}
	return cookie, nil
	//POST  — аутентификация пользователя;

}
func helperOrdersGet(ctx context.Context, input string) {
	//GET — получение списка загруженных пользователем номеров заказов,
	//		статусов их обработки и информации о начислениях;
}
func helperBalance(ctx context.Context, input string) {
	//GET — получение текущего баланса счёта баллов лояльности пользователя;
}
func helperWithdraw(ctx context.Context, input string) {
	//POST — запрос на списание баллов с накопительного счёта в счёт оплаты нового заказа;
}
func helperWithdrawals(ctx context.Context, input string) {
	//GET — получение информации о выводе средств с накопительного счёта пользователем.
}
