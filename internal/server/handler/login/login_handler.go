package login

import (
	"go-chat/internal/database"
	"go-chat/internal/server/handler_validator"
	"go-chat/internal/util"
	"net/http"

	"github.com/gin-gonic/gin"
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

	ok = util.CompareHash(body.Password, user.Password)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Password incorrect"})
		return
	}

	c.JSON(http.StatusOK, user)
}
