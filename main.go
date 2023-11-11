package main

import (
	"go-chat/internal/database"
	"go-chat/internal/server"
	"go-chat/internal/util"
)

func main() {
	util.ReadEnvs()
	db := database.Connect()
	server.Listen(db)

	select {}
}
