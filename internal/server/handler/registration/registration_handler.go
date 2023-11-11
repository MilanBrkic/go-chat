package registration

import (
	"go-chat/internal/database"
	"go-chat/internal/model"
	"go-chat/internal/server/handler_validator"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type body struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegistrationHandler struct {
	UserDb *database.UserDatabase
}

func (handler *RegistrationHandler) Handle(c *gin.Context) {
	body, ok := handler_validator.IsValid[body](c)

	if !ok {
		return
	}

	_, exists := handler.UserDb.GetByUsername(body.Username)

	if exists {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Username already exists", "username": body.Username})
		return
	}

	hashedPassword, err := hashPassword(body.Password)

	if err != nil {
		c.JSON(500, gin.H{
			"error":   err,
			"message": "Error while hashing password",
		})
		return
	}

	handler.UserDb.Insert(model.User{Username: body.Username, Password: hashedPassword})

	c.JSON(http.StatusOK, gin.H{"message": "User registered"})
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
