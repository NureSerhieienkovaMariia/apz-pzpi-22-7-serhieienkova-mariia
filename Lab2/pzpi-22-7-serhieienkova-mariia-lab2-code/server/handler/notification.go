package handler

import (
	"clinic/server/structures"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) createIndicatorsNotification(c *gin.Context) {
	var input structures.IndicatorsNotification
	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id, err := h.services.IndicatorsNotificationAction.Create(input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"id": id})
}

func (h *Handler) getAllIndicatorsNotifications(c *gin.Context) {
	indicatorsNotifications, err := h.services.IndicatorsNotificationAction.GetAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, indicatorsNotifications)
}

func (h *Handler) getIndicatorsNotificationById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	indicatorsNotification, err := h.services.IndicatorsNotificationAction.Get(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, indicatorsNotification)
}

func (h *Handler) getAllIndicatorsNotificationsByPatientID(c *gin.Context) {
	patientID, err := strconv.Atoi(c.Param("patient_id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid patient id"})
		return
	}

	indicatorsNotifications, err := h.services.IndicatorsNotificationAction.GetAllByPatientID(patientID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, indicatorsNotifications)
}
