package main

import (
	"go_graduation/internal/cfg"
	"go_graduation/internal/server"
)

func main() {
	cfg.Initialize()
	//database.InitDB()/
	server.Run()

	//eng, _ := security.Init()
	//cookie, s, err := eng.GenerateCookie()
	//if err != nil {
	//	return
	//}
	//fmt.Println(cookie, s)
}
