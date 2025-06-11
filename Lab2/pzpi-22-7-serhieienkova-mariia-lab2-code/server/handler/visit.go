package handler

import (
	"clinic/server/service"
	"clinic/server/structures"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createVisit(c *gin.Context) {
	var input structures.Visit
	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id, err := h.services.VisitAction.Create(input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"id": id})
}

func (h *Handler) getAllVisits(c *gin.Context) {
	visits, err := h.services.VisitAction.GetAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, visits)
}

func (h *Handler) getVisitById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	visit, err := h.services.VisitAction.Get(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, visit)
}

func (h *Handler) updateVisit(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	var input structures.Visit
	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := h.services.VisitAction.Update(id, input); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "updated"})
}

func (h *Handler) deleteVisit(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	if err := h.services.VisitAction.Delete(id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "deleted"})
}

func (h *Handler) getVisitsForCurrentWeekOfDoctor(c *gin.Context) {
	rawAuthToken := readRawAuthToken(c)

	tokenClaims, err := service.ParseToken(rawAuthToken)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	visits, err := h.services.VisitAction.GetAllWeeksVisits()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Filter visits by the current doctor's ID
	doctorId := tokenClaims.Id
	var filteredVisits []structures.Visit
	for _, visit := range visits {
		treatmentPlanId := visit.TreatmentPlanID
		treatmentPlan, err := h.services.TreatmentPlanAction.Get(treatmentPlanId)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		if treatmentPlan.DoctorID == doctorId {
			filteredVisits = append(filteredVisits, visit)
		}
	}

	c.JSON(200, filteredVisits)
}

func (h *Handler) getAllVisitsOfDoctor(c *gin.Context) {
	rawAuthToken := readRawAuthToken(c)

	tokenClaims, err := service.ParseToken(rawAuthToken)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	doctorId := tokenClaims.Id
	visits, err := h.services.VisitAction.GetAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var doctorVisits []structures.Visit
	for _, visit := range visits {
		treatmentPlanId := visit.TreatmentPlanID
		treatmentPlan, err := h.services.TreatmentPlanAction.Get(treatmentPlanId)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		if treatmentPlan.DoctorID == doctorId {
			doctorVisits = append(doctorVisits, visit)
		}
	}

	c.JSON(200, doctorVisits)
}
