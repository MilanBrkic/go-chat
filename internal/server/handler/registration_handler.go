package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Registration struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func HandleRegistration(c *gin.Context) {
	body, ok := isValid(c)

	if !ok {
		return
	}

	c.JSON(http.StatusOK, body)
}

func isValid(c *gin.Context) (*Registration, bool) {
	var jsonBody Registration
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
