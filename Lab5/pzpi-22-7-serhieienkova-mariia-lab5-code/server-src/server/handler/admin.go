package handler

import (
	"clinic/server/service"
	"clinic/server/structures"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) adminSignUp(c *gin.Context) {
	var input structures.Admin
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.AdminAction.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := h.services.AdminAction.GenerateToken(input.Email, input.PasswordHash)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id":          id,
		"user_type":        structures.AdminType,
		"access_jwt_token": token.Token,
	})
}

func (h *Handler) adminSignIn(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.AdminAction.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id":          token.UserId,
		"user_type":        structures.AdminType,
		"access_jwt_token": token.Token,
	})
}

func (h *Handler) adminRefreshToken(c *gin.Context) {
	var input refreshTokenInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	newAccessToken, newRefreshToken, err := h.services.AdminAction.RefreshToken(input.RefreshToken)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id":           newAccessToken.UserId,
		"user_type":         structures.AdminType,
		"access_jwt_token":  newAccessToken.Token,
		"refresh_jwt_token": newRefreshToken.Token,
	})
}

func (h *Handler) adminCurrentUser(c *gin.Context) {
	rawAuthToken := readRawAuthToken(c)

	tokenClaims, err := service.ParseToken(rawAuthToken)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	user, err := h.services.AdminAction.GetById(tokenClaims.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":        user.Id,
		"email":     user.Email,
		"name":      user.Name,
		"surname":   user.Surname,
		"user_type": structures.AdminType,
	})
}

func (h *Handler) adminDeletePatient(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.PatientAction.Delete(intId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Patient deleted successfully",
	})
}

func (h *Handler) adminUpdatePatient(c *gin.Context) {
	id := c.Param("id")
	var input structures.Patient
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := h.services.PatientAction.Update(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Patient updated successfully",
	})
}

func (h *Handler) adminDeleteRelative(c *gin.Context) {
	id := c.Param("id")

	err := h.services.RelativeAction.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Relative deleted successfully",
	})
}

func (h *Handler) adminUpdateRelative(c *gin.Context) {
	id := c.Param("id")
	var input structures.Relative
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := h.services.RelativeAction.Update(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Relative updated successfully",
	})
}
