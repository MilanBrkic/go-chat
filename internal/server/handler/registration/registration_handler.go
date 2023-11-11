package registration

import (
	"fmt"
	"go-chat/internal/database"
	"go-chat/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type registrationBody struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegistrationHandler struct {
	UserDb *database.UserDatabase
}

func (handler *RegistrationHandler) isValid(c *gin.Context) (*registrationBody, bool) {
	var jsonBody registrationBody
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

func (handler *RegistrationHandler) Handle(c *gin.Context) {
	body, ok := handler.isValid(c)

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
