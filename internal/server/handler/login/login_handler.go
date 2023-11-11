package login

import (
	"fmt"
	"go-chat/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type loginBody struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginHandler struct {
	UserDb *database.UserDatabase
}

func (handler *LoginHandler) isValid(c *gin.Context) (*loginBody, bool) {
	var jsonBody loginBody
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

func (handler *LoginHandler) Handle(c *gin.Context) {
	body, ok := handler.isValid(c)

	if !ok {
		return
	}

	user, exists := handler.UserDb.GetByUsername(body.Username)

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
