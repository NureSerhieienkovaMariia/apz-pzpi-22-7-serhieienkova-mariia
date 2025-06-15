package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type refreshTokenInput struct {
	RefreshToken string `json:"refresh_jwt_token" binding:"required"`
}

func readRawAuthToken(c *gin.Context) string {
	header := c.GetHeader("Authorization")
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		newErrorResponse(c, http.StatusUnauthorized, "token is empty")
	}
	return headerParts[1]
}
