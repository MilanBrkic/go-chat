package login

import (
	"fmt"
	"go-chat/internal/database"
	"go-chat/internal/server/handler_validator"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type body struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginHandler struct {
	UserDb *database.UserDatabase
}

func (h *LoginHandler) Handle(c *gin.Context) {
	body, ok := handler_validator.IsValid[body](c)

	if !ok {
		return
	}

	user, exists := h.UserDb.GetByUsername(body.Username)

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Username does not exist", "username": body.Username})
		return
	}

	ok = compareHash(body.Password, user.Password)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Password incorrect"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User logged in"})
}

func compareHash(enteredPassword string, storedHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(enteredPassword))
	if err == nil {
		return true
	} else if err == bcrypt.ErrMismatchedHashAndPassword {
		return false
	} else {
		panic(fmt.Sprintf("Error: %s", err))
	}
}
