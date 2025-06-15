package handler

import (
	"clinic/server/service"
	"clinic/server/structures"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) patientSignUp(c *gin.Context) {
	var input structures.Patient
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.PatientAction.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := h.services.PatientAction.GenerateToken(input.Email, input.PasswordHash)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id":          id,
		"user_type":        structures.PatientType,
		"access_jwt_token": token.Token,
	})
}

func (h *Handler) patientSignIn(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.PatientAction.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id":          token.UserId,
		"user_type":        structures.PatientType,
		"access_jwt_token": token.Token,
	})
}

func (h *Handler) patientRefreshToken(c *gin.Context) {
	var input refreshTokenInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	newAccessToken, newRefreshToken, err := h.services.PatientAction.RefreshToken(input.RefreshToken)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id":           newAccessToken.UserId,
		"user_type":         structures.PatientType,
		"access_jwt_token":  newAccessToken.Token,
		"refresh_jwt_token": newRefreshToken.Token,
	})
}

func (h *Handler) patientCurrentUser(c *gin.Context) {
	rawAuthToken := readRawAuthToken(c)

	tokenClaims, err := service.ParseToken(rawAuthToken)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	user, err := h.services.PatientAction.GetById(tokenClaims.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":        user.Id,
		"email":     user.Email,
		"name":      user.Name,
		"surname":   user.Surname,
		"user_type": structures.PatientType,
		"birthday":  user.Birthday,
	})
}

func (h *Handler) getAllPatients(c *gin.Context) {
	patients, err := h.services.PatientAction.GetAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Clear password hashes for security before sending to client
	patientsResponse := []structures.Patient{}
	for _, patient := range patients {
		patient.PasswordHash = "" // Clear password hash for security
		patientsResponse = append(patientsResponse, patient)
	}

	c.JSON(200, patientsResponse)
}

func (h *Handler) getPatientById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	patient, err := h.services.PatientAction.GetById(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Clear password hash for security before sending to client
	patient.PasswordHash = "" // Clear password hash for security

	c.JSON(200, patient)
}

func (h *Handler) deletePatient(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	if err := h.services.PatientAction.Delete(id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "deleted"})
}

