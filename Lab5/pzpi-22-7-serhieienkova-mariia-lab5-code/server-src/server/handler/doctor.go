package handler

import (
	"clinic/server/service"
	"clinic/server/structures"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) doctorSignUp(c *gin.Context) {
	var input structures.Doctor
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.DoctorAction.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := h.services.DoctorAction.GenerateToken(input.Email, input.PasswordHash)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id":          id,
		"user_type":        structures.DoctorType,
		"access_jwt_token": token.Token,
	})
}

func (h *Handler) doctorSignIn(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.DoctorAction.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id":          token.UserId,
		"user_type":        structures.DoctorType,
		"access_jwt_token": token.Token,
	})
}

func (h *Handler) doctorRefreshToken(c *gin.Context) {
	var input refreshTokenInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	newAccessToken, newRefreshToken, err := h.services.DoctorAction.RefreshToken(input.RefreshToken)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id":           newAccessToken.UserId,
		"user_type":         structures.DoctorType,
		"access_jwt_token":  newAccessToken.Token,
		"refresh_jwt_token": newRefreshToken.Token,
	})
}

func (h *Handler) doctorCurrentUser(c *gin.Context) {
	rawAuthToken := readRawAuthToken(c)

	tokenClaims, err := service.ParseToken(rawAuthToken)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	user, err := h.services.DoctorAction.GetById(tokenClaims.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":        user.Id,
		"email":     user.Email,
		"name":      user.Name,
		"surname":   user.Surname,
		"user_type": structures.DoctorType,
	})
}
