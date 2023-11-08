package main

import (
	"go-chat/internal/database"
	"go-chat/internal/server"
	"go-chat/internal/util"
)

func main() {
	util.ReadEnvs()
	database.Connect()
	server.Listen()

	select {}
}
