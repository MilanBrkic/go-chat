package handler

import (
	"fmt"
	"go-chat/internal/database"
	"go-chat/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type registration struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type registrationHandler struct {
	userDb *database.UserDatabase
}

func GetRegistrationHandler(userDb *database.UserDatabase) func(c *gin.Context) {
	return (&registrationHandler{userDb: userDb}).handle
}

func (handler *registrationHandler) isValid(c *gin.Context) (*registration, bool) {
	var jsonBody registration
	var validate = validator.New()

	if err := c.ShouldBindJSON(&jsonBody); err != nil {
		fmt.Println("JSONUnmarshalMiddleware error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		c.Abort()
		return nil, false
	}

	if err := validate.Struct(jsonBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, false
	}

	return &jsonBody, true
}

func (handler *registrationHandler) handle(c *gin.Context) {
	body, ok := handler.isValid(c)

	if !ok {
		return
	}

	_, ok := handler.userDb.GetByUsername(body.Username)

	if ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Username already exists", "username": body.Username})
	}

	handler.userDb.Insert(model.User{Username: body.Username, Password: body.Password})

	c.JSON(http.StatusOK, gin.H{"message": "User registered"})
}
