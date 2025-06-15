package structures

import "time"

type PatientNotification struct {
	ID        int       `json:"id" db:"id"`
	Timestamp time.Time `json:"timestamp" db:"timestamp"`
	PatientID int       `json:"patient_id" db:"patient_id"`
	Topic     string    `json:"topic" db:"topic"`
	Message   string    `json:"message" db:"message"`
}
