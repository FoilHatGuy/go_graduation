package main

import (
	"fmt"
	"go_graduation/internal/security"
)

func main() {
	//server.Run()

	eng, _ := security.Init()
	cookie, s, err := eng.GenerateCookie()
	if err != nil {
		return
	}
	fmt.Println(cookie, s)
}
