package handler_validator

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func IsValid[T any](c *gin.Context) (*T, bool) {
	var jsonBody T
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
