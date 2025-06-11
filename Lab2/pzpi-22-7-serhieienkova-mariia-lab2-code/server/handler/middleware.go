package handler

import (
	"clinic/server/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userIdCtx           = "userId"
)

func (h *Handler) userIdentity(userType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader(authorizationHeader)
		if header == "" {
			newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
			return
		}

		if len(headerParts[1]) == 0 {
			newErrorResponse(c, http.StatusUnauthorized, "token is empty")
			return
		}

		tokenClaims, err := service.ParseToken(headerParts[1])
		if err != nil {
			newErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}

		if tokenClaims.UserType != userType {
			newErrorResponse(c, http.StatusUnauthorized, "invalid user type to execute this action")
			return
		}

		c.Set(userIdCtx, tokenClaims.Id)
	}
}

//func getUserId(c *gin.Context) (int, error) {
//	id, ok := c.Get(userIdCtx)
//	if !ok {
//		return 0, errors.New("user id not found")
//	}
//
//	idInt, ok := id.(int)
//	if !ok {
//		return 0, errors.New("user id is of invalid type")
//	}
//
//	return idInt, nil
//}
