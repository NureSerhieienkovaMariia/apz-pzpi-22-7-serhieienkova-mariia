package service

import (
	"clinic/server/repository"
	"clinic/server/structures"
)

type AdminAction interface {
	Create(admin structures.Admin) (int, error)
	GenerateToken(email, password string) (structures.UserToken, error)
	RefreshToken(refreshToken string) (structures.UserToken, structures.UserToken, error)
	GetById(adminId int) (structures.Admin, error)
}

type RelativeAction interface {
	Create(relative structures.Relative) (int, error)
	GenerateToken(email, password string) (structures.UserToken, error)
	RefreshToken(refreshToken string) (structures.UserToken, structures.UserToken, error)
	GetById(relativeId int) (structures.Relative, error)
	GetAllByPatientId(patientId int) ([]structures.Relative, error)
	Delete(id string) error
	Update(id string, input structures.Relative) error
}

type DoctorAction interface {
	Create(doctor structures.Doctor) (int, error)
	GenerateToken(email, password string) (structures.UserToken, error)
	RefreshToken(refreshToken string) (structures.UserToken, structures.UserToken, error)
	GetById(doctorId int) (structures.Doctor, error)
}

type PatientAction interface {
	Create(patient structures.Patient) (int, error)
	GenerateToken(email, password string) (structures.UserToken, error)
	RefreshToken(refreshToken string) (structures.UserToken, structures.UserToken, error)
	GetById(patientId int) (structures.Patient, error)
	GetAll() ([]structures.Patient, error)
	AttachRelative(patientRelative structures.PatientRelative) error
	GetAllByRelativeId(relativeId int) ([]structures.Patient, error)
	GetAllByDoctorId(doctorId int) ([]structures.Patient, error)
	Delete(id int) error
	IsRelativeAllowedToSeeRecords(relativeId int, patientId int) (bool, error)
	Update(id string, input structures.Patient) error
}

type MedicineAction interface {
	Create(medicine structures.Medicine) (int, error)
	GetAll() ([]structures.Medicine, error)
	Get(id int) (structures.Medicine, error)
	Update(id int, input structures.Medicine) error
	Delete(id int) error
}

type DiagnosisAction interface {
	Create(diagnosis structures.Diagnosis) (int, error)
	GetAll() ([]structures.Diagnosis, error)
	Get(id int) (structures.Diagnosis, error)
	Update(id int, input structures.Diagnosis) error
	Delete(id int) error
	GetAllByPatientId(patientId int) ([]structures.Diagnosis, error)
	AttachDiagnosisToPatient(patientDiagnosis structures.PatientDiagnosis) error
}

type DeviceAction interface {
	Create(device structures.Device) (int, error)
	GetAll() ([]structures.Device, error)
	Get(id int) (structures.Device, error)
	Update(id int, input structures.Device) error
	Delete(id int) error
}

type IndicatorsNotificationAction interface {
	Create(indicatorsNotification structures.IndicatorsNotification) (int, error)
	GetAll() ([]structures.IndicatorsNotification, error)
	Get(id int) (structures.IndicatorsNotification, error)
	GetAllByPatientID(patientID int) ([]structures.IndicatorsNotification, error)
}

type IndicatorsStampAction interface {
	Create(input structures.IndicatorsStamp) error
	GetAll() ([]structures.IndicatorsStamp, error)
	GetById(id int) (structures.IndicatorsStamp, error)
}

type VisitAction interface {
	Create(visit structures.Visit) (int, error)
	GetAll() ([]structures.Visit, error)
	Get(id int) (structures.Visit, error)
	Update(id int, input structures.Visit) error
	Delete(id int) error
	GetAllByTreatmentPlanId(treatmentPlanId int) ([]structures.Visit, error)
	GetAllTodaysVisits() ([]structures.Visit, error)
	GetAllWeeksVisits() ([]structures.Visit, error)
}

type TreatmentPlanAction interface {
	Create(treatmentPlan structures.TreatmentPlan) (int, error)
	GetAll() ([]structures.TreatmentPlan, error)
	Get(id int) (structures.TreatmentPlan, error)
	Update(id int, input structures.TreatmentPlan) error
	Delete(id int) error
	GetByPatientId(patientId int) ([]structures.TreatmentPlan, error)
	GetAllByDoctorId(doctorId int) ([]structures.TreatmentPlan, error)
}

type PrescriptionAction interface {
	Create(prescription structures.Prescription) (int, error)
	GetAll() ([]structures.Prescription, error)
	Get(id int) (structures.Prescription, error)
	Update(id int, input structures.Prescription) error
	Delete(id int) error
	GetAllByTreatmentPlanId(patientId int) ([]structures.Prescription, error)
}

type HealthNoteAction interface {
	Create(healthNote structures.HealthNote) (int, error)
	GetAll() ([]structures.HealthNote, error)
	Get(id int) (structures.HealthNote, error)
	Update(id int, input structures.HealthNote) error
	Delete(id int) error
	GetAllByPatientId(patientId int) ([]structures.HealthNote, error)
}

type PatientNotificationAction interface {
	Create(notification structures.PatientNotification) (int, error)
	GetAll() ([]structures.PatientNotification, error)
	Get(id int) (structures.PatientNotification, error)
	Update(id int, input structures.PatientNotification) error
	Delete(id int) error
}

type Service struct {
	AdminAction
	RelativeAction
	DoctorAction
	PatientAction
	MedicineAction
	DiagnosisAction
	DeviceAction
	IndicatorsNotificationAction
	IndicatorsStampAction
	VisitAction
	TreatmentPlanAction
	PrescriptionAction
	HealthNoteAction
	PatientNotificationAction
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		AdminAction:                  NewAdminService(repos.AdminRepo),
		RelativeAction:               NewRelativeService(repos.RelativeRepo),
		DoctorAction:                 NewDoctorService(repos.DoctorRepo),
		PatientAction:                NewPatientActionService(repos.PatientRepo, repos.PatientRelativeRepo),
		MedicineAction:               NewMedicineActionService(repos.MedicineRepo),
		DiagnosisAction:              NewDiagnosisActionService(repos.DiagnosisRepo),
		DeviceAction:                 NewDeviceActionService(repos.DeviceRepo),
		IndicatorsNotificationAction: NewIndicatorsNotificationActionService(repos.IndicatorsNotificationRepo),
		IndicatorsStampAction:        NewIndicatorsStampActionService(repos.IndicatorsStampRepo, repos.IndicatorsNotificationRepo, repos.DeviceRepo, repos.PatientRepo, repos.PatientRelativeRepo, repos.RelativeRepo, repos.TreatmentPlanRepo, repos.DoctorRepo),
		VisitAction:                  NewVisitActionService(repos.VisitRepo),
		TreatmentPlanAction:          NewTreatmentPlanActionService(repos.TreatmentPlanRepo),
		PrescriptionAction:           NewPrescriptionActionService(repos.PrescriptionRepo),
		HealthNoteAction:             NewHealthNoteActionService(repos.HealthNoteRepo),
		PatientNotificationAction:    NewPatientNotificationService(repos.PatientNotificationRepo),
	}
}
