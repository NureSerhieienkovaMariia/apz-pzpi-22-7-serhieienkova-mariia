package handler

import (
	"clinic/server/service"
	"clinic/server/structures"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) relativeSignUp(c *gin.Context) {
	var input structures.Relative
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.RelativeAction.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := h.services.RelativeAction.GenerateToken(input.Email, input.PasswordHash)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id":          id,
		"user_type":        structures.RelativeType,
		"access_jwt_token": token.Token,
	})
}

func (h *Handler) relativeSignIn(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.RelativeAction.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id":          token.UserId,
		"user_type":        structures.RelativeType,
		"access_jwt_token": token.Token,
	})
}

func (h *Handler) relativeRefreshToken(c *gin.Context) {
	var input refreshTokenInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	newAccessToken, newRefreshToken, err := h.services.RelativeAction.RefreshToken(input.RefreshToken)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id":           newAccessToken.UserId,
		"user_type":         structures.RelativeType,
		"access_jwt_token":  newAccessToken.Token,
		"refresh_jwt_token": newRefreshToken.Token,
	})
}

func (h *Handler) relativeCurrentUser(c *gin.Context) {
	rawAuthToken := readRawAuthToken(c)

	tokenClaims, err := service.ParseToken(rawAuthToken)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	user, err := h.services.RelativeAction.GetById(tokenClaims.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":        user.Id,
		"email":     user.Email,
		"name":      user.Name,
		"surname":   user.Surname,
		"user_type": structures.RelativeType,
	})
}

func (h *Handler) getAllPatientsByRelativeId(context *gin.Context) {
	rawAuthToken := readRawAuthToken(context)

	tokenClaims, err := service.ParseToken(rawAuthToken)
	if err != nil {
		newErrorResponse(context, http.StatusUnauthorized, err.Error())
		return
	}

	patients, err := h.services.PatientAction.GetAllByRelativeId(tokenClaims.Id)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, patients)
}

func (h *Handler) isRelativeAllowedToSeeRecords(context *gin.Context) {
	rawAuthToken := readRawAuthToken(context)

	tokenClaims, err := service.ParseToken(rawAuthToken)
	if err != nil {
		newErrorResponse(context, http.StatusUnauthorized, err.Error())
		return
	}

	patientId := context.Param("id")
	if patientId == "" {
		newErrorResponse(context, http.StatusBadRequest, "invalid patient id")
		return
	}

	patientIdInt, err := strconv.Atoi(patientId)
	if err != nil {
		newErrorResponse(context, http.StatusBadRequest, "invalid patient id")
		return
	}

	isAllowed, err := h.services.PatientAction.IsRelativeAllowedToSeeRecords(tokenClaims.Id, patientIdInt)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	if !isAllowed {
		newErrorResponse(context, http.StatusForbidden, "you are not allowed to see this patient's records")
		return
	}
}
