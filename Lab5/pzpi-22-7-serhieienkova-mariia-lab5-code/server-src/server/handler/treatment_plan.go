package handler

import (
	"clinic/server/structures"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) createTreatmentPlan(c *gin.Context) {
	var input structures.TreatmentPlan
	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id, err := h.services.TreatmentPlanAction.Create(input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"id": id})
}

func (h *Handler) getAllTreatmentPlans(c *gin.Context) {
	treatmentPlans, err := h.services.TreatmentPlanAction.GetAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, treatmentPlans)
}

func (h *Handler) getTreatmentPlanById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	treatmentPlan, err := h.services.TreatmentPlanAction.Get(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, treatmentPlan)
}

func (h *Handler) updateTreatmentPlan(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	var input structures.TreatmentPlan
	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := h.services.TreatmentPlanAction.Update(id, input); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "updated"})
}

func (h *Handler) deleteTreatmentPlan(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	if err := h.services.TreatmentPlanAction.Delete(id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "deleted"})
}
