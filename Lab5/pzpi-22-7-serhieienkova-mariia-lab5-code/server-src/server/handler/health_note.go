package handler

import (
	"clinic/server/service"
	"clinic/server/structures"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createHealthNote(c *gin.Context) {
	var input structures.HealthNote
	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("Received health note input:", input)
	id, err := h.services.HealthNoteAction.Create(input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"id": id})
}

func (h *Handler) getAllHealthNotes(c *gin.Context) {
	rawAuthToken := readRawAuthToken(c)

	tokenClaims, err := service.ParseToken(rawAuthToken)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	patientID := tokenClaims.Id
	healthNotes, err := h.services.HealthNoteAction.GetAllByPatientId(patientID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, healthNotes)
}

func (h *Handler) getHealthNoteById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	healthNote, err := h.services.HealthNoteAction.Get(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, healthNote)
}

func (h *Handler) updateHealthNote(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	var input structures.HealthNote
	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := h.services.HealthNoteAction.Update(id, input); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "updated"})
}

func (h *Handler) deleteHealthNote(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	if err := h.services.HealthNoteAction.Delete(id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "deleted"})
}

func (h *Handler) getHealthNotesByPatientId(context *gin.Context) {
	patientId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(400, gin.H{"error": "invalid patient id"})
		return
	}
	healthNotes, err := h.services.HealthNoteAction.GetAllByPatientId(patientId)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}
	context.JSON(200, healthNotes)
}