func (h *Handler) getPatientFullInfo(c *gin.Context) {
	patientID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid patient id")
		return
	}

	patient, err := h.services.PatientAction.GetById(patientID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	patient.PasswordHash = "" // Clear password hash for security

	// get diagnoses by patient id, and treatment plan by patient id
	diagnoses, err := h.services.DiagnosisAction.GetAllByPatientId(patientID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	treatmentPlans, err := h.services.TreatmentPlanAction.GetByPatientId(patientID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	treatmentPlansFullInfo := []structures.TreatmentPlanFullInfo{}

	for _, treatmentPlan := range treatmentPlans {

		// get doctor by doctor id from treatment plan
		doctor, err := h.services.DoctorAction.GetById(treatmentPlan.DoctorID)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		doctor.PasswordHash = "" // Clear password hash for security

		// get visits and prescriptions by treatment plan id
		visits, err := h.services.VisitAction.GetAllByTreatmentPlanId(treatmentPlan.ID)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		prescriptions, err := h.services.PrescriptionAction.GetAllByTreatmentPlanId(treatmentPlan.ID)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		// get array of PrescriptionInfo from prescriptions
		var prescriptionsFullInfo []structures.PrescriptionInfo
		for _, prescription := range prescriptions {
			medicine, err := h.services.MedicineAction.Get(prescription.MedicineID)
			if err != nil {
				newErrorResponse(c, http.StatusInternalServerError, err.Error())
				return
			}
			prescriptionsFullInfo = append(prescriptionsFullInfo, structures.PrescriptionInfo{
				Medicine:  medicine,
				Dosage:    prescription.Dosage,
				Frequency: prescription.Frequency,
			})
		}

		treatmentPlansFullInfo = append(treatmentPlansFullInfo, structures.TreatmentPlanFullInfo{
			Doctor:        doctor,
			StartDate:     treatmentPlans[0].StartDate,
			EndDate:       treatmentPlans[0].EndDate,
			Visits:        visits,
			Prescriptions: prescriptionsFullInfo,
		})
	}

	fullInfo := structures.PatientFullInfo{
		Patient:        patient,
		Diagnoses:      diagnoses,
		TreatmentPlans: treatmentPlansFullInfo,
	}

	c.JSON(http.StatusOK, fullInfo)
}

func (h *Handler) getPatientsProfile(c *gin.Context) {
	rawAuthToken := readRawAuthToken(c)

	tokenClaims, err := service.ParseToken(rawAuthToken)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	patientID := tokenClaims.Id
	patient, err := h.services.PatientAction.GetById(patientID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	patient.PasswordHash = "" // Clear password hash for security

	// get diagnoses by patient id, and treatment plan by patient id
	diagnoses, err := h.services.DiagnosisAction.GetAllByPatientId(patientID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	treatmentPlans, err := h.services.TreatmentPlanAction.GetByPatientId(patientID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	treatmentPlansFullInfo := []structures.TreatmentPlanFullInfo{}

	for _, treatmentPlan := range treatmentPlans {

		// get doctor by doctor id from treatment plan
		doctor, err := h.services.DoctorAction.GetById(treatmentPlan.DoctorID)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		doctor.PasswordHash = "" // Clear password hash for security

		// get visits and prescriptions by treatment plan id
		visits, err := h.services.VisitAction.GetAllByTreatmentPlanId(treatmentPlan.ID)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		prescriptions, err := h.services.PrescriptionAction.GetAllByTreatmentPlanId(treatmentPlan.ID)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		// get array of PrescriptionInfo from prescriptions
		var prescriptionsFullInfo []structures.PrescriptionInfo
		for _, prescription := range prescriptions {
			medicine, err := h.services.MedicineAction.Get(prescription.MedicineID)
			if err != nil {
				newErrorResponse(c, http.StatusInternalServerError, err.Error())
				return
			}
			prescriptionsFullInfo = append(prescriptionsFullInfo, structures.PrescriptionInfo{
				Medicine:  medicine,
				Dosage:    prescription.Dosage,
				Frequency: prescription.Frequency,
			})
		}

		treatmentPlansFullInfo = append(treatmentPlansFullInfo, structures.TreatmentPlanFullInfo{
			Doctor:        doctor,
			StartDate:     treatmentPlans[0].StartDate,
			EndDate:       treatmentPlans[0].EndDate,
			Visits:        visits,
			Prescriptions: prescriptionsFullInfo,
		})
	}

	fullInfo := structures.PatientFullInfo{
		Patient:        patient,
		Diagnoses:      diagnoses,
		TreatmentPlans: treatmentPlansFullInfo,
	}

	c.JSON(http.StatusOK, fullInfo)
}

// Add a relative to a patient
func (h *Handler) addRelativeToPatient(c *gin.Context) {
	var input structures.PatientRelative
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := h.services.PatientAction.AttachRelative(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "relative added"})
}

// Add a diagnosis to a patient
func (h *Handler) addDiagnosisToPatient(c *gin.Context) {
	var input structures.PatientDiagnosis
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := h.services.DiagnosisAction.AttachDiagnosisToPatient(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "diagnosis added"})
}

func (h *Handler) getAttachedRelatives(context *gin.Context) {
	rawAuthToken := readRawAuthToken(context)

	tokenClaims, err := service.ParseToken(rawAuthToken)
	if err != nil {
		newErrorResponse(context, http.StatusUnauthorized, err.Error())
		return
	}

	patientID := tokenClaims.Id
	relatives, err := h.services.RelativeAction.GetAllByPatientId(patientID)
	if err != nil {
		newErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, relatives)

}
