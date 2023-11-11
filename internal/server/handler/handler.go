package handler

import (
	"go-chat/internal/database"
	"go-chat/internal/server/handler/login"
	"go-chat/internal/server/handler/registration"

	"github.com/gin-gonic/gin"
)

func GetRegistrationHandler(userDb *database.UserDatabase) func(c *gin.Context) {
	return (&registration.RegistrationHandler{UserDb: userDb}).Handle
}

func GetLoginHandler(userDb *database.UserDatabase) func(c *gin.Context) {
	return (&login.LoginHandler{UserDb: userDb}).Handle
}
