package server

import (
	"fmt"
	"go-chat/internal/config"
	"go-chat/internal/server/handler"

	"github.com/gin-gonic/gin"
)

func Listen() {
	port := config.SERVER_PORT

	r := gin.Default()

	gin.SetMode(gin.ReleaseMode)

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/process-json", handler.HandleRegistration)

	go func() {
		if err := r.Run(":" + port); err != nil {
			fmt.Printf("Server error: %v\n", err)
		}
	}()

	fmt.Printf("Server listening on port %s\n", port)
}
