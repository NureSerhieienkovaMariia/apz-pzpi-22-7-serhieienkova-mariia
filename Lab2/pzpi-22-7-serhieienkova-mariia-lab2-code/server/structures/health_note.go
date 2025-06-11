package structures

import (
	"time"
)

type HealthNote struct {
	ID        int       `json:"id" db:"id"`
	PatientID int       `json:"patient_id" db:"patient_id"`
	Timestamp time.Time `json:"timestamp" db:"timestamp"`
	Note      string    `json:"note" db:"note"`
}
