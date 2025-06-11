package handler

import (
	"clinic/server/service"
	"clinic/server/structures"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(CORSMiddleware())

	admin := router.Group("/admin")
	{
		adminAuth := admin.Group("/auth")
		{
			adminAuth.POST("/sign-up", h.adminSignUp)
			adminAuth.POST("/sign-in", h.adminSignIn)
			adminAuth.POST("/refresh-token", h.adminRefreshToken)
			adminAuth.GET("/current-user", h.adminCurrentUser)
		}

		adminApi := admin.Group("/api", h.userIdentity(structures.AdminType))
		{
			devices := adminApi.Group("/device")
			{
				devices.POST("/", h.createDevice)
				devices.GET("/", h.getAllDevices)
				devices.GET("/:id", h.getDeviceById)
				devices.PUT("/:id", h.updateDevice)
				devices.DELETE("/:id", h.deleteDevice)
			}
		}
	}

	doctor := router.Group("/doctor")
	{
		doctorAuth := doctor.Group("/auth")
		{
			doctorAuth.POST("/sign-up", h.doctorSignUp)
			doctorAuth.POST("/sign-in", h.doctorSignIn)
			doctorAuth.POST("/refresh-token", h.doctorRefreshToken)
			doctorAuth.GET("/current-user", h.doctorCurrentUser)
		}

		doctorApi := doctor.Group("/api", h.userIdentity(structures.DoctorType))
		{
			patients := doctorApi.Group("/patients")
			{
				patients.GET("/", h.getAllPatients)
				patients.GET("/:id", h.getPatientById)
				patients.GET("/:id/full-info", h.getPatientFullInfo)
				patients.POST("/attach-relative", h.addRelativeToPatient)
				patients.POST("/attach-diagnosis", h.addDiagnosisToPatient)
			}
			medicines := doctorApi.Group("/medicine")
			{
				medicines.POST("/", h.createMedicine)
				medicines.GET("/", h.getAllMedicines)
				medicines.GET("/:id", h.getMedicineById)
				medicines.PUT("/:id", h.updateMedicine)
				medicines.DELETE("/:id", h.deleteMedicine)
			}
			diagnoses := doctorApi.Group("/diagnosis")
			{
				diagnoses.POST("/", h.createDiagnosis)
				diagnoses.GET("/", h.getAllDiagnoses)
				diagnoses.GET("/:id", h.getDiagnosisById)
				diagnoses.PUT("/:id", h.updateDiagnosis)
				diagnoses.DELETE("/:id", h.deleteDiagnosis)
			}
			treatmentPlans := doctorApi.Group("/treatmentplan")
			{
				treatmentPlans.POST("/", h.createTreatmentPlan)
				treatmentPlans.GET("/", h.getAllTreatmentPlans)
				treatmentPlans.GET("/:id", h.getTreatmentPlanById)
				treatmentPlans.PUT("/:id", h.updateTreatmentPlan)
				treatmentPlans.DELETE("/:id", h.deleteTreatmentPlan)
			}
			prescriptions := doctorApi.Group("/prescription")
			{
				prescriptions.POST("/", h.createPrescription)
				prescriptions.GET("/", h.getAllPrescriptions)
				prescriptions.GET("/:id", h.getPrescriptionById)
				prescriptions.PUT("/:id", h.updatePrescription)
				prescriptions.DELETE("/:id", h.deletePrescription)
			}
			healthNotes := doctorApi.Group("/healthnote")
			{
				healthNotes.GET("/patient/:id", h.getHealthNotesByPatientId)
			}
			visits := doctorApi.Group("/visit")
			{
				visits.POST("/", h.createVisit)
				visits.GET("/", h.getAllVisitsOfDoctor)
				visits.GET("/:id", h.getVisitById)
				visits.PUT("/:id", h.updateVisit)
				visits.DELETE("/:id", h.deleteVisit)
				visits.GET("/week", h.getVisitsForCurrentWeekOfDoctor)
			}

		}
	}

	relative := router.Group("/relative")
	{
		relativeAuth := relative.Group("/auth")
		{
			relativeAuth.POST("/sign-up", h.relativeSignUp)
			relativeAuth.POST("/sign-in", h.relativeSignIn)
			relativeAuth.POST("/refresh-token", h.relativeRefreshToken)
			relativeAuth.GET("/current-user", h.relativeCurrentUser)
		}

		relativeApi := relative.Group("/api", h.userIdentity(structures.RelativeType))
		{
			patients := relativeApi.Group("/patients")
			{
				patients.GET("/", h.getAllPatientsByRelativeId)
				patients.GET("/:id", h.isRelativeAllowedToSeeRecords, h.getPatientFullInfo)
				patients.GET("/:id/notes", h.isRelativeAllowedToSeeRecords, h.getHealthNotesByPatientId)
			}
		}
	}

	patient := router.Group("/patient")
	{
		patientAuth := patient.Group("/auth")
		{
			patientAuth.POST("/sign-up", h.patientSignUp)
			patientAuth.POST("/sign-in", h.patientSignIn)
			patientAuth.POST("/refresh-token", h.patientRefreshToken)
			patientAuth.GET("/current-user", h.patientCurrentUser)
		}

		patientApi := patient.Group("/api", h.userIdentity(structures.PatientType))
		{
			healthNotes := patientApi.Group("/healthnote")
			{
				healthNotes.POST("/", h.createHealthNote)
				healthNotes.GET("/", h.getAllHealthNotes)
				healthNotes.DELETE("/:id", h.deleteHealthNote)
			}
			patientApi.GET("/profile", h.getPatientsProfile)
			patientApi.GET("/relatives", h.getAttachedRelatives)
		}
	}

	indicators := router.Group("/indicators")
	{
		indicators.POST("/", h.createIndicatorsStamp)
	}

	indicatorsNotifications := router.Group("/indicatorsNotification")
	{
		indicatorsNotifications.GET("/", h.getAllIndicatorsNotifications)
		indicatorsNotifications.GET("/:id", h.getIndicatorsNotificationById)
	}

	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, DELETE, OPTIONS, PATCH")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
