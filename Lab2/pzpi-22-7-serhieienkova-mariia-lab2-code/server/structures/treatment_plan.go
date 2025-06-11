package structures

import "time"

type TreatmentPlan struct {
	ID        int       `json:"id" db:"id"`
	PatientID int       `json:"patient_id" db:"patient_id"`
	DoctorID  int       `json:"doctor_id" db:"doctor_id"`
	StartDate time.Time `json:"start_date" db:"start_date"`
	EndDate   time.Time `json:"end_date" db:"end_date"`
}
