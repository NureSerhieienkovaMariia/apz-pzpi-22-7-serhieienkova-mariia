package structures

import "time"

type Visit struct {
	ID              int       `json:"id" db:"id"`
	TreatmentPlanID int       `json:"treatment_plan_id" db:"treatment_plan_id"`
	Reason          string    `json:"reason" db:"reason"`
	Date            time.Time `json:"date" db:"date"`
	Notes           string    `json:"notes" db:"note"`
}
