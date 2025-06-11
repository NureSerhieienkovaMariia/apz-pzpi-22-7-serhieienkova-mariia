package repository

import (
	"clinic/server/structures"
	"github.com/jmoiron/sqlx"
)

type AdminRepo interface {
	Create(admin structures.Admin) (int, error)
	GetByCreds(email, password string) (structures.Admin, error)
	GetById(adminId int) (structures.Admin, error)
}

type RelativeRepo interface {
	Create(relative structures.Relative) (int, error)
	GetByCreds(email, password string) (structures.Relative, error)
	GetById(relativeId int) (structures.Relative, error)
	GetAllByPatientId(id int) ([]structures.Relative, error)
}

type DoctorRepo interface {
	Create(doctor structures.Doctor) (int, error)
	GetByCreds(email, password string) (structures.Doctor, error)
	GetById(doctorId int) (structures.Doctor, error)
}

type PatientRepo interface {
	Create(patient structures.Patient) (int, error)
	GetAll() ([]structures.Patient, error)
	GetById(id int) (structures.Patient, error)
	Update(id int, input structures.Patient) error
	Delete(id int) error
	GetDiagnosesByPatientID(patientID int) ([]structures.Diagnosis, error)
	GetMedicinesByPatientID(patientID int) ([]structures.Medicine, error)
	GetDevicesByPatientID(patientID int) ([]structures.Device, error)
	GetIndicatorsByPatientID(patientID int) ([]structures.IndicatorsStamp, error)
	GetIndicatorsNotificationsByPatientID(patientID int) ([]structures.IndicatorsNotification, error)
	GetByCreds(email string, hash string) (structures.Patient, error)
	AttachRelative(relative structures.PatientRelative) error
	GetAllByRelativeId(id int) ([]structures.Patient, error)
	GetAllByDoctorId(id int) ([]structures.Patient, error)
}

type MedicineRepo interface {
	Create(medicine structures.Medicine) (int, error)
	GetAll() ([]structures.Medicine, error)
	Get(id int) (structures.Medicine, error)
	Update(id int, input structures.Medicine) error
	Delete(id int) error
}

type DiagnosisRepo interface {
	Create(diagnosis structures.Diagnosis) (int, error)
	GetAll() ([]structures.Diagnosis, error)
	Get(id int) (structures.Diagnosis, error)
	Update(id int, input structures.Diagnosis) error
	Delete(id int) error
	GetAllByPatientId(id int) ([]structures.Diagnosis, error)
	AttachDiagnosisToPatient(diagnosis structures.PatientDiagnosis) error
}

type DeviceRepo interface {
	Create(device structures.Device) (int, error)
	GetAll() ([]structures.Device, error)
	Get(id int) (structures.Device, error)
	Update(id int, input structures.Device) error
	Delete(id int) error
}

type IndicatorsNotificationRepo interface {
	Create(indicatorsNotification structures.IndicatorsNotification) (int, error)
	GetAll() ([]structures.IndicatorsNotification, error)
	Get(id int) (structures.IndicatorsNotification, error)
	GetAllByPatientID(patientID int) ([]structures.IndicatorsNotification, error)
}

type IndicatorsStampRepo interface {
	Create(indicatorsStamp structures.IndicatorsStamp) (int, error)
	GetAll() ([]structures.IndicatorsStamp, error)
	GetById(id int) (structures.IndicatorsStamp, error)
}

type HealthNoteRepo interface {
	Create(healthNote structures.HealthNote) (int, error)
	GetAll() ([]structures.HealthNote, error)
	Get(id int) (structures.HealthNote, error)
	Update(id int, input structures.HealthNote) error
	Delete(id int) error
	GetAllByPatientId(id int) ([]structures.HealthNote, error)
}

type PatientDiagnosisRepo interface {
	Create(patientDiagnosis structures.PatientDiagnosis) (int, error)
	GetAll() ([]structures.PatientDiagnosis, error)
	Get(id int) (structures.PatientDiagnosis, error)
	Update(id int, input structures.PatientDiagnosis) error
	Delete(id int) error
}

type PatientRelativeRepo interface {
	Create(patientRelative structures.PatientRelative) (int, error)
	GetAll() ([]structures.PatientRelative, error)
	Get(id int) (structures.PatientRelative, error)
	Update(id int, input structures.PatientRelative) error
	Delete(id int) error
	GetAllByRelativeId(relativeId int) ([]structures.PatientRelative, error)
	GetAllByPatientId(patientId int) ([]structures.PatientRelative, error)
}

type PrescriptionRepo interface {
	Create(prescription structures.Prescription) (int, error)
	GetAll() ([]structures.Prescription, error)
	Get(id int) (structures.Prescription, error)
	Update(id int, input structures.Prescription) error
	Delete(id int) error
	GetAllByTreatmentPlanId(id int) ([]structures.Prescription, error)
}

type TreatmentPlanRepo interface {
	Create(treatmentPlan structures.TreatmentPlan) (int, error)
	GetAll() ([]structures.TreatmentPlan, error)
	Get(id int) (structures.TreatmentPlan, error)
	Update(id int, input structures.TreatmentPlan) error
	Delete(id int) error
	GetByPatientId(id int) ([]structures.TreatmentPlan, error)
	GetAllByDoctorId(id int) ([]structures.TreatmentPlan, error)
}

type VisitRepo interface {
	Create(visit structures.Visit) (int, error)
	GetAll() ([]structures.Visit, error)
	Get(id int) (structures.Visit, error)
	Update(id int, input structures.Visit) error
	Delete(id int) error
	GetAllByTreatmentPlanId(id int) ([]structures.Visit, error)
	GetVisitsForNextDays(i int) ([]structures.Visit, error)
}

type PatientNotificationRepo interface {
	Create(notification structures.PatientNotification) (int, error)
	GetAll() ([]structures.PatientNotification, error)
	Get(id int) (structures.PatientNotification, error)
	Update(id int, input structures.PatientNotification) error
	Delete(id int) error
}

type Repository struct {
	AdminRepo
	RelativeRepo
	DoctorRepo
	PatientRepo
	MedicineRepo
	DiagnosisRepo
	DeviceRepo
	IndicatorsNotificationRepo
	IndicatorsStampRepo
	HealthNoteRepo
	PatientDiagnosisRepo
	PatientRelativeRepo
	PrescriptionRepo
	TreatmentPlanRepo
	VisitRepo
	PatientNotificationRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		AdminRepo:                  NewAdminPostgres(db),
		RelativeRepo:               NewRelativePostgres(db),
		DoctorRepo:                 NewDoctorPostgres(db),
		PatientRepo:                NewPatientPostgres(db),
		MedicineRepo:               NewMedicinePostgres(db),
		DiagnosisRepo:              NewDiagnosisPostgres(db),
		DeviceRepo:                 NewDevicePostgres(db),
		IndicatorsNotificationRepo: NewIndicatorsNotificationPostgres(db),
		IndicatorsStampRepo:        NewIndicatorsStampPostgres(db),
		HealthNoteRepo:             NewHealthNotePostgres(db),
		PatientDiagnosisRepo:       NewPatientDiagnosisPostgres(db),
		PatientRelativeRepo:        NewPatientRelativePostgres(db),
		PrescriptionRepo:           NewPrescriptionPostgres(db),
		TreatmentPlanRepo:          NewTreatmentPlanPostgres(db),
		VisitRepo:                  NewVisitPostgres(db),
		PatientNotificationRepo:    NewPatientNotificationPostgres(db),
	}
}
