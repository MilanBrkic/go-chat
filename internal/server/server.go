package server

import (
	"fmt"
	"go-chat/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSONUnmarshalMiddleware(c *gin.Context) {
	var requestData map[string]interface{}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		c.Abort()
		return
	}

	c.Set("jsonData", requestData)
	c.Next()
}

func Listen() {
	port := config.SERVER_PORT

	r := gin.Default()

	gin.SetMode(gin.ReleaseMode)

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(JSONUnmarshalMiddleware)

	r.POST("/process-json", func(c *gin.Context) {
		jsonData, ok := c.MustGet("jsonData").(map[string]interface{})
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "JSON data not found in request"})
			return
		}

		field1, _ := jsonData["milan"].(string)

		response := gin.H{
			"Field1": field1,
		}

		c.JSON(http.StatusOK, response)
	})

	go func() {
		if err := r.Run(":" + port); err != nil {
			fmt.Printf("Server error: %v\n", err)
		}
	}()

	fmt.Printf("Server listening on port %s\n", port)
}
